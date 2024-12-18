// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
	_ = abi.ConvertType
)

// ISignatureUtilsSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// ZenbtcbatcherMetaData contains all meta data concerning the Zenbtcbatcher contract.
var ZenbtcbatcherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zenBTC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rockBTC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_eglStrategyManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_eglDelegationManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_eglStrategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EGL_DELEGATION_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EGL_STRATEGY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EGL_STRATEGY_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROCK_BTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZEN_BTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fee\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"approverSalt\",\"type\":\"bytes32\"}],\"name\":\"wrapZenBTC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106101005760003560e01c80637ddaab5711610097578063b52e4e8c11610066578063b52e4e8c146102a8578063cf02d3d5146102bb578063d5391393146102e2578063d547741f1461030957600080fd5b80637ddaab571461023f57806391d1485414610266578063983d273714610279578063a217fddf146102a057600080fd5b806336568abe116100d357806336568abe1461019f578063443a8c23146101b2578063524d4227146101f157806364046b481461021857600080fd5b806301ffc9a7146101055780631626ba7e1461012d578063248a9ca3146101595780632f2ff15d1461018a575b600080fd5b610118610113366004610b48565b61031c565b60405190151581526020015b60405180910390f35b61014061013b366004610b8f565b610353565b6040516001600160e01b03199091168152602001610124565b61017c610167366004610c4a565b60009081526020819052604090206001015490565b604051908152602001610124565b61019d610198366004610c7f565b6103ba565b005b61019d6101ad366004610c7f565b6103e5565b6101d97f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610124565b6101d97f000000000000000000000000000000000000000000000000000000000000000081565b6101d97f000000000000000000000000000000000000000000000000000000000000000081565b6101d97f000000000000000000000000000000000000000000000000000000000000000081565b610118610274366004610c7f565b61041d565b6101d97f000000000000000000000000000000000000000000000000000000000000000081565b61017c600081565b61019d6102b6366004610cc3565b610446565b6101d97f000000000000000000000000000000000000000000000000000000000000000081565b61017c7fb458d13b0f4ce9e6aa65d297a27b10f75fdc6d0957bb29e1f2a30c8766b3541581565b61019d610317366004610c7f565b6107d8565b60006001600160e01b03198216637965db0b60e01b148061034d57506301ffc9a760e01b6001600160e01b03198316145b92915050565b60008061036084846107fd565b905061038c7fb458d13b0f4ce9e6aa65d297a27b10f75fdc6d0957bb29e1f2a30c8766b354158261041d565b156103a15750630b135d3f60e11b905061034d565b604051638baa579f60e01b815260040160405180910390fd5b6000828152602081905260409020600101546103d581610827565b6103df8383610834565b50505050565b6001600160a01b038116331461040e5760405163334bd91960e11b815260040160405180910390fd5b61041882826108c6565b505050565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b7fb458d13b0f4ce9e6aa65d297a27b10f75fdc6d0957bb29e1f2a30c8766b3541561047081610827565b6040516340c10f1960e01b815230600482015267ffffffffffffffff861660248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906340c10f1990604401600060405180830381600087803b1580156104e157600080fd5b505af11580156104f5573d6000803e3d6000fd5b505060405163095ea7b360e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116600483015267ffffffffffffffff891660248301527f000000000000000000000000000000000000000000000000000000000000000016925063095ea7b391506044016020604051808303816000875af1158015610591573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b59190610d3f565b506040516373d0285560e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301527f00000000000000000000000000000000000000000000000000000000000000008116602483015267ffffffffffffffff871660448301527f0000000000000000000000000000000000000000000000000000000000000000169063e7a050aa906064016020604051808303816000875af1158015610676573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069a9190610d61565b5060405163eea9064b60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eea9064b9061070b907f00000000000000000000000000000000000000000000000000000000000000009087908790600401610d7a565b600060405180830381600087803b15801561072557600080fd5b505af1158015610739573d6000803e3d6000fd5b505060405163f1d96dd360e01b81526001600160a01b03898116600483015267ffffffffffffffff808a166024840152881660448301527f000000000000000000000000000000000000000000000000000000000000000016925063f1d96dd39150606401600060405180830381600087803b1580156107b857600080fd5b505af11580156107cc573d6000803e3d6000fd5b50505050505050505050565b6000828152602081905260409020600101546107f381610827565b6103df83836108c6565b60008060008061080d8686610931565b92509250925061081d828261097e565b5090949350505050565b6108318133610a40565b50565b6000610840838361041d565b6108be576000838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556108763390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161034d565b50600061034d565b60006108d2838361041d565b156108be576000838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161034d565b6000806000835160410361096b5760208401516040850151606086015160001a61095d88828585610a79565b955095509550505050610977565b50508151600091506002905b9250925092565b600082600381111561099257610992610e16565b0361099b575050565b60018260038111156109af576109af610e16565b036109cd5760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156109e1576109e1610e16565b03610a075760405163fce698f760e01b8152600481018290526024015b60405180910390fd5b6003826003811115610a1b57610a1b610e16565b03610a3c576040516335e2f38360e21b8152600481018290526024016109fe565b5050565b610a4a828261041d565b610a3c5760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016109fe565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115610ab45750600091506003905082610b3e565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610b08573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610b3457506000925060019150829050610b3e565b9250600091508190505b9450945094915050565b600060208284031215610b5a57600080fd5b81356001600160e01b031981168114610b7257600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215610ba257600080fd5b82359150602083013567ffffffffffffffff80821115610bc157600080fd5b818501915085601f830112610bd557600080fd5b813581811115610be757610be7610b79565b604051601f8201601f19908116603f01168101908382118183101715610c0f57610c0f610b79565b81604052828152886020848701011115610c2857600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b600060208284031215610c5c57600080fd5b5035919050565b80356001600160a01b0381168114610c7a57600080fd5b919050565b60008060408385031215610c9257600080fd5b82359150610ca260208401610c63565b90509250929050565b803567ffffffffffffffff81168114610c7a57600080fd5b600080600080600060a08688031215610cdb57600080fd5b610ce486610c63565b9450610cf260208701610cab565b9350610d0060408701610cab565b9250606086013567ffffffffffffffff811115610d1c57600080fd5b860160408189031215610d2e57600080fd5b949793965091946080013592915050565b600060208284031215610d5157600080fd5b81518015158114610b7257600080fd5b600060208284031215610d7357600080fd5b5051919050565b6001600160a01b0384168152606060208201526000833536859003601e19018112610da457600080fd5b840160208101903567ffffffffffffffff811115610dc157600080fd5b803603821315610dd057600080fd5b604060608501528060a0850152808260c0860137600060c082860101526020860135608085015260c0601f19601f83011685010192505050826040830152949350505050565b634e487b7160e01b600052602160045260246000fdfea2646970667358221220a1de1a5a4b043740ce26414fd69f3cbb78bd06490271989241a140483d33a87c64736f6c63430008180033",
}

