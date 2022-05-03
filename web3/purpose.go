// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package web3

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Mission is an auto generated low-level Go binding around an user-defined struct.
type Mission struct {
	Title     string
	Payout    uint64
	Completed bool
}

// MissionInterfaceMetaData contains all meta data concerning the MissionInterface contract.
var MissionInterfaceMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220aa026ee85e264ae0f32569cf68564235726afe06700143340a53b3dd9cb0728d64736f6c634300080d0033",
}

// MissionInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MissionInterfaceMetaData.ABI instead.
var MissionInterfaceABI = MissionInterfaceMetaData.ABI

// MissionInterfaceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MissionInterfaceMetaData.Bin instead.
var MissionInterfaceBin = MissionInterfaceMetaData.Bin

// DeployMissionInterface deploys a new Ethereum contract, binding an instance of MissionInterface to it.
func DeployMissionInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MissionInterface, error) {
	parsed, err := MissionInterfaceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MissionInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MissionInterface{MissionInterfaceCaller: MissionInterfaceCaller{contract: contract}, MissionInterfaceTransactor: MissionInterfaceTransactor{contract: contract}, MissionInterfaceFilterer: MissionInterfaceFilterer{contract: contract}}, nil
}

// MissionInterface is an auto generated Go binding around an Ethereum contract.
type MissionInterface struct {
	MissionInterfaceCaller     // Read-only binding to the contract
	MissionInterfaceTransactor // Write-only binding to the contract
	MissionInterfaceFilterer   // Log filterer for contract events
}

// MissionInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MissionInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MissionInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MissionInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MissionInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MissionInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MissionInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MissionInterfaceSession struct {
	Contract     *MissionInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MissionInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MissionInterfaceCallerSession struct {
	Contract *MissionInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MissionInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MissionInterfaceTransactorSession struct {
	Contract     *MissionInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MissionInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MissionInterfaceRaw struct {
	Contract *MissionInterface // Generic contract binding to access the raw methods on
}

// MissionInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MissionInterfaceCallerRaw struct {
	Contract *MissionInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// MissionInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MissionInterfaceTransactorRaw struct {
	Contract *MissionInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMissionInterface creates a new instance of MissionInterface, bound to a specific deployed contract.
func NewMissionInterface(address common.Address, backend bind.ContractBackend) (*MissionInterface, error) {
	contract, err := bindMissionInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MissionInterface{MissionInterfaceCaller: MissionInterfaceCaller{contract: contract}, MissionInterfaceTransactor: MissionInterfaceTransactor{contract: contract}, MissionInterfaceFilterer: MissionInterfaceFilterer{contract: contract}}, nil
}

// NewMissionInterfaceCaller creates a new read-only instance of MissionInterface, bound to a specific deployed contract.
func NewMissionInterfaceCaller(address common.Address, caller bind.ContractCaller) (*MissionInterfaceCaller, error) {
	contract, err := bindMissionInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MissionInterfaceCaller{contract: contract}, nil
}

// NewMissionInterfaceTransactor creates a new write-only instance of MissionInterface, bound to a specific deployed contract.
func NewMissionInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MissionInterfaceTransactor, error) {
	contract, err := bindMissionInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MissionInterfaceTransactor{contract: contract}, nil
}

// NewMissionInterfaceFilterer creates a new log filterer instance of MissionInterface, bound to a specific deployed contract.
func NewMissionInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MissionInterfaceFilterer, error) {
	contract, err := bindMissionInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MissionInterfaceFilterer{contract: contract}, nil
}

// bindMissionInterface binds a generic wrapper to an already deployed contract.
func bindMissionInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MissionInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MissionInterface *MissionInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MissionInterface.Contract.MissionInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MissionInterface *MissionInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MissionInterface.Contract.MissionInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MissionInterface *MissionInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MissionInterface.Contract.MissionInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MissionInterface *MissionInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MissionInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MissionInterface *MissionInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MissionInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MissionInterface *MissionInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MissionInterface.Contract.contract.Transact(opts, method, params...)
}

// PurposeMetaData contains all meta data concerning the Purpose contract.
var PurposeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"payout\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structMission\",\"name\":\"mission\",\"type\":\"tuple\"}],\"name\":\"AddedMission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"payout\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structMission\",\"name\":\"mission\",\"type\":\"tuple\"}],\"name\":\"RemovedMission\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"addMission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allMissions\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"payout\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"internalType\":\"structMission[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"missionByID\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"payout\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"internalType\":\"structMission\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"missionByName\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"payout\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"internalType\":\"structMission\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"removeMission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"72254ef8": "addMission(string)",
		"2da76955": "allMissions()",
		"f17fcb2f": "missionByID(uint256)",
		"7509d95d": "missionByName(string)",
		"d4b127ff": "removeMission(string)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610d79806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632da769551461005c57806372254ef81461007a5780637509d95d1461008f578063d4b127ff146100af578063f17fcb2f146100c2575b600080fd5b6100646100d5565b6040516100719190610b06565b60405180910390f35b61008d610088366004610b7e565b6100e6565b005b6100a261009d366004610b7e565b61013a565b6040516100719190610c2f565b61008d6100bd366004610b7e565b610168565b6100a26100d0366004610c49565b6101ae565b60606100e160006101d6565b905090565b6000806100f38184610469565b915091508015610135577fdba8f119d1a15584e96e4b7347aaca727202e849378f6f178a37a24cc6244ad58260405161012c9190610c2f565b60405180910390a15b505050565b604080516060808201835281526000602082018190529181019190915261016260008361061b565b92915050565b6000806101758184610724565b915091508015610135577f27e8d524e71df3d22630ccf6e2b48aa67b9dbcc29e37e3e23175e8045b7d141c8260405161012c9190610c2f565b60408051606080820183528152600060208201819052918101919091526101626000836108c1565b6060600080836002015467ffffffffffffffff8111156101f8576101f8610b68565b60405190808252806020026020018201604052801561024357816020015b60408051606080820183528152600060208083018290529282015282526000199092019101816102165790505b50905060005b846003015481101561046157600085600001828154811061026c5761026c610c62565b600091825260209182902060408051808201909152600290920201805460ff161515825260018101805492939192918401916102a790610c78565b80601f01602080910402602001604051908101604052809291908181526020018280546102d390610c78565b80156103205780601f106102f557610100808354040283529160200191610320565b820191906000526020600020905b81548152906001019060200180831161030357829003601f168201915b5050505050815250509050806000015161044e5785600101816020015160405161034a9190610cb2565b908152602001604051809103902060010160405180606001604052908160008201805461037690610c78565b80601f01602080910402602001604051908101604052809291908181526020018280546103a290610c78565b80156103ef5780601f106103c4576101008083540402835291602001916103ef565b820191906000526020600020905b8154815290600101906020018083116103d257829003601f168201915b50505091835250506001919091015467ffffffffffffffff81166020830152600160401b900460ff161515604090910152835184908690811061043457610434610c62565b6020026020010181905250838061044a90610ce4565b9450505b508061045981610ce4565b915050610249565b509392505050565b6040805160608082018352815260006020820181905291810191909152600080846001018460405161049b9190610cb2565b90815260405190819003602001902054905080156104d9576104cd6040518060200160405280600081525060006109a3565b60009250925050610614565b60006104e68560006109a3565b86546001018088556000889052600388015491925086918891811061050d5761050d610c62565b906000526020600020906002020160010190805190602001906105319291906109e9565b5060405180604001604052808760030154600161054e9190610cfd565b81526020018281525086600101866040516105699190610cb2565b9081526040516020918190038201902082518155828201518051805192939192600185019261059c9284929101906109e9565b506020820151600190910180546040909301511515600160401b0268ffffffffffffffffff1990931667ffffffffffffffff9092169190911791909117905550506002860180549060006105ef83610ce4565b909155505060038601805490600061060683610ce4565b909155509093506001925050505b9250929050565b604080516060808201835281526000602082018190528183015290516001840190610647908490610cb2565b908152602001604051809103902060010160405180606001604052908160008201805461067390610c78565b80601f016020809104026020016040519081016040528092919081815260200182805461069f90610c78565b80156106ec5780601f106106c1576101008083540402835291602001916106ec565b820191906000526020600020905b8154815290600101906020018083116106cf57829003601f168201915b50505091835250506001919091015467ffffffffffffffff81166020830152600160401b900460ff1615156040909101529392505050565b604080516060808201835281526000602082018190529181019190915260008084600101846040516107569190610cb2565b90815260200160405180910390206000015490508060000361078c576104cd6040518060200160405280600081525060006109a3565b600085600101856040516107a09190610cb2565b90815260200160405180910390206001016040518060600160405290816000820180546107cc90610c78565b80601f01602080910402602001604051908101604052809291908181526020018280546107f890610c78565b80156108455780601f1061081a57610100808354040283529160200191610845565b820191906000526020600020905b81548152906001019060200180831161082857829003601f168201915b505050918352505060019182015467ffffffffffffffff81166020830152600160401b900460ff161515604090910152909150866108838285610d15565b8154811061089357610893610c62565b600091825260208220600291820201805460ff19169315159390931790925590870180549161060683610d2c565b604080516060808201835281526000602082018190529181019190915260008360000183815481106108f5576108f5610c62565b9060005260206000209060020201600101805461091190610c78565b80601f016020809104026020016040519081016040528092919081815260200182805461093d90610c78565b801561098a5780601f1061095f5761010080835404028352916020019161098a565b820191906000526020600020905b81548152906001019060200180831161096d57829003601f168201915b5050505050905061099b848261061b565b949350505050565b6040805160608082018352815260006020820181905291810191909152506040805160608101825292835267ffffffffffffffff91909116602083015260009082015290565b8280546109f590610c78565b90600052602060002090601f016020900481019282610a175760008555610a5d565b82601f10610a3057805160ff1916838001178555610a5d565b82800160010185558215610a5d579182015b82811115610a5d578251825591602001919060010190610a42565b50610a69929150610a6d565b5090565b5b80821115610a695760008155600101610a6e565b60005b83811015610a9d578181015183820152602001610a85565b83811115610aac576000848401525b50505050565b60008151606084528051806060860152610ad3816080870160208501610a82565b60208481015167ffffffffffffffff16908601526040938401511515938501939093525050601f01601f19160160800190565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015610b5b57603f19888603018452610b49858351610ab2565b94509285019290850190600101610b2d565b5092979650505050505050565b634e487b7160e01b600052604160045260246000fd5b600060208284031215610b9057600080fd5b813567ffffffffffffffff80821115610ba857600080fd5b818401915084601f830112610bbc57600080fd5b813581811115610bce57610bce610b68565b604051601f8201601f19908116603f01168101908382118183101715610bf657610bf6610b68565b81604052828152876020848701011115610c0f57600080fd5b826020860160208301376000928101602001929092525095945050505050565b602081526000610c426020830184610ab2565b9392505050565b600060208284031215610c5b57600080fd5b5035919050565b634e487b7160e01b600052603260045260246000fd5b600181811c90821680610c8c57607f821691505b602082108103610cac57634e487b7160e01b600052602260045260246000fd5b50919050565b60008251610cc4818460208701610a82565b9190910192915050565b634e487b7160e01b600052601160045260246000fd5b600060018201610cf657610cf6610cce565b5060010190565b60008219821115610d1057610d10610cce565b500190565b600082821015610d2757610d27610cce565b500390565b600081610d3b57610d3b610cce565b50600019019056fea2646970667358221220a7b18211cd7c36133196283978dbc23225640147e7bd7659a121941aa5ac041b64736f6c634300080d0033",
}

// PurposeABI is the input ABI used to generate the binding from.
// Deprecated: Use PurposeMetaData.ABI instead.
var PurposeABI = PurposeMetaData.ABI

// Deprecated: Use PurposeMetaData.Sigs instead.
// PurposeFuncSigs maps the 4-byte function signature to its string representation.
var PurposeFuncSigs = PurposeMetaData.Sigs

// PurposeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PurposeMetaData.Bin instead.
var PurposeBin = PurposeMetaData.Bin

// DeployPurpose deploys a new Ethereum contract, binding an instance of Purpose to it.
func DeployPurpose(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Purpose, error) {
	parsed, err := PurposeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PurposeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Purpose{PurposeCaller: PurposeCaller{contract: contract}, PurposeTransactor: PurposeTransactor{contract: contract}, PurposeFilterer: PurposeFilterer{contract: contract}}, nil
}

// Purpose is an auto generated Go binding around an Ethereum contract.
type Purpose struct {
	PurposeCaller     // Read-only binding to the contract
	PurposeTransactor // Write-only binding to the contract
	PurposeFilterer   // Log filterer for contract events
}

// PurposeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PurposeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PurposeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PurposeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PurposeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PurposeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PurposeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PurposeSession struct {
	Contract     *Purpose          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PurposeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PurposeCallerSession struct {
	Contract *PurposeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PurposeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PurposeTransactorSession struct {
	Contract     *PurposeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PurposeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PurposeRaw struct {
	Contract *Purpose // Generic contract binding to access the raw methods on
}

// PurposeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PurposeCallerRaw struct {
	Contract *PurposeCaller // Generic read-only contract binding to access the raw methods on
}

// PurposeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PurposeTransactorRaw struct {
	Contract *PurposeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPurpose creates a new instance of Purpose, bound to a specific deployed contract.
func NewPurpose(address common.Address, backend bind.ContractBackend) (*Purpose, error) {
	contract, err := bindPurpose(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Purpose{PurposeCaller: PurposeCaller{contract: contract}, PurposeTransactor: PurposeTransactor{contract: contract}, PurposeFilterer: PurposeFilterer{contract: contract}}, nil
}

// NewPurposeCaller creates a new read-only instance of Purpose, bound to a specific deployed contract.
func NewPurposeCaller(address common.Address, caller bind.ContractCaller) (*PurposeCaller, error) {
	contract, err := bindPurpose(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PurposeCaller{contract: contract}, nil
}

// NewPurposeTransactor creates a new write-only instance of Purpose, bound to a specific deployed contract.
func NewPurposeTransactor(address common.Address, transactor bind.ContractTransactor) (*PurposeTransactor, error) {
	contract, err := bindPurpose(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PurposeTransactor{contract: contract}, nil
}

// NewPurposeFilterer creates a new log filterer instance of Purpose, bound to a specific deployed contract.
func NewPurposeFilterer(address common.Address, filterer bind.ContractFilterer) (*PurposeFilterer, error) {
	contract, err := bindPurpose(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PurposeFilterer{contract: contract}, nil
}

// bindPurpose binds a generic wrapper to an already deployed contract.
func bindPurpose(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PurposeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Purpose *PurposeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Purpose.Contract.PurposeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Purpose *PurposeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Purpose.Contract.PurposeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Purpose *PurposeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Purpose.Contract.PurposeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Purpose *PurposeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Purpose.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Purpose *PurposeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Purpose.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Purpose *PurposeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Purpose.Contract.contract.Transact(opts, method, params...)
}

// AllMissions is a free data retrieval call binding the contract method 0x2da76955.
//
// Solidity: function allMissions() view returns((string,uint64,bool)[])
func (_Purpose *PurposeCaller) AllMissions(opts *bind.CallOpts) ([]Mission, error) {
	var out []interface{}
	err := _Purpose.contract.Call(opts, &out, "allMissions")

	if err != nil {
		return *new([]Mission), err
	}

	out0 := *abi.ConvertType(out[0], new([]Mission)).(*[]Mission)

	return out0, err

}

// AllMissions is a free data retrieval call binding the contract method 0x2da76955.
//
// Solidity: function allMissions() view returns((string,uint64,bool)[])
func (_Purpose *PurposeSession) AllMissions() ([]Mission, error) {
	return _Purpose.Contract.AllMissions(&_Purpose.CallOpts)
}

// AllMissions is a free data retrieval call binding the contract method 0x2da76955.
//
// Solidity: function allMissions() view returns((string,uint64,bool)[])
func (_Purpose *PurposeCallerSession) AllMissions() ([]Mission, error) {
	return _Purpose.Contract.AllMissions(&_Purpose.CallOpts)
}

// MissionByID is a free data retrieval call binding the contract method 0xf17fcb2f.
//
// Solidity: function missionByID(uint256 id) view returns((string,uint64,bool))
func (_Purpose *PurposeCaller) MissionByID(opts *bind.CallOpts, id *big.Int) (Mission, error) {
	var out []interface{}
	err := _Purpose.contract.Call(opts, &out, "missionByID", id)

	if err != nil {
		return *new(Mission), err
	}

	out0 := *abi.ConvertType(out[0], new(Mission)).(*Mission)

	return out0, err

}

// MissionByID is a free data retrieval call binding the contract method 0xf17fcb2f.
//
// Solidity: function missionByID(uint256 id) view returns((string,uint64,bool))
func (_Purpose *PurposeSession) MissionByID(id *big.Int) (Mission, error) {
	return _Purpose.Contract.MissionByID(&_Purpose.CallOpts, id)
}

// MissionByID is a free data retrieval call binding the contract method 0xf17fcb2f.
//
// Solidity: function missionByID(uint256 id) view returns((string,uint64,bool))
func (_Purpose *PurposeCallerSession) MissionByID(id *big.Int) (Mission, error) {
	return _Purpose.Contract.MissionByID(&_Purpose.CallOpts, id)
}

// MissionByName is a free data retrieval call binding the contract method 0x7509d95d.
//
// Solidity: function missionByName(string name) view returns((string,uint64,bool))
func (_Purpose *PurposeCaller) MissionByName(opts *bind.CallOpts, name string) (Mission, error) {
	var out []interface{}
	err := _Purpose.contract.Call(opts, &out, "missionByName", name)

	if err != nil {
		return *new(Mission), err
	}

	out0 := *abi.ConvertType(out[0], new(Mission)).(*Mission)

	return out0, err

}

// MissionByName is a free data retrieval call binding the contract method 0x7509d95d.
//
// Solidity: function missionByName(string name) view returns((string,uint64,bool))
func (_Purpose *PurposeSession) MissionByName(name string) (Mission, error) {
	return _Purpose.Contract.MissionByName(&_Purpose.CallOpts, name)
}

// MissionByName is a free data retrieval call binding the contract method 0x7509d95d.
//
// Solidity: function missionByName(string name) view returns((string,uint64,bool))
func (_Purpose *PurposeCallerSession) MissionByName(name string) (Mission, error) {
	return _Purpose.Contract.MissionByName(&_Purpose.CallOpts, name)
}

// AddMission is a paid mutator transaction binding the contract method 0x72254ef8.
//
// Solidity: function addMission(string name) returns()
func (_Purpose *PurposeTransactor) AddMission(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Purpose.contract.Transact(opts, "addMission", name)
}

// AddMission is a paid mutator transaction binding the contract method 0x72254ef8.
//
// Solidity: function addMission(string name) returns()
func (_Purpose *PurposeSession) AddMission(name string) (*types.Transaction, error) {
	return _Purpose.Contract.AddMission(&_Purpose.TransactOpts, name)
}

// AddMission is a paid mutator transaction binding the contract method 0x72254ef8.
//
// Solidity: function addMission(string name) returns()
func (_Purpose *PurposeTransactorSession) AddMission(name string) (*types.Transaction, error) {
	return _Purpose.Contract.AddMission(&_Purpose.TransactOpts, name)
}

// RemoveMission is a paid mutator transaction binding the contract method 0xd4b127ff.
//
// Solidity: function removeMission(string name) returns()
func (_Purpose *PurposeTransactor) RemoveMission(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Purpose.contract.Transact(opts, "removeMission", name)
}

// RemoveMission is a paid mutator transaction binding the contract method 0xd4b127ff.
//
// Solidity: function removeMission(string name) returns()
func (_Purpose *PurposeSession) RemoveMission(name string) (*types.Transaction, error) {
	return _Purpose.Contract.RemoveMission(&_Purpose.TransactOpts, name)
}

// RemoveMission is a paid mutator transaction binding the contract method 0xd4b127ff.
//
// Solidity: function removeMission(string name) returns()
func (_Purpose *PurposeTransactorSession) RemoveMission(name string) (*types.Transaction, error) {
	return _Purpose.Contract.RemoveMission(&_Purpose.TransactOpts, name)
}

// PurposeAddedMissionIterator is returned from FilterAddedMission and is used to iterate over the raw logs and unpacked data for AddedMission events raised by the Purpose contract.
type PurposeAddedMissionIterator struct {
	Event *PurposeAddedMission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PurposeAddedMissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PurposeAddedMission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PurposeAddedMission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PurposeAddedMissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PurposeAddedMissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PurposeAddedMission represents a AddedMission event raised by the Purpose contract.
type PurposeAddedMission struct {
	Mission Mission
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddedMission is a free log retrieval operation binding the contract event 0xdba8f119d1a15584e96e4b7347aaca727202e849378f6f178a37a24cc6244ad5.
//
// Solidity: event AddedMission((string,uint64,bool) mission)
func (_Purpose *PurposeFilterer) FilterAddedMission(opts *bind.FilterOpts) (*PurposeAddedMissionIterator, error) {

	logs, sub, err := _Purpose.contract.FilterLogs(opts, "AddedMission")
	if err != nil {
		return nil, err
	}
	return &PurposeAddedMissionIterator{contract: _Purpose.contract, event: "AddedMission", logs: logs, sub: sub}, nil
}

// WatchAddedMission is a free log subscription operation binding the contract event 0xdba8f119d1a15584e96e4b7347aaca727202e849378f6f178a37a24cc6244ad5.
//
// Solidity: event AddedMission((string,uint64,bool) mission)
func (_Purpose *PurposeFilterer) WatchAddedMission(opts *bind.WatchOpts, sink chan<- *PurposeAddedMission) (event.Subscription, error) {

	logs, sub, err := _Purpose.contract.WatchLogs(opts, "AddedMission")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PurposeAddedMission)
				if err := _Purpose.contract.UnpackLog(event, "AddedMission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddedMission is a log parse operation binding the contract event 0xdba8f119d1a15584e96e4b7347aaca727202e849378f6f178a37a24cc6244ad5.
//
// Solidity: event AddedMission((string,uint64,bool) mission)
func (_Purpose *PurposeFilterer) ParseAddedMission(log types.Log) (*PurposeAddedMission, error) {
	event := new(PurposeAddedMission)
	if err := _Purpose.contract.UnpackLog(event, "AddedMission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PurposeRemovedMissionIterator is returned from FilterRemovedMission and is used to iterate over the raw logs and unpacked data for RemovedMission events raised by the Purpose contract.
type PurposeRemovedMissionIterator struct {
	Event *PurposeRemovedMission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PurposeRemovedMissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PurposeRemovedMission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PurposeRemovedMission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PurposeRemovedMissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PurposeRemovedMissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PurposeRemovedMission represents a RemovedMission event raised by the Purpose contract.
type PurposeRemovedMission struct {
	Mission Mission
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemovedMission is a free log retrieval operation binding the contract event 0x27e8d524e71df3d22630ccf6e2b48aa67b9dbcc29e37e3e23175e8045b7d141c.
//
// Solidity: event RemovedMission((string,uint64,bool) mission)
func (_Purpose *PurposeFilterer) FilterRemovedMission(opts *bind.FilterOpts) (*PurposeRemovedMissionIterator, error) {

	logs, sub, err := _Purpose.contract.FilterLogs(opts, "RemovedMission")
	if err != nil {
		return nil, err
	}
	return &PurposeRemovedMissionIterator{contract: _Purpose.contract, event: "RemovedMission", logs: logs, sub: sub}, nil
}

// WatchRemovedMission is a free log subscription operation binding the contract event 0x27e8d524e71df3d22630ccf6e2b48aa67b9dbcc29e37e3e23175e8045b7d141c.
//
// Solidity: event RemovedMission((string,uint64,bool) mission)
func (_Purpose *PurposeFilterer) WatchRemovedMission(opts *bind.WatchOpts, sink chan<- *PurposeRemovedMission) (event.Subscription, error) {

	logs, sub, err := _Purpose.contract.WatchLogs(opts, "RemovedMission")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PurposeRemovedMission)
				if err := _Purpose.contract.UnpackLog(event, "RemovedMission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemovedMission is a log parse operation binding the contract event 0x27e8d524e71df3d22630ccf6e2b48aa67b9dbcc29e37e3e23175e8045b7d141c.
//
// Solidity: event RemovedMission((string,uint64,bool) mission)
func (_Purpose *PurposeFilterer) ParseRemovedMission(log types.Log) (*PurposeRemovedMission, error) {
	event := new(PurposeRemovedMission)
	if err := _Purpose.contract.UnpackLog(event, "RemovedMission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
