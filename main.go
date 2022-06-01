//go:generate abigen --sol web3/Purpose.sol --pkg web3 --out web3/purpose.go
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/purpose.proto
//go:generate protoc -I=./grpc purpose.proto --js_out=import_style=commonjs:../frontend/src/grpc/build --grpc-web_out=import_style=typescript,mode=grpcwebtext:../frontend/src/grpc/build

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/adamavixio/goethix"
	"github.com/adamavixio/logger"
	pb "github.com/adamavixio/purpose-backend/gRPC"
	"github.com/adamavixio/purpose-backend/web3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc"
)

var (
	PORT        = *flag.Int("port", 50051, "server port")
	NODE        = *flag.String("eth-address", "ws://127.0.0.1:7545", "eth server address")
	CONTRACT    = *flag.String("contract", "0x5a0994de6d4596d2867A481feF494B447655a6B3", "contract address")
	PRIVATE_KEY = *flag.String("private-key", "84bf34442cd7d3dee01465398716034734aa445ac1fdef3dd4836cea6b5f4e37", "private key")

	ethix   *goethix.Ethix
	purpose *web3.Purpose
	pubsub  = PubSub{sync.Mutex{}, map[pb.Purpose_MissionStreamServer]struct{}{}}
)

type Server struct {
	pb.UnimplementedPurposeServer
}

//
// Procedure
//

func main() {
	done := make(chan struct{})

	go ConnectWeb3()
	go StartServer()

	<-done
}

//
// Tasks
//

func ConnectWeb3() {
	ethix = goethix.NewEthix()

	err := ethix.Dial(NODE)
	if err != nil {
		logger.Fatal(err, "unable to dial ethereum node")
	}

	contract := common.HexToAddress(CONTRACT)
	purpose, err = web3.NewPurpose(contract, ethix.Client())
	if err != nil {
		logger.Fatal(err, "unable to create purpose web3 client")
	}

	addMissionChannel := make(chan *web3.PurposeAddedMission)
	addMissionSub, err := purpose.WatchAddedMission(&bind.WatchOpts{}, addMissionChannel)
	if err != nil {
		logger.Fatal(err, "unable to subscribe to addMission channel")
	}

	removeMissionChannel := make(chan *web3.PurposeRemovedMission)
	removeMissionSub, err := purpose.WatchRemovedMission(&bind.WatchOpts{}, removeMissionChannel)
	if err != nil {
		logger.Fatal(err, "unable to subscribe to removeMission channel")
	}

	for {
		select {
		case addedMission := <-addMissionChannel:
			pubsub.Publish(&pb.MissionStreamResponse{
				Type:     "ADD",
				Missions: MissionsWebToPB([]web3.Mission{addedMission.Mission}),
			})
		case err := <-addMissionSub.Err():
			logger.Warn(err, "error on emit addMission")
		case removedMission := <-removeMissionChannel:
			pubsub.Publish(&pb.MissionStreamResponse{
				Type:     "REMOVE",
				Missions: MissionsWebToPB([]web3.Mission{removedMission.Mission}),
			})
		case err := <-removeMissionSub.Err():
			logger.Warn(err, "error on emit removeMission")
		}
	}
}

func StartServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPurposeServer(s, &Server{})

	logger.Info("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//
// Actions
//

func (s *Server) MissionStream(in *pb.MissionStreamRequest, stream pb.Purpose_MissionStreamServer) error {
	done, err := pubsub.Subscribe(stream)
	if err != nil {
		pubsub.Unsubscribe(stream)
		return err
	}

	<-done
	return nil
}

func (s *Server) AddMission(ctx context.Context, in *pb.AddMissionRequest) (*pb.AddMissionResponse, error) {
	auth, err := ethix.Authorize(PRIVATE_KEY)
	if err != nil {
		return nil, err
	}

	_, err = purpose.AddMission(auth, in.Mission.Title)
	if err != nil {
		return nil, err
	}

	return &pb.AddMissionResponse{}, nil
}

func (s *Server) RemoveMission(ctx context.Context, in *pb.RemoveMissionRequest) (*pb.RemoveMissionResponse, error) {
	auth, err := ethix.Authorize(PRIVATE_KEY)
	if err != nil {
		return nil, err
	}

	_, err = purpose.RemoveMission(auth, in.Mission.Title)
	if err != nil {
		return nil, err
	}

	return &pb.RemoveMissionResponse{}, nil
}

//
// Transform
//

func MissionsWebToPB(web3Missions []web3.Mission) []*pb.Mission {
	pbMissions := []*pb.Mission{}
	for _, mission := range web3Missions {
		pbMissions = append(pbMissions, &pb.Mission{
			Title:     mission.Title,
			Payout:    mission.Payout,
			Completed: mission.Completed,
		})
	}

	return pbMissions
}

func MissionWebToPB(mission *web3.Mission) *pb.Mission {
	return &pb.Mission{
		Title:     mission.Title,
		Payout:    mission.Payout,
		Completed: mission.Completed,
	}
}

//
// Tools
//

type PubSub struct {
	mutex       sync.Mutex
	connections map[pb.Purpose_MissionStreamServer]struct{}
}

func (ps *PubSub) Publish(message *pb.MissionStreamResponse) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	for connection := range ps.connections {
		if err := connection.Send(message); err != nil {
			logger.Warn(err, "unable to send message")
		}
	}
}

func (ps *PubSub) Subscribe(stream pb.Purpose_MissionStreamServer) (<-chan struct{}, error) {
	missions, err := purpose.AllMissions(nil)
	if err != nil {
		return nil, err
	}

	pbMissions := MissionsWebToPB(missions)
	err = stream.Send(&pb.MissionStreamResponse{
		Type:     "ALL",
		Missions: pbMissions,
	})
	if err != nil {
		return nil, err
	}

	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	ps.connections[stream] = struct{}{}
	return stream.Context().Done(), nil
}

func (ps *PubSub) Unsubscribe(stream pb.Purpose_MissionStreamServer) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	delete(ps.connections, stream)
}