// ZenbtcbatcherABI is the input ABI used to generate the binding from.
// Deprecated: Use ZenbtcbatcherMetaData.ABI instead.
var ZenbtcbatcherABI = ZenbtcbatcherMetaData.ABI

// ZenbtcbatcherBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZenbtcbatcherMetaData.Bin instead.
var ZenbtcbatcherBin = ZenbtcbatcherMetaData.Bin

// DeployZenbtcbatcher deploys a new Ethereum contract, binding an instance of Zenbtcbatcher to it.
func DeployZenbtcbatcher(auth *bind.TransactOpts, backend bind.ContractBackend, _zenBTC common.Address, _rockBTC common.Address, _eglStrategyManager common.Address, _eglDelegationManager common.Address, _eglStrategy common.Address, _operator common.Address) (common.Address, *types.Transaction, *Zenbtcbatcher, error) {
	parsed, err := ZenbtcbatcherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZenbtcbatcherBin), backend, _zenBTC, _rockBTC, _eglStrategyManager, _eglDelegationManager, _eglStrategy, _operator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Zenbtcbatcher{ZenbtcbatcherCaller: ZenbtcbatcherCaller{contract: contract}, ZenbtcbatcherTransactor: ZenbtcbatcherTransactor{contract: contract}, ZenbtcbatcherFilterer: ZenbtcbatcherFilterer{contract: contract}}, nil
}

// Zenbtcbatcher is an auto generated Go binding around an Ethereum contract.
type Zenbtcbatcher struct {
	ZenbtcbatcherCaller     // Read-only binding to the contract
	ZenbtcbatcherTransactor // Write-only binding to the contract
	ZenbtcbatcherFilterer   // Log filterer for contract events
}

// ZenbtcbatcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZenbtcbatcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZenbtcbatcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZenbtcbatcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZenbtcbatcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZenbtcbatcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZenbtcbatcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZenbtcbatcherSession struct {
	Contract     *Zenbtcbatcher    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZenbtcbatcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZenbtcbatcherCallerSession struct {
	Contract *ZenbtcbatcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ZenbtcbatcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZenbtcbatcherTransactorSession struct {
	Contract     *ZenbtcbatcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ZenbtcbatcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZenbtcbatcherRaw struct {
	Contract *Zenbtcbatcher // Generic contract binding to access the raw methods on
}

// ZenbtcbatcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZenbtcbatcherCallerRaw struct {
	Contract *ZenbtcbatcherCaller // Generic read-only contract binding to access the raw methods on
}

// ZenbtcbatcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZenbtcbatcherTransactorRaw struct {
	Contract *ZenbtcbatcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZenbtcbatcher creates a new instance of Zenbtcbatcher, bound to a specific deployed contract.
func NewZenbtcbatcher(address common.Address, backend bind.ContractBackend) (*Zenbtcbatcher, error) {
	contract, err := bindZenbtcbatcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zenbtcbatcher{ZenbtcbatcherCaller: ZenbtcbatcherCaller{contract: contract}, ZenbtcbatcherTransactor: ZenbtcbatcherTransactor{contract: contract}, ZenbtcbatcherFilterer: ZenbtcbatcherFilterer{contract: contract}}, nil
}

// NewZenbtcbatcherCaller creates a new read-only instance of Zenbtcbatcher, bound to a specific deployed contract.
func NewZenbtcbatcherCaller(address common.Address, caller bind.ContractCaller) (*ZenbtcbatcherCaller, error) {
	contract, err := bindZenbtcbatcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZenbtcbatcherCaller{contract: contract}, nil
}

// NewZenbtcbatcherTransactor creates a new write-only instance of Zenbtcbatcher, bound to a specific deployed contract.
func NewZenbtcbatcherTransactor(address common.Address, transactor bind.ContractTransactor) (*ZenbtcbatcherTransactor, error) {
	contract, err := bindZenbtcbatcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZenbtcbatcherTransactor{contract: contract}, nil
}

// NewZenbtcbatcherFilterer creates a new log filterer instance of Zenbtcbatcher, bound to a specific deployed contract.
func NewZenbtcbatcherFilterer(address common.Address, filterer bind.ContractFilterer) (*ZenbtcbatcherFilterer, error) {
	contract, err := bindZenbtcbatcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZenbtcbatcherFilterer{contract: contract}, nil
}

// bindZenbtcbatcher binds a generic wrapper to an already deployed contract.
func bindZenbtcbatcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZenbtcbatcherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zenbtcbatcher *ZenbtcbatcherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zenbtcbatcher.Contract.ZenbtcbatcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zenbtcbatcher *ZenbtcbatcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zenbtcbatcher.Contract.ZenbtcbatcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zenbtcbatcher *ZenbtcbatcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zenbtcbatcher.Contract.ZenbtcbatcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zenbtcbatcher *ZenbtcbatcherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zenbtcbatcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zenbtcbatcher *ZenbtcbatcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zenbtcbatcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zenbtcbatcher *ZenbtcbatcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zenbtcbatcher.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Zenbtcbatcher *ZenbtcbatcherCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Zenbtcbatcher.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Zenbtcbatcher *ZenbtcbatcherSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Zenbtcbatcher.Contract.DEFAULTADMINROLE(&_Zenbtcbatcher.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Zenbtcbatcher *ZenbtcbatcherCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Zenbtcbatcher.Contract.DEFAULTADMINROLE(&_Zenbtcbatcher.CallOpts)
}

// EGLDELEGATIONMANAGER is a free data retrieval call binding the contract method 0xcf02d3d5.
//
// Solidity: function EGL_DELEGATION_MANAGER() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherCaller) EGLDELEGATIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zenbtcbatcher.contract.Call(opts, &out, "EGL_DELEGATION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EGLDELEGATIONMANAGER is a free data retrieval call binding the contract method 0xcf02d3d5.
//
// Solidity: function EGL_DELEGATION_MANAGER() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherSession) EGLDELEGATIONMANAGER() (common.Address, error) {
	return _Zenbtcbatcher.Contract.EGLDELEGATIONMANAGER(&_Zenbtcbatcher.CallOpts)
}

// EGLDELEGATIONMANAGER is a free data retrieval call binding the contract method 0xcf02d3d5.
//
// Solidity: function EGL_DELEGATION_MANAGER() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherCallerSession) EGLDELEGATIONMANAGER() (common.Address, error) {
	return _Zenbtcbatcher.Contract.EGLDELEGATIONMANAGER(&_Zenbtcbatcher.CallOpts)
}

// EGLSTRATEGY is a free data retrieval call binding the contract method 0x443a8c23.
//
// Solidity: function EGL_STRATEGY() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherCaller) EGLSTRATEGY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zenbtcbatcher.contract.Call(opts, &out, "EGL_STRATEGY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EGLSTRATEGY is a free data retrieval call binding the contract method 0x443a8c23.
//
// Solidity: function EGL_STRATEGY() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherSession) EGLSTRATEGY() (common.Address, error) {
	return _Zenbtcbatcher.Contract.EGLSTRATEGY(&_Zenbtcbatcher.CallOpts)
}

// EGLSTRATEGY is a free data retrieval call binding the contract method 0x443a8c23.
//
// Solidity: function EGL_STRATEGY() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherCallerSession) EGLSTRATEGY() (common.Address, error) {
	return _Zenbtcbatcher.Contract.EGLSTRATEGY(&_Zenbtcbatcher.CallOpts)
}

// EGLSTRATEGYMANAGER is a free data retrieval call binding the contract method 0x524d4227.
//
// Solidity: function EGL_STRATEGY_MANAGER() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherCaller) EGLSTRATEGYMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zenbtcbatcher.contract.Call(opts, &out, "EGL_STRATEGY_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EGLSTRATEGYMANAGER is a free data retrieval call binding the contract method 0x524d4227.
//
// Solidity: function EGL_STRATEGY_MANAGER() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherSession) EGLSTRATEGYMANAGER() (common.Address, error) {
	return _Zenbtcbatcher.Contract.EGLSTRATEGYMANAGER(&_Zenbtcbatcher.CallOpts)
}

// EGLSTRATEGYMANAGER is a free data retrieval call binding the contract method 0x524d4227.
//
// Solidity: function EGL_STRATEGY_MANAGER() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherCallerSession) EGLSTRATEGYMANAGER() (common.Address, error) {
	return _Zenbtcbatcher.Contract.EGLSTRATEGYMANAGER(&_Zenbtcbatcher.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Zenbtcbatcher *ZenbtcbatcherCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Zenbtcbatcher.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Zenbtcbatcher *ZenbtcbatcherSession) MINTERROLE() ([32]byte, error) {
	return _Zenbtcbatcher.Contract.MINTERROLE(&_Zenbtcbatcher.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Zenbtcbatcher *ZenbtcbatcherCallerSession) MINTERROLE() ([32]byte, error) {
	return _Zenbtcbatcher.Contract.MINTERROLE(&_Zenbtcbatcher.CallOpts)
}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherCaller) OPERATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zenbtcbatcher.contract.Call(opts, &out, "OPERATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherSession) OPERATOR() (common.Address, error) {
	return _Zenbtcbatcher.Contract.OPERATOR(&_Zenbtcbatcher.CallOpts)
}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherCallerSession) OPERATOR() (common.Address, error) {
	return _Zenbtcbatcher.Contract.OPERATOR(&_Zenbtcbatcher.CallOpts)
}

// ROCKBTC is a free data retrieval call binding the contract method 0x7ddaab57.
//
// Solidity: function ROCK_BTC() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherCaller) ROCKBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zenbtcbatcher.contract.Call(opts, &out, "ROCK_BTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ROCKBTC is a free data retrieval call binding the contract method 0x7ddaab57.
//
// Solidity: function ROCK_BTC() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherSession) ROCKBTC() (common.Address, error) {
	return _Zenbtcbatcher.Contract.ROCKBTC(&_Zenbtcbatcher.CallOpts)
}

// ROCKBTC is a free data retrieval call binding the contract method 0x7ddaab57.
//
// Solidity: function ROCK_BTC() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherCallerSession) ROCKBTC() (common.Address, error) {
	return _Zenbtcbatcher.Contract.ROCKBTC(&_Zenbtcbatcher.CallOpts)
}

// ZENBTC is a free data retrieval call binding the contract method 0x64046b48.
//
// Solidity: function ZEN_BTC() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherCaller) ZENBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Zenbtcbatcher.contract.Call(opts, &out, "ZEN_BTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZENBTC is a free data retrieval call binding the contract method 0x64046b48.
//
// Solidity: function ZEN_BTC() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherSession) ZENBTC() (common.Address, error) {
	return _Zenbtcbatcher.Contract.ZENBTC(&_Zenbtcbatcher.CallOpts)
}

// ZENBTC is a free data retrieval call binding the contract method 0x64046b48.
//
// Solidity: function ZEN_BTC() view returns(address)
func (_Zenbtcbatcher *ZenbtcbatcherCallerSession) ZENBTC() (common.Address, error) {
	return _Zenbtcbatcher.Contract.ZENBTC(&_Zenbtcbatcher.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Zenbtcbatcher *ZenbtcbatcherCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Zenbtcbatcher.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Zenbtcbatcher *ZenbtcbatcherSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Zenbtcbatcher.Contract.GetRoleAdmin(&_Zenbtcbatcher.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Zenbtcbatcher *ZenbtcbatcherCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Zenbtcbatcher.Contract.GetRoleAdmin(&_Zenbtcbatcher.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Zenbtcbatcher *ZenbtcbatcherCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Zenbtcbatcher.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Zenbtcbatcher *ZenbtcbatcherSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Zenbtcbatcher.Contract.HasRole(&_Zenbtcbatcher.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Zenbtcbatcher *ZenbtcbatcherCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Zenbtcbatcher.Contract.HasRole(&_Zenbtcbatcher.CallOpts, role, account)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4 magicValue)
func (_Zenbtcbatcher *ZenbtcbatcherCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _Zenbtcbatcher.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4 magicValue)
func (_Zenbtcbatcher *ZenbtcbatcherSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _Zenbtcbatcher.Contract.IsValidSignature(&_Zenbtcbatcher.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4 magicValue)
func (_Zenbtcbatcher *ZenbtcbatcherCallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _Zenbtcbatcher.Contract.IsValidSignature(&_Zenbtcbatcher.CallOpts, _hash, _signature)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Zenbtcbatcher *ZenbtcbatcherCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Zenbtcbatcher.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Zenbtcbatcher *ZenbtcbatcherSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Zenbtcbatcher.Contract.SupportsInterface(&_Zenbtcbatcher.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Zenbtcbatcher *ZenbtcbatcherCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Zenbtcbatcher.Contract.SupportsInterface(&_Zenbtcbatcher.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Zenbtcbatcher *ZenbtcbatcherTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Zenbtcbatcher.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Zenbtcbatcher *ZenbtcbatcherSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Zenbtcbatcher.Contract.GrantRole(&_Zenbtcbatcher.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Zenbtcbatcher *ZenbtcbatcherTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Zenbtcbatcher.Contract.GrantRole(&_Zenbtcbatcher.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Zenbtcbatcher *ZenbtcbatcherTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Zenbtcbatcher.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Zenbtcbatcher *ZenbtcbatcherSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Zenbtcbatcher.Contract.RenounceRole(&_Zenbtcbatcher.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Zenbtcbatcher *ZenbtcbatcherTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Zenbtcbatcher.Contract.RenounceRole(&_Zenbtcbatcher.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Zenbtcbatcher *ZenbtcbatcherTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Zenbtcbatcher.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Zenbtcbatcher *ZenbtcbatcherSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Zenbtcbatcher.Contract.RevokeRole(&_Zenbtcbatcher.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Zenbtcbatcher *ZenbtcbatcherTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Zenbtcbatcher.Contract.RevokeRole(&_Zenbtcbatcher.TransactOpts, role, account)
}

// WrapZenBTC is a paid mutator transaction binding the contract method 0xb52e4e8c.
//
// Solidity: function wrapZenBTC(address to, uint64 value, uint64 fee, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_Zenbtcbatcher *ZenbtcbatcherTransactor) WrapZenBTC(opts *bind.TransactOpts, to common.Address, value uint64, fee uint64, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _Zenbtcbatcher.contract.Transact(opts, "wrapZenBTC", to, value, fee, approverSignatureAndExpiry, approverSalt)
}

// WrapZenBTC is a paid mutator transaction binding the contract method 0xb52e4e8c.
//
// Solidity: function wrapZenBTC(address to, uint64 value, uint64 fee, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_Zenbtcbatcher *ZenbtcbatcherSession) WrapZenBTC(to common.Address, value uint64, fee uint64, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _Zenbtcbatcher.Contract.WrapZenBTC(&_Zenbtcbatcher.TransactOpts, to, value, fee, approverSignatureAndExpiry, approverSalt)
}

// WrapZenBTC is a paid mutator transaction binding the contract method 0xb52e4e8c.
//
// Solidity: function wrapZenBTC(address to, uint64 value, uint64 fee, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_Zenbtcbatcher *ZenbtcbatcherTransactorSession) WrapZenBTC(to common.Address, value uint64, fee uint64, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _Zenbtcbatcher.Contract.WrapZenBTC(&_Zenbtcbatcher.TransactOpts, to, value, fee, approverSignatureAndExpiry, approverSalt)
}

// ZenbtcbatcherRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Zenbtcbatcher contract.
type ZenbtcbatcherRoleAdminChangedIterator struct {
	Event *ZenbtcbatcherRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZenbtcbatcherRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenbtcbatcherRoleAdminChanged)
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
		it.Event = new(ZenbtcbatcherRoleAdminChanged)
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
func (it *ZenbtcbatcherRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenbtcbatcherRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenbtcbatcherRoleAdminChanged represents a RoleAdminChanged event raised by the Zenbtcbatcher contract.
type ZenbtcbatcherRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Zenbtcbatcher *ZenbtcbatcherFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZenbtcbatcherRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Zenbtcbatcher.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZenbtcbatcherRoleAdminChangedIterator{contract: _Zenbtcbatcher.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Zenbtcbatcher *ZenbtcbatcherFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZenbtcbatcherRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Zenbtcbatcher.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenbtcbatcherRoleAdminChanged)
				if err := _Zenbtcbatcher.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Zenbtcbatcher *ZenbtcbatcherFilterer) ParseRoleAdminChanged(log types.Log) (*ZenbtcbatcherRoleAdminChanged, error) {
	event := new(ZenbtcbatcherRoleAdminChanged)
	if err := _Zenbtcbatcher.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenbtcbatcherRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Zenbtcbatcher contract.
type ZenbtcbatcherRoleGrantedIterator struct {
	Event *ZenbtcbatcherRoleGranted // Event containing the contract specifics and raw log

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
func (it *ZenbtcbatcherRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenbtcbatcherRoleGranted)
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
		it.Event = new(ZenbtcbatcherRoleGranted)
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
func (it *ZenbtcbatcherRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenbtcbatcherRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenbtcbatcherRoleGranted represents a RoleGranted event raised by the Zenbtcbatcher contract.
type ZenbtcbatcherRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Zenbtcbatcher *ZenbtcbatcherFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZenbtcbatcherRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Zenbtcbatcher.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZenbtcbatcherRoleGrantedIterator{contract: _Zenbtcbatcher.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Zenbtcbatcher *ZenbtcbatcherFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZenbtcbatcherRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Zenbtcbatcher.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenbtcbatcherRoleGranted)
				if err := _Zenbtcbatcher.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Zenbtcbatcher *ZenbtcbatcherFilterer) ParseRoleGranted(log types.Log) (*ZenbtcbatcherRoleGranted, error) {
	event := new(ZenbtcbatcherRoleGranted)
	if err := _Zenbtcbatcher.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenbtcbatcherRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Zenbtcbatcher contract.
type ZenbtcbatcherRoleRevokedIterator struct {
	Event *ZenbtcbatcherRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ZenbtcbatcherRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenbtcbatcherRoleRevoked)
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
		it.Event = new(ZenbtcbatcherRoleRevoked)
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
func (it *ZenbtcbatcherRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenbtcbatcherRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenbtcbatcherRoleRevoked represents a RoleRevoked event raised by the Zenbtcbatcher contract.
type ZenbtcbatcherRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Zenbtcbatcher *ZenbtcbatcherFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZenbtcbatcherRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Zenbtcbatcher.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZenbtcbatcherRoleRevokedIterator{contract: _Zenbtcbatcher.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Zenbtcbatcher *ZenbtcbatcherFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZenbtcbatcherRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Zenbtcbatcher.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenbtcbatcherRoleRevoked)
				if err := _Zenbtcbatcher.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Zenbtcbatcher *ZenbtcbatcherFilterer) ParseRoleRevoked(log types.Log) (*ZenbtcbatcherRoleRevoked, error) {
	event := new(ZenbtcbatcherRoleRevoked)
	if err := _Zenbtcbatcher.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}