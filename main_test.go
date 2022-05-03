package main

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/adamavixio/purpose-backend/goethix"
	"github.com/adamavixio/purpose-backend/web3"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ADDRESS = "http://127.0.0.1:7545"

	ACC_1_PRIVATE = "84bf34442cd7d3dee01465398716034734aa445ac1fdef3dd4836cea6b5f4e37"
	ACC_1_ADDRESS = "0xbA3A2950a2178153716b8f1917992746AA166A31"

	ACC_2_PRIVATE = "15699e8369f6c1cb3019aa5119a6fd8491a2e259dc32c1fc00b1bbd018876c1b"
	ACC_2_ADDRESS = "0xe40bBa4959E9CB1bb614eC2228B15b7c554bc44A"

	TEST_CONTRACT = "0x5a0994de6d4596d2867A481feF494B447655a6B3"
)

func TestNewContract(t *testing.T) {
	ethix = goethix.NewEthix()

	err := ethix.Dial(ADDRESS)
	if err != nil {
		t.Fatal(err, "unable to dial ethereum node")
	}

	auth, err := ethix.Authorize(ACC_1_PRIVATE)
	if err != nil {
		t.Fatal(err, "unable to dial ethereum node")
	}

	address, _, _, err := web3.DeployPurpose(auth, ethix.Client())
	if err != nil {
		t.Fatal(err, "unable to dial ethereum node")
	}

	t.Log(address)
}

func TestAddMission(t *testing.T) {
	ethix = goethix.NewEthix()

	err := ethix.Dial(ADDRESS)
	if err != nil {
		t.Fatal(err, "unable to dial ethereum node")
	}

	contract := common.HexToAddress(TEST_CONTRACT)
	purpose, err = web3.NewPurpose(contract, ethix.Client())
	if err != nil {
		t.Fatal(err, "unable to create purpose web3 client")
	}

	auth, err := ethix.Authorize(ACC_1_PRIVATE)
	if err != nil {
		t.Fatal(err, "unable to create purpose web3 client")
	}

	_, err = purpose.AddMission(auth, "test"+strconv.Itoa(rand.Intn(100)))
	if err != nil {
		t.Fatal(err, "unable to add mission")
	}
}
