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

// IDelegationManagerQueuedWithdrawalParams is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerQueuedWithdrawalParams struct {
	Strategies []common.Address
	Shares     []*big.Int
	Withdrawer common.Address
}

// IDelegationManagerWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerWithdrawal struct {
	Staker      common.Address
	DelegatedTo common.Address
	Withdrawer  common.Address
	Nonce       *big.Int
	StartBlock  uint32
	Strategies  []common.Address
	Shares      []*big.Int
}

// ISignatureUtilsSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// ZenBTControllerMetaData contains all meta data concerning the ZenBTController contract.
var ZenBTControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"RockBTCStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startBlock\",\"type\":\"uint32\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structIDelegationManager.Withdrawal\",\"name\":\"\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"receiveAsTokens\",\"type\":\"bool\"}],\"name\":\"WithdrawalCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"withdrawalIds\",\"type\":\"bytes32[]\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fee\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ZenBTCWrapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEGLDelegationManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEGLStrategyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRockBTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRockBTCStrategy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getZenBTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zenBTC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rockBTC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_eglStrategyManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_eglDelegationManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rockBTCStrategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"initializeV1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"approverSalt\",\"type\":\"bytes32\"}],\"name\":\"stakeRockBTC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startBlock\",\"type\":\"uint32\"},{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"}],\"internalType\":\"structIDelegationManager.Withdrawal\",\"name\":\"withdrawal\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"middlewareTimesIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"receiveAsTokens\",\"type\":\"bool\"}],\"name\":\"unwrapComplete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIStrategy[]\",\"name\":\"strategies\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"}],\"internalType\":\"structIDelegationManager.QueuedWithdrawalParams[]\",\"name\":\"queuedWithdrawalParams\",\"type\":\"tuple[]\"}],\"name\":\"unwrapInit\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"withdrawalIds\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fee\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"approverSalt\",\"type\":\"bytes32\"}],\"name\":\"wrapZenBTC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106101375760003560e01c8063931432ab116100b8578063d53913931161007c578063d5391393146102f7578063d547741f1461030c578063e2d1249a1461031f578063e7f43c681461033f578063f7d664941461036f578063fd26fc661461038d57600080fd5b8063931432ab146102915780639915c15f146102c1578063a217fddf146102d4578063b52e4e8c146102dc578063d087d288146102ef57600080fd5b80632f2ff15d116100ff5780632f2ff15d1461022557806336568abe1461023a5780633725d2fc1461024d5780636296ae271461026b57806391d148541461027e57600080fd5b806301ffc9a71461013c57806304611b06146101645780631626ba7e146101a85780631f05b551146101d4578063248a9ca314610204575b600080fd5b61014f61014a3660046111b3565b6103a0565b60405190151581526020015b60405180910390f35b7f1f7056f62fdbe84d13291834ab6e74f080827e3f8f77bb615d9db2295a18b501546001600160a01b03165b6040516001600160a01b03909116815260200161015b565b6101bb6101b636600461122a565b6103d7565b6040516001600160e01b0319909116815260200161015b565b7f1f7056f62fdbe84d13291834ab6e74f080827e3f8f77bb615d9db2295a18b503546001600160a01b0316610190565b6102176102123660046112c8565b61042c565b60405190815260200161015b565b6102386102333660046112f6565b61044e565b005b6102386102483660046112f6565b610470565b600080516020611acb833981519152546001600160a01b0316610190565b61023861027936600461135a565b6104a8565b61014f61028c3660046112f6565b6104cb565b7f1f7056f62fdbe84d13291834ab6e74f080827e3f8f77bb615d9db2295a18b505546001600160a01b0316610190565b6102386102cf366004611409565b610503565b610217600081565b6102386102ea366004611498565b61066e565b610217610775565b610217600080516020611b2b83398151915281565b61023861031a3660046112f6565b610801565b61033261032d366004611510565b61081d565b60405161015b9190611551565b7f1f7056f62fdbe84d13291834ab6e74f080827e3f8f77bb615d9db2295a18b502546001600160a01b0316610190565b600080516020611aeb833981519152546001600160a01b0316610190565b61023861039b366004611595565b61090c565b60006001600160e01b03198216637965db0b60e01b14806103d157506301ffc9a760e01b6001600160e01b03198316145b92915050565b6000806103e48484610a2c565b90506103fe600080516020611b2b833981519152826104cb565b156104135750630b135d3f60e11b90506103d1565b604051638baa579f60e01b815260040160405180910390fd5b6000908152600080516020611b0b833981519152602052604090206001015490565b6104578261042c565b61046081610a56565b61046a8383610a63565b50505050565b6001600160a01b03811633146104995760405163334bd91960e11b815260040160405180910390fd5b6104a38282610b08565b505050565b600080516020611b2b8339815191526104c081610a56565b61046a848484610b84565b6000918252600080516020611b0b833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b600080516020611b2b83398151915261051b81610a56565b600080516020611acb833981519152546040516360d7faed60e01b8152600080516020611aeb833981519152916001600160a01b0316906360d7faed9061056e908a908a908a908a908a9060040161179b565b600060405180830381600087803b15801561058857600080fd5b505af115801561059c573d6000803e3d6000fd5b5050505060018101546001600160a01b03166342966c686105c060c08a018a6117db565b60008181106105d1576105d1611824565b905060200201356040518263ffffffff1660e01b81526004016105f691815260200190565b600060405180830381600087803b15801561061057600080fd5b505af1158015610624573d6000803e3d6000fd5b505050507fadd5e0a5a1ec72776e3a7800fa2aec5ca0f7865cf68592ff037e468cb272a0a28787878660405161065d949392919061183a565b60405180910390a150505050505050565b600080516020611b2b83398151915261068681610a56565b600080516020611aeb8339815191526106a0868585610b84565b805460405163f1d96dd360e01b81526001600160a01b0389811660048301526001600160401b03808a166024840152881660448301529091169063f1d96dd390606401600060405180830381600087803b1580156106fd57600080fd5b505af1158015610711573d6000803e3d6000fd5b505050506002810154604080516001600160401b03808a168252881660208201526001600160a01b03928316928a16917ffd60f0aa3e690ac3a0ef44fb08201a6d258f27749c174539d4d1cf9c2ca0f8d7910160405180910390a350505050505050565b600080600080516020611aeb8339815191526004818101546040516329c77d4f60e01b815230928101929092529192506001600160a01b03909116906329c77d4f90602401602060405180830381865afa1580156107d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107fb9190611874565b91505090565b61080a8261042c565b61081381610a56565b61046a8383610b08565b6060600080516020611b2b83398151915261083781610a56565b600080516020611acb833981519152546040516306ec6e8160e11b8152600080516020611aeb833981519152916001600160a01b031690630dd8dd0290610884908890889060040161188d565b6000604051808303816000875af11580156108a3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526108cb9190810190611957565b92507f3ca8ee7ddc9cba27865c2837ec26df7a102ce7b8f538096db5b5d42e863502af836040516108fc9190611551565b60405180910390a1505092915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156109515750825b90506000826001600160401b0316600114801561096d5750303b155b90508115801561097b575080155b156109995760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156109c357845460ff60401b1916600160401b1785555b6109cb610ddc565b6109d98b8b8b8b8b8b610df6565b8315610a1f57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b600080600080610a3c8686610e1c565b925092509250610a4c8282610e69565b5090949350505050565b610a608133610f2b565b50565b6000600080516020611b0b833981519152610a7e84846104cb565b610afe576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055610ab43390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506103d1565b60009150506103d1565b6000600080516020611b0b833981519152610b2384846104cb565b15610afe576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506103d1565b6000600080516020611aeb83398151915260018101546040516340c10f1960e01b81523060048201526001600160401b03871660248201529192506001600160a01b0316906340c10f1990604401600060405180830381600087803b158015610bec57600080fd5b505af1158015610c00573d6000803e3d6000fd5b505050506001810154600382015460405163095ea7b360e01b81526001600160a01b0391821660048201526001600160401b038716602482015291169063095ea7b3906044016020604051808303816000875af1158015610c65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c8991906119fc565b506003810154600582015460018301546040516373d0285560e11b81526001600160a01b03928316600482015290821660248201526001600160401b038716604482015291169063e7a050aa906064016020604051808303816000875af1158015610cf8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1c9190611874565b50600480820154600283015460405163eea9064b60e01b81526001600160a01b039283169363eea9064b93610d579316918891889101611a19565b600060405180830381600087803b158015610d7157600080fd5b505af1158015610d85573d6000803e3d6000fd5b5050505060028101546040516001600160401b03861681526001600160a01b03909116907f79c412fdc602b15f36aa9b42405335e8dac8e322a516e71a0a3881d5ed41e3bc9060200160405180910390a250505050565b610de4610f64565b610dec610fad565b610df4610fbd565b565b610dfe610f64565b610e06610ddc565b610e14868686868686610fc5565b505050505050565b60008060008351604103610e565760208401516040850151606086015160001a610e48888285856110d1565b955095509550505050610e62565b50508151600091506002905b9250925092565b6000826003811115610e7d57610e7d611ab4565b03610e86575050565b6001826003811115610e9a57610e9a611ab4565b03610eb85760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115610ecc57610ecc611ab4565b03610ef25760405163fce698f760e01b8152600481018290526024015b60405180910390fd5b6003826003811115610f0657610f06611ab4565b03610f27576040516335e2f38360e21b815260048101829052602401610ee9565b5050565b610f3582826104cb565b610f275760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610ee9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610df457604051631afcd79f60e31b815260040160405180910390fd5b610fb5610f64565b610df46111a0565b610df4610f64565b610fcd610f64565b600080516020611aeb83398151915280546001600160a01b03199081166001600160a01b03988916179091557f1f7056f62fdbe84d13291834ab6e74f080827e3f8f77bb615d9db2295a18b50180548216968816969096179095557f1f7056f62fdbe84d13291834ab6e74f080827e3f8f77bb615d9db2295a18b502805486169187169190911790557f1f7056f62fdbe84d13291834ab6e74f080827e3f8f77bb615d9db2295a18b5038054851693861693909317909255600080516020611acb833981519152805484169185169190911790557f1f7056f62fdbe84d13291834ab6e74f080827e3f8f77bb615d9db2295a18b50580549092169216919091179055565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561110c5750600091506003905082611196565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611160573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661118c57506000925060019150829050611196565b9250600091508190505b9450945094915050565b6111a8610f64565b610a60600033610a63565b6000602082840312156111c557600080fd5b81356001600160e01b0319811681146111dd57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715611222576112226111e4565b604052919050565b6000806040838503121561123d57600080fd5b823591506020808401356001600160401b038082111561125c57600080fd5b818601915086601f83011261127057600080fd5b813581811115611282576112826111e4565b611294601f8201601f191685016111fa565b915080825287848285010111156112aa57600080fd5b80848401858401376000848284010152508093505050509250929050565b6000602082840312156112da57600080fd5b5035919050565b6001600160a01b0381168114610a6057600080fd5b6000806040838503121561130957600080fd5b82359150602083013561131b816112e1565b809150509250929050565b80356001600160401b038116811461133d57600080fd5b919050565b60006040828403121561135457600080fd5b50919050565b60008060006060848603121561136f57600080fd5b61137884611326565b925060208401356001600160401b0381111561139357600080fd5b61139f86828701611342565b925050604084013590509250925092565b60008083601f8401126113c257600080fd5b5081356001600160401b038111156113d957600080fd5b6020830191508360208260051b85010111156113f457600080fd5b9250929050565b8015158114610a6057600080fd5b60008060008060006080868803121561142157600080fd5b85356001600160401b038082111561143857600080fd5b9087019060e0828a03121561144c57600080fd5b9095506020870135908082111561146257600080fd5b5061146f888289016113b0565b90955093505060408601359150606086013561148a816113fb565b809150509295509295909350565b600080600080600060a086880312156114b057600080fd5b85356114bb816112e1565b94506114c960208701611326565b93506114d760408701611326565b925060608601356001600160401b038111156114f257600080fd5b6114fe88828901611342565b95989497509295608001359392505050565b6000806020838503121561152357600080fd5b82356001600160401b0381111561153957600080fd5b611545858286016113b0565b90969095509350505050565b6020808252825182820181905260009190848201906040850190845b818110156115895783518352928401929184019160010161156d565b50909695505050505050565b60008060008060008060c087890312156115ae57600080fd5b86356115b9816112e1565b955060208701356115c9816112e1565b945060408701356115d9816112e1565b935060608701356115e9816112e1565b925060808701356115f9816112e1565b915060a0870135611609816112e1565b809150509295509295509295565b6000808335601e1984360301811261162e57600080fd5b83016020810192503590506001600160401b0381111561164d57600080fd5b8060051b36038213156113f457600080fd5b8183526000602080850194508260005b8581101561169d578135611682816112e1565b6001600160a01b03168752958201959082019060010161166f565b509495945050505050565b81835260006001600160fb1b038311156116c157600080fd5b8260051b80836020870137939093016020019392505050565b600081356116e7816112e1565b6001600160a01b039081168452602083013590611703826112e1565b908116602085015260408301359061171a826112e1565b16604084015260608281013590840152608082013563ffffffff8116811461174157600080fd5b63ffffffff16608084015261175960a0830183611617565b60e060a086015261176e60e08601828461165f565b91505061177e60c0840184611617565b85830360c08701526117918382846116a8565b9695505050505050565b6080815260006117ae60808301886116da565b82810360208401526117c181878961165f565b604084019590955250509015156060909101529392505050565b6000808335601e198436030181126117f257600080fd5b8301803591506001600160401b0382111561180c57600080fd5b6020019150600581901b36038213156113f457600080fd5b634e487b7160e01b600052603260045260246000fd5b60608152600061184d60608301876116da565b828103602084015261186081868861165f565b915050821515604083015295945050505050565b60006020828403121561188657600080fd5b5051919050565b60208082528181018390526000906040808401600586901b850182018785805b8981101561194857888403603f190185528235368c9003605e190181126118d2578283fd5b8b0160606118e08280611617565b8288526118f0838901828461165f565b9250505061190089830183611617565b8783038b8901526119128382846116a8565b92505050878201359150611925826112e1565b6001600160a01b03919091169487019490945293860193918601916001016118ad565b50919998505050505050505050565b6000602080838503121561196a57600080fd5b82516001600160401b038082111561198157600080fd5b818501915085601f83011261199557600080fd5b8151818111156119a7576119a76111e4565b8060051b91506119b88483016111fa565b81815291830184019184810190888411156119d257600080fd5b938501935b838510156119f0578451825293850193908501906119d7565b98975050505050505050565b600060208284031215611a0e57600080fd5b81516111dd816113fb565b6001600160a01b0384168152606060208201526000833536859003601e19018112611a4357600080fd5b84016020810190356001600160401b03811115611a5f57600080fd5b803603821315611a6e57600080fd5b604060608501528060a0850152808260c0860137600060c082860101526020860135608085015260c0601f19601f83011685010192505050826040830152949350505050565b634e487b7160e01b600052602160045260246000fdfe1f7056f62fdbe84d13291834ab6e74f080827e3f8f77bb615d9db2295a18b5041f7056f62fdbe84d13291834ab6e74f080827e3f8f77bb615d9db2295a18b50002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800b458d13b0f4ce9e6aa65d297a27b10f75fdc6d0957bb29e1f2a30c8766b35415a2646970667358221220fa04ee4c26de0e58269ce720f14bf201e6fcb5bf0ae1aa128c1a452aa8c8e2e164736f6c63430008180033",
}

// ZenBTControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ZenBTControllerMetaData.ABI instead.
var ZenBTControllerABI = ZenBTControllerMetaData.ABI

// ZenBTControllerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZenBTControllerMetaData.Bin instead.
var ZenBTControllerBin = ZenBTControllerMetaData.Bin

// DeployZenBTController deploys a new Ethereum contract, binding an instance of ZenBTController to it.
func DeployZenBTController(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZenBTController, error) {
	parsed, err := ZenBTControllerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZenBTControllerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZenBTController{ZenBTControllerCaller: ZenBTControllerCaller{contract: contract}, ZenBTControllerTransactor: ZenBTControllerTransactor{contract: contract}, ZenBTControllerFilterer: ZenBTControllerFilterer{contract: contract}}, nil
}

// ZenBTController is an auto generated Go binding around an Ethereum contract.
type ZenBTController struct {
	ZenBTControllerCaller     // Read-only binding to the contract
	ZenBTControllerTransactor // Write-only binding to the contract
	ZenBTControllerFilterer   // Log filterer for contract events
}

// ZenBTControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZenBTControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZenBTControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZenBTControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZenBTControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZenBTControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZenBTControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZenBTControllerSession struct {
	Contract     *ZenBTController  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZenBTControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZenBTControllerCallerSession struct {
	Contract *ZenBTControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ZenBTControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZenBTControllerTransactorSession struct {
	Contract     *ZenBTControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ZenBTControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZenBTControllerRaw struct {
	Contract *ZenBTController // Generic contract binding to access the raw methods on
}

// ZenBTControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZenBTControllerCallerRaw struct {
	Contract *ZenBTControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ZenBTControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZenBTControllerTransactorRaw struct {
	Contract *ZenBTControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZenBTController creates a new instance of ZenBTController, bound to a specific deployed contract.
func NewZenBTController(address common.Address, backend bind.ContractBackend) (*ZenBTController, error) {
	contract, err := bindZenBTController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZenBTController{ZenBTControllerCaller: ZenBTControllerCaller{contract: contract}, ZenBTControllerTransactor: ZenBTControllerTransactor{contract: contract}, ZenBTControllerFilterer: ZenBTControllerFilterer{contract: contract}}, nil
}

// NewZenBTControllerCaller creates a new read-only instance of ZenBTController, bound to a specific deployed contract.
func NewZenBTControllerCaller(address common.Address, caller bind.ContractCaller) (*ZenBTControllerCaller, error) {
	contract, err := bindZenBTController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZenBTControllerCaller{contract: contract}, nil
}

// NewZenBTControllerTransactor creates a new write-only instance of ZenBTController, bound to a specific deployed contract.
func NewZenBTControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ZenBTControllerTransactor, error) {
	contract, err := bindZenBTController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZenBTControllerTransactor{contract: contract}, nil
}

// NewZenBTControllerFilterer creates a new log filterer instance of ZenBTController, bound to a specific deployed contract.
func NewZenBTControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ZenBTControllerFilterer, error) {
	contract, err := bindZenBTController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZenBTControllerFilterer{contract: contract}, nil
}

// bindZenBTController binds a generic wrapper to an already deployed contract.
func bindZenBTController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZenBTControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZenBTController *ZenBTControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZenBTController.Contract.ZenBTControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZenBTController *ZenBTControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZenBTController.Contract.ZenBTControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZenBTController *ZenBTControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZenBTController.Contract.ZenBTControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZenBTController *ZenBTControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZenBTController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZenBTController *ZenBTControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZenBTController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZenBTController *ZenBTControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZenBTController.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZenBTController *ZenBTControllerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZenBTController.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZenBTController *ZenBTControllerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZenBTController.Contract.DEFAULTADMINROLE(&_ZenBTController.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZenBTController *ZenBTControllerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZenBTController.Contract.DEFAULTADMINROLE(&_ZenBTController.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ZenBTController *ZenBTControllerCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZenBTController.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ZenBTController *ZenBTControllerSession) MINTERROLE() ([32]byte, error) {
	return _ZenBTController.Contract.MINTERROLE(&_ZenBTController.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ZenBTController *ZenBTControllerCallerSession) MINTERROLE() ([32]byte, error) {
	return _ZenBTController.Contract.MINTERROLE(&_ZenBTController.CallOpts)
}

// GetEGLDelegationManager is a free data retrieval call binding the contract method 0x3725d2fc.
//
// Solidity: function getEGLDelegationManager() view returns(address)
func (_ZenBTController *ZenBTControllerCaller) GetEGLDelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZenBTController.contract.Call(opts, &out, "getEGLDelegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEGLDelegationManager is a free data retrieval call binding the contract method 0x3725d2fc.
//
// Solidity: function getEGLDelegationManager() view returns(address)
func (_ZenBTController *ZenBTControllerSession) GetEGLDelegationManager() (common.Address, error) {
	return _ZenBTController.Contract.GetEGLDelegationManager(&_ZenBTController.CallOpts)
}

// GetEGLDelegationManager is a free data retrieval call binding the contract method 0x3725d2fc.
//
// Solidity: function getEGLDelegationManager() view returns(address)
func (_ZenBTController *ZenBTControllerCallerSession) GetEGLDelegationManager() (common.Address, error) {
	return _ZenBTController.Contract.GetEGLDelegationManager(&_ZenBTController.CallOpts)
}

// GetEGLStrategyManager is a free data retrieval call binding the contract method 0x1f05b551.
//
// Solidity: function getEGLStrategyManager() view returns(address)
func (_ZenBTController *ZenBTControllerCaller) GetEGLStrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZenBTController.contract.Call(opts, &out, "getEGLStrategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEGLStrategyManager is a free data retrieval call binding the contract method 0x1f05b551.
//
// Solidity: function getEGLStrategyManager() view returns(address)
func (_ZenBTController *ZenBTControllerSession) GetEGLStrategyManager() (common.Address, error) {
	return _ZenBTController.Contract.GetEGLStrategyManager(&_ZenBTController.CallOpts)
}

// GetEGLStrategyManager is a free data retrieval call binding the contract method 0x1f05b551.
//
// Solidity: function getEGLStrategyManager() view returns(address)
func (_ZenBTController *ZenBTControllerCallerSession) GetEGLStrategyManager() (common.Address, error) {
	return _ZenBTController.Contract.GetEGLStrategyManager(&_ZenBTController.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_ZenBTController *ZenBTControllerCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZenBTController.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_ZenBTController *ZenBTControllerSession) GetNonce() (*big.Int, error) {
	return _ZenBTController.Contract.GetNonce(&_ZenBTController.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_ZenBTController *ZenBTControllerCallerSession) GetNonce() (*big.Int, error) {
	return _ZenBTController.Contract.GetNonce(&_ZenBTController.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_ZenBTController *ZenBTControllerCaller) GetOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZenBTController.contract.Call(opts, &out, "getOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_ZenBTController *ZenBTControllerSession) GetOperator() (common.Address, error) {
	return _ZenBTController.Contract.GetOperator(&_ZenBTController.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_ZenBTController *ZenBTControllerCallerSession) GetOperator() (common.Address, error) {
	return _ZenBTController.Contract.GetOperator(&_ZenBTController.CallOpts)
}

// GetRockBTC is a free data retrieval call binding the contract method 0x04611b06.
//
// Solidity: function getRockBTC() view returns(address)
func (_ZenBTController *ZenBTControllerCaller) GetRockBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZenBTController.contract.Call(opts, &out, "getRockBTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRockBTC is a free data retrieval call binding the contract method 0x04611b06.
//
// Solidity: function getRockBTC() view returns(address)
func (_ZenBTController *ZenBTControllerSession) GetRockBTC() (common.Address, error) {
	return _ZenBTController.Contract.GetRockBTC(&_ZenBTController.CallOpts)
}

// GetRockBTC is a free data retrieval call binding the contract method 0x04611b06.
//
// Solidity: function getRockBTC() view returns(address)
func (_ZenBTController *ZenBTControllerCallerSession) GetRockBTC() (common.Address, error) {
	return _ZenBTController.Contract.GetRockBTC(&_ZenBTController.CallOpts)
}

// GetRockBTCStrategy is a free data retrieval call binding the contract method 0x931432ab.
//
// Solidity: function getRockBTCStrategy() view returns(address)
func (_ZenBTController *ZenBTControllerCaller) GetRockBTCStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZenBTController.contract.Call(opts, &out, "getRockBTCStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRockBTCStrategy is a free data retrieval call binding the contract method 0x931432ab.
//
// Solidity: function getRockBTCStrategy() view returns(address)
func (_ZenBTController *ZenBTControllerSession) GetRockBTCStrategy() (common.Address, error) {
	return _ZenBTController.Contract.GetRockBTCStrategy(&_ZenBTController.CallOpts)
}

// GetRockBTCStrategy is a free data retrieval call binding the contract method 0x931432ab.
//
// Solidity: function getRockBTCStrategy() view returns(address)
func (_ZenBTController *ZenBTControllerCallerSession) GetRockBTCStrategy() (common.Address, error) {
	return _ZenBTController.Contract.GetRockBTCStrategy(&_ZenBTController.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZenBTController *ZenBTControllerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZenBTController.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZenBTController *ZenBTControllerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZenBTController.Contract.GetRoleAdmin(&_ZenBTController.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZenBTController *ZenBTControllerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZenBTController.Contract.GetRoleAdmin(&_ZenBTController.CallOpts, role)
}

// GetZenBTC is a free data retrieval call binding the contract method 0xf7d66494.
//
// Solidity: function getZenBTC() view returns(address)
func (_ZenBTController *ZenBTControllerCaller) GetZenBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZenBTController.contract.Call(opts, &out, "getZenBTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetZenBTC is a free data retrieval call binding the contract method 0xf7d66494.
//
// Solidity: function getZenBTC() view returns(address)
func (_ZenBTController *ZenBTControllerSession) GetZenBTC() (common.Address, error) {
	return _ZenBTController.Contract.GetZenBTC(&_ZenBTController.CallOpts)
}

// GetZenBTC is a free data retrieval call binding the contract method 0xf7d66494.
//
// Solidity: function getZenBTC() view returns(address)
func (_ZenBTController *ZenBTControllerCallerSession) GetZenBTC() (common.Address, error) {
	return _ZenBTController.Contract.GetZenBTC(&_ZenBTController.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZenBTController *ZenBTControllerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZenBTController.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZenBTController *ZenBTControllerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZenBTController.Contract.HasRole(&_ZenBTController.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZenBTController *ZenBTControllerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZenBTController.Contract.HasRole(&_ZenBTController.CallOpts, role, account)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4 magicValue)
func (_ZenBTController *ZenBTControllerCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _ZenBTController.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4 magicValue)
func (_ZenBTController *ZenBTControllerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _ZenBTController.Contract.IsValidSignature(&_ZenBTController.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4 magicValue)
func (_ZenBTController *ZenBTControllerCallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _ZenBTController.Contract.IsValidSignature(&_ZenBTController.CallOpts, _hash, _signature)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZenBTController *ZenBTControllerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZenBTController.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZenBTController *ZenBTControllerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZenBTController.Contract.SupportsInterface(&_ZenBTController.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZenBTController *ZenBTControllerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZenBTController.Contract.SupportsInterface(&_ZenBTController.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZenBTController *ZenBTControllerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZenBTController.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZenBTController *ZenBTControllerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZenBTController.Contract.GrantRole(&_ZenBTController.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZenBTController *ZenBTControllerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZenBTController.Contract.GrantRole(&_ZenBTController.TransactOpts, role, account)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0xfd26fc66.
//
// Solidity: function initializeV1(address _zenBTC, address _rockBTC, address _eglStrategyManager, address _eglDelegationManager, address _rockBTCStrategy, address _operator) returns()
func (_ZenBTController *ZenBTControllerTransactor) InitializeV1(opts *bind.TransactOpts, _zenBTC common.Address, _rockBTC common.Address, _eglStrategyManager common.Address, _eglDelegationManager common.Address, _rockBTCStrategy common.Address, _operator common.Address) (*types.Transaction, error) {
	return _ZenBTController.contract.Transact(opts, "initializeV1", _zenBTC, _rockBTC, _eglStrategyManager, _eglDelegationManager, _rockBTCStrategy, _operator)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0xfd26fc66.
//
// Solidity: function initializeV1(address _zenBTC, address _rockBTC, address _eglStrategyManager, address _eglDelegationManager, address _rockBTCStrategy, address _operator) returns()
func (_ZenBTController *ZenBTControllerSession) InitializeV1(_zenBTC common.Address, _rockBTC common.Address, _eglStrategyManager common.Address, _eglDelegationManager common.Address, _rockBTCStrategy common.Address, _operator common.Address) (*types.Transaction, error) {
	return _ZenBTController.Contract.InitializeV1(&_ZenBTController.TransactOpts, _zenBTC, _rockBTC, _eglStrategyManager, _eglDelegationManager, _rockBTCStrategy, _operator)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0xfd26fc66.
//
// Solidity: function initializeV1(address _zenBTC, address _rockBTC, address _eglStrategyManager, address _eglDelegationManager, address _rockBTCStrategy, address _operator) returns()
func (_ZenBTController *ZenBTControllerTransactorSession) InitializeV1(_zenBTC common.Address, _rockBTC common.Address, _eglStrategyManager common.Address, _eglDelegationManager common.Address, _rockBTCStrategy common.Address, _operator common.Address) (*types.Transaction, error) {
	return _ZenBTController.Contract.InitializeV1(&_ZenBTController.TransactOpts, _zenBTC, _rockBTC, _eglStrategyManager, _eglDelegationManager, _rockBTCStrategy, _operator)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZenBTController *ZenBTControllerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZenBTController.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZenBTController *ZenBTControllerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZenBTController.Contract.RenounceRole(&_ZenBTController.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZenBTController *ZenBTControllerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZenBTController.Contract.RenounceRole(&_ZenBTController.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZenBTController *ZenBTControllerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZenBTController.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZenBTController *ZenBTControllerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZenBTController.Contract.RevokeRole(&_ZenBTController.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZenBTController *ZenBTControllerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZenBTController.Contract.RevokeRole(&_ZenBTController.TransactOpts, role, account)
}

// StakeRockBTC is a paid mutator transaction binding the contract method 0x6296ae27.
//
// Solidity: function stakeRockBTC(uint64 value, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ZenBTController *ZenBTControllerTransactor) StakeRockBTC(opts *bind.TransactOpts, value uint64, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ZenBTController.contract.Transact(opts, "stakeRockBTC", value, approverSignatureAndExpiry, approverSalt)
}

// StakeRockBTC is a paid mutator transaction binding the contract method 0x6296ae27.
//
// Solidity: function stakeRockBTC(uint64 value, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ZenBTController *ZenBTControllerSession) StakeRockBTC(value uint64, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ZenBTController.Contract.StakeRockBTC(&_ZenBTController.TransactOpts, value, approverSignatureAndExpiry, approverSalt)
}

// StakeRockBTC is a paid mutator transaction binding the contract method 0x6296ae27.
//
// Solidity: function stakeRockBTC(uint64 value, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ZenBTController *ZenBTControllerTransactorSession) StakeRockBTC(value uint64, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ZenBTController.Contract.StakeRockBTC(&_ZenBTController.TransactOpts, value, approverSignatureAndExpiry, approverSalt)
}

// UnwrapComplete is a paid mutator transaction binding the contract method 0x9915c15f.
//
// Solidity: function unwrapComplete((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_ZenBTController *ZenBTControllerTransactor) UnwrapComplete(opts *bind.TransactOpts, withdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _ZenBTController.contract.Transact(opts, "unwrapComplete", withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// UnwrapComplete is a paid mutator transaction binding the contract method 0x9915c15f.
//
// Solidity: function unwrapComplete((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_ZenBTController *ZenBTControllerSession) UnwrapComplete(withdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _ZenBTController.Contract.UnwrapComplete(&_ZenBTController.TransactOpts, withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// UnwrapComplete is a paid mutator transaction binding the contract method 0x9915c15f.
//
// Solidity: function unwrapComplete((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_ZenBTController *ZenBTControllerTransactorSession) UnwrapComplete(withdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _ZenBTController.Contract.UnwrapComplete(&_ZenBTController.TransactOpts, withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// UnwrapInit is a paid mutator transaction binding the contract method 0xe2d1249a.
//
// Solidity: function unwrapInit((address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes32[] withdrawalIds)
func (_ZenBTController *ZenBTControllerTransactor) UnwrapInit(opts *bind.TransactOpts, queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _ZenBTController.contract.Transact(opts, "unwrapInit", queuedWithdrawalParams)
}

// UnwrapInit is a paid mutator transaction binding the contract method 0xe2d1249a.
//
// Solidity: function unwrapInit((address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes32[] withdrawalIds)
func (_ZenBTController *ZenBTControllerSession) UnwrapInit(queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _ZenBTController.Contract.UnwrapInit(&_ZenBTController.TransactOpts, queuedWithdrawalParams)
}

// UnwrapInit is a paid mutator transaction binding the contract method 0xe2d1249a.
//
// Solidity: function unwrapInit((address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes32[] withdrawalIds)
func (_ZenBTController *ZenBTControllerTransactorSession) UnwrapInit(queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _ZenBTController.Contract.UnwrapInit(&_ZenBTController.TransactOpts, queuedWithdrawalParams)
}

// WrapZenBTC is a paid mutator transaction binding the contract method 0xb52e4e8c.
//
// Solidity: function wrapZenBTC(address to, uint64 value, uint64 fee, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ZenBTController *ZenBTControllerTransactor) WrapZenBTC(opts *bind.TransactOpts, to common.Address, value uint64, fee uint64, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ZenBTController.contract.Transact(opts, "wrapZenBTC", to, value, fee, approverSignatureAndExpiry, approverSalt)
}

// WrapZenBTC is a paid mutator transaction binding the contract method 0xb52e4e8c.
//
// Solidity: function wrapZenBTC(address to, uint64 value, uint64 fee, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ZenBTController *ZenBTControllerSession) WrapZenBTC(to common.Address, value uint64, fee uint64, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ZenBTController.Contract.WrapZenBTC(&_ZenBTController.TransactOpts, to, value, fee, approverSignatureAndExpiry, approverSalt)
}

// WrapZenBTC is a paid mutator transaction binding the contract method 0xb52e4e8c.
//
// Solidity: function wrapZenBTC(address to, uint64 value, uint64 fee, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ZenBTController *ZenBTControllerTransactorSession) WrapZenBTC(to common.Address, value uint64, fee uint64, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ZenBTController.Contract.WrapZenBTC(&_ZenBTController.TransactOpts, to, value, fee, approverSignatureAndExpiry, approverSalt)
}

// ZenBTControllerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZenBTController contract.
type ZenBTControllerInitializedIterator struct {
	Event *ZenBTControllerInitialized // Event containing the contract specifics and raw log

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
func (it *ZenBTControllerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTControllerInitialized)
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
		it.Event = new(ZenBTControllerInitialized)
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
func (it *ZenBTControllerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTControllerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTControllerInitialized represents a Initialized event raised by the ZenBTController contract.
type ZenBTControllerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZenBTController *ZenBTControllerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZenBTControllerInitializedIterator, error) {

	logs, sub, err := _ZenBTController.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZenBTControllerInitializedIterator{contract: _ZenBTController.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZenBTController *ZenBTControllerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZenBTControllerInitialized) (event.Subscription, error) {

	logs, sub, err := _ZenBTController.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTControllerInitialized)
				if err := _ZenBTController.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZenBTController *ZenBTControllerFilterer) ParseInitialized(log types.Log) (*ZenBTControllerInitialized, error) {
	event := new(ZenBTControllerInitialized)
	if err := _ZenBTController.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTControllerRockBTCStakedIterator is returned from FilterRockBTCStaked and is used to iterate over the raw logs and unpacked data for RockBTCStaked events raised by the ZenBTController contract.
type ZenBTControllerRockBTCStakedIterator struct {
	Event *ZenBTControllerRockBTCStaked // Event containing the contract specifics and raw log

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
func (it *ZenBTControllerRockBTCStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTControllerRockBTCStaked)
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
		it.Event = new(ZenBTControllerRockBTCStaked)
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
func (it *ZenBTControllerRockBTCStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTControllerRockBTCStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTControllerRockBTCStaked represents a RockBTCStaked event raised by the ZenBTController contract.
type ZenBTControllerRockBTCStaked struct {
	Operator common.Address
	Value    uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRockBTCStaked is a free log retrieval operation binding the contract event 0x79c412fdc602b15f36aa9b42405335e8dac8e322a516e71a0a3881d5ed41e3bc.
//
// Solidity: event RockBTCStaked(address indexed operator, uint64 value)
func (_ZenBTController *ZenBTControllerFilterer) FilterRockBTCStaked(opts *bind.FilterOpts, operator []common.Address) (*ZenBTControllerRockBTCStakedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ZenBTController.contract.FilterLogs(opts, "RockBTCStaked", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ZenBTControllerRockBTCStakedIterator{contract: _ZenBTController.contract, event: "RockBTCStaked", logs: logs, sub: sub}, nil
}

// WatchRockBTCStaked is a free log subscription operation binding the contract event 0x79c412fdc602b15f36aa9b42405335e8dac8e322a516e71a0a3881d5ed41e3bc.
//
// Solidity: event RockBTCStaked(address indexed operator, uint64 value)
func (_ZenBTController *ZenBTControllerFilterer) WatchRockBTCStaked(opts *bind.WatchOpts, sink chan<- *ZenBTControllerRockBTCStaked, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ZenBTController.contract.WatchLogs(opts, "RockBTCStaked", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTControllerRockBTCStaked)
				if err := _ZenBTController.contract.UnpackLog(event, "RockBTCStaked", log); err != nil {
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

// ParseRockBTCStaked is a log parse operation binding the contract event 0x79c412fdc602b15f36aa9b42405335e8dac8e322a516e71a0a3881d5ed41e3bc.
//
// Solidity: event RockBTCStaked(address indexed operator, uint64 value)
func (_ZenBTController *ZenBTControllerFilterer) ParseRockBTCStaked(log types.Log) (*ZenBTControllerRockBTCStaked, error) {
	event := new(ZenBTControllerRockBTCStaked)
	if err := _ZenBTController.contract.UnpackLog(event, "RockBTCStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTControllerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZenBTController contract.
type ZenBTControllerRoleAdminChangedIterator struct {
	Event *ZenBTControllerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZenBTControllerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTControllerRoleAdminChanged)
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
		it.Event = new(ZenBTControllerRoleAdminChanged)
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
func (it *ZenBTControllerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTControllerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTControllerRoleAdminChanged represents a RoleAdminChanged event raised by the ZenBTController contract.
type ZenBTControllerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZenBTController *ZenBTControllerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZenBTControllerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ZenBTController.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZenBTControllerRoleAdminChangedIterator{contract: _ZenBTController.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZenBTController *ZenBTControllerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZenBTControllerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ZenBTController.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTControllerRoleAdminChanged)
				if err := _ZenBTController.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ZenBTController *ZenBTControllerFilterer) ParseRoleAdminChanged(log types.Log) (*ZenBTControllerRoleAdminChanged, error) {
	event := new(ZenBTControllerRoleAdminChanged)
	if err := _ZenBTController.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTControllerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZenBTController contract.
type ZenBTControllerRoleGrantedIterator struct {
	Event *ZenBTControllerRoleGranted // Event containing the contract specifics and raw log

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
func (it *ZenBTControllerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTControllerRoleGranted)
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
		it.Event = new(ZenBTControllerRoleGranted)
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
func (it *ZenBTControllerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTControllerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTControllerRoleGranted represents a RoleGranted event raised by the ZenBTController contract.
type ZenBTControllerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZenBTController *ZenBTControllerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZenBTControllerRoleGrantedIterator, error) {

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

	logs, sub, err := _ZenBTController.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZenBTControllerRoleGrantedIterator{contract: _ZenBTController.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZenBTController *ZenBTControllerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZenBTControllerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ZenBTController.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTControllerRoleGranted)
				if err := _ZenBTController.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ZenBTController *ZenBTControllerFilterer) ParseRoleGranted(log types.Log) (*ZenBTControllerRoleGranted, error) {
	event := new(ZenBTControllerRoleGranted)
	if err := _ZenBTController.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTControllerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZenBTController contract.
type ZenBTControllerRoleRevokedIterator struct {
	Event *ZenBTControllerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ZenBTControllerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTControllerRoleRevoked)
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
		it.Event = new(ZenBTControllerRoleRevoked)
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
func (it *ZenBTControllerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTControllerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTControllerRoleRevoked represents a RoleRevoked event raised by the ZenBTController contract.
type ZenBTControllerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZenBTController *ZenBTControllerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZenBTControllerRoleRevokedIterator, error) {

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

	logs, sub, err := _ZenBTController.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZenBTControllerRoleRevokedIterator{contract: _ZenBTController.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZenBTController *ZenBTControllerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZenBTControllerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ZenBTController.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTControllerRoleRevoked)
				if err := _ZenBTController.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ZenBTController *ZenBTControllerFilterer) ParseRoleRevoked(log types.Log) (*ZenBTControllerRoleRevoked, error) {
	event := new(ZenBTControllerRoleRevoked)
	if err := _ZenBTController.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTControllerWithdrawalCompletedIterator is returned from FilterWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for WithdrawalCompleted events raised by the ZenBTController contract.
type ZenBTControllerWithdrawalCompletedIterator struct {
	Event *ZenBTControllerWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *ZenBTControllerWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTControllerWithdrawalCompleted)
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
		it.Event = new(ZenBTControllerWithdrawalCompleted)
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
func (it *ZenBTControllerWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTControllerWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTControllerWithdrawalCompleted represents a WithdrawalCompleted event raised by the ZenBTController contract.
type ZenBTControllerWithdrawalCompleted struct {
	Arg0            IDelegationManagerWithdrawal
	Tokens          []common.Address
	ReceiveAsTokens bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCompleted is a free log retrieval operation binding the contract event 0xadd5e0a5a1ec72776e3a7800fa2aec5ca0f7865cf68592ff037e468cb272a0a2.
//
// Solidity: event WithdrawalCompleted((address,address,address,uint256,uint32,address[],uint256[]) arg0, address[] tokens, bool receiveAsTokens)
func (_ZenBTController *ZenBTControllerFilterer) FilterWithdrawalCompleted(opts *bind.FilterOpts) (*ZenBTControllerWithdrawalCompletedIterator, error) {

	logs, sub, err := _ZenBTController.contract.FilterLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return &ZenBTControllerWithdrawalCompletedIterator{contract: _ZenBTController.contract, event: "WithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCompleted is a free log subscription operation binding the contract event 0xadd5e0a5a1ec72776e3a7800fa2aec5ca0f7865cf68592ff037e468cb272a0a2.
//
// Solidity: event WithdrawalCompleted((address,address,address,uint256,uint32,address[],uint256[]) arg0, address[] tokens, bool receiveAsTokens)
func (_ZenBTController *ZenBTControllerFilterer) WatchWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ZenBTControllerWithdrawalCompleted) (event.Subscription, error) {

	logs, sub, err := _ZenBTController.contract.WatchLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTControllerWithdrawalCompleted)
				if err := _ZenBTController.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
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

// ParseWithdrawalCompleted is a log parse operation binding the contract event 0xadd5e0a5a1ec72776e3a7800fa2aec5ca0f7865cf68592ff037e468cb272a0a2.
//
// Solidity: event WithdrawalCompleted((address,address,address,uint256,uint32,address[],uint256[]) arg0, address[] tokens, bool receiveAsTokens)
func (_ZenBTController *ZenBTControllerFilterer) ParseWithdrawalCompleted(log types.Log) (*ZenBTControllerWithdrawalCompleted, error) {
	event := new(ZenBTControllerWithdrawalCompleted)
	if err := _ZenBTController.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTControllerWithdrawalInitiatedIterator is returned from FilterWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for WithdrawalInitiated events raised by the ZenBTController contract.
type ZenBTControllerWithdrawalInitiatedIterator struct {
	Event *ZenBTControllerWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *ZenBTControllerWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTControllerWithdrawalInitiated)
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
		it.Event = new(ZenBTControllerWithdrawalInitiated)
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
func (it *ZenBTControllerWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTControllerWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTControllerWithdrawalInitiated represents a WithdrawalInitiated event raised by the ZenBTController contract.
type ZenBTControllerWithdrawalInitiated struct {
	WithdrawalIds [][32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiated is a free log retrieval operation binding the contract event 0x3ca8ee7ddc9cba27865c2837ec26df7a102ce7b8f538096db5b5d42e863502af.
//
// Solidity: event WithdrawalInitiated(bytes32[] withdrawalIds)
func (_ZenBTController *ZenBTControllerFilterer) FilterWithdrawalInitiated(opts *bind.FilterOpts) (*ZenBTControllerWithdrawalInitiatedIterator, error) {

	logs, sub, err := _ZenBTController.contract.FilterLogs(opts, "WithdrawalInitiated")
	if err != nil {
		return nil, err
	}
	return &ZenBTControllerWithdrawalInitiatedIterator{contract: _ZenBTController.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiated is a free log subscription operation binding the contract event 0x3ca8ee7ddc9cba27865c2837ec26df7a102ce7b8f538096db5b5d42e863502af.
//
// Solidity: event WithdrawalInitiated(bytes32[] withdrawalIds)
func (_ZenBTController *ZenBTControllerFilterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *ZenBTControllerWithdrawalInitiated) (event.Subscription, error) {

	logs, sub, err := _ZenBTController.contract.WatchLogs(opts, "WithdrawalInitiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTControllerWithdrawalInitiated)
				if err := _ZenBTController.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
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

// ParseWithdrawalInitiated is a log parse operation binding the contract event 0x3ca8ee7ddc9cba27865c2837ec26df7a102ce7b8f538096db5b5d42e863502af.
//
// Solidity: event WithdrawalInitiated(bytes32[] withdrawalIds)
func (_ZenBTController *ZenBTControllerFilterer) ParseWithdrawalInitiated(log types.Log) (*ZenBTControllerWithdrawalInitiated, error) {
	event := new(ZenBTControllerWithdrawalInitiated)
	if err := _ZenBTController.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTControllerZenBTCWrappedIterator is returned from FilterZenBTCWrapped and is used to iterate over the raw logs and unpacked data for ZenBTCWrapped events raised by the ZenBTController contract.
type ZenBTControllerZenBTCWrappedIterator struct {
	Event *ZenBTControllerZenBTCWrapped // Event containing the contract specifics and raw log

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
func (it *ZenBTControllerZenBTCWrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTControllerZenBTCWrapped)
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
		it.Event = new(ZenBTControllerZenBTCWrapped)
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
func (it *ZenBTControllerZenBTCWrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTControllerZenBTCWrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTControllerZenBTCWrapped represents a ZenBTCWrapped event raised by the ZenBTController contract.
type ZenBTControllerZenBTCWrapped struct {
	To       common.Address
	Value    uint64
	Fee      uint64
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterZenBTCWrapped is a free log retrieval operation binding the contract event 0xfd60f0aa3e690ac3a0ef44fb08201a6d258f27749c174539d4d1cf9c2ca0f8d7.
//
// Solidity: event ZenBTCWrapped(address indexed to, uint64 value, uint64 fee, address indexed operator)
func (_ZenBTController *ZenBTControllerFilterer) FilterZenBTCWrapped(opts *bind.FilterOpts, to []common.Address, operator []common.Address) (*ZenBTControllerZenBTCWrappedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ZenBTController.contract.FilterLogs(opts, "ZenBTCWrapped", toRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ZenBTControllerZenBTCWrappedIterator{contract: _ZenBTController.contract, event: "ZenBTCWrapped", logs: logs, sub: sub}, nil
}

// WatchZenBTCWrapped is a free log subscription operation binding the contract event 0xfd60f0aa3e690ac3a0ef44fb08201a6d258f27749c174539d4d1cf9c2ca0f8d7.
//
// Solidity: event ZenBTCWrapped(address indexed to, uint64 value, uint64 fee, address indexed operator)
func (_ZenBTController *ZenBTControllerFilterer) WatchZenBTCWrapped(opts *bind.WatchOpts, sink chan<- *ZenBTControllerZenBTCWrapped, to []common.Address, operator []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ZenBTController.contract.WatchLogs(opts, "ZenBTCWrapped", toRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTControllerZenBTCWrapped)
				if err := _ZenBTController.contract.UnpackLog(event, "ZenBTCWrapped", log); err != nil {
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

// ParseZenBTCWrapped is a log parse operation binding the contract event 0xfd60f0aa3e690ac3a0ef44fb08201a6d258f27749c174539d4d1cf9c2ca0f8d7.
//
// Solidity: event ZenBTCWrapped(address indexed to, uint64 value, uint64 fee, address indexed operator)
func (_ZenBTController *ZenBTControllerFilterer) ParseZenBTCWrapped(log types.Log) (*ZenBTControllerZenBTCWrapped, error) {
	event := new(ZenBTControllerZenBTCWrapped)
	if err := _ZenBTController.contract.UnpackLog(event, "ZenBTCWrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}