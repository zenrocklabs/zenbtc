// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ZenBTCMetaData contains all meta data concerning the ZenBTC contract.
var ZenBTCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"basisPointsRateUpdated\",\"type\":\"uint16\"}],\"name\":\"BasisPointsRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeAddress\",\"type\":\"address\"}],\"name\":\"FeeAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRedeemFee\",\"type\":\"uint256\"}],\"name\":\"RedeemFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TokenRedemption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fee\",\"type\":\"uint64\"}],\"name\":\"TokensMintedWithFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBasisPointsRate\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"numBlocks\",\"type\":\"uint16\"}],\"name\":\"getRecentRedemptionData\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"ids\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"amounts\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"destination\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initializeV1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destAddr\",\"type\":\"bytes\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"basisPointsRate_\",\"type\":\"uint16\"}],\"name\":\"updateBasisPointsRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAddress_\",\"type\":\"address\"}],\"name\":\"updateFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fee\",\"type\":\"uint64\"}],\"name\":\"wrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106101cf5760003560e01c80634e7ceacb11610104578063bbcaac38116100a2578063f1690dbe11610071578063f1690dbe1461043b578063f1d96dd31461045d578063fd154cc714610470578063fded25291461048357600080fd5b8063bbcaac38146103ed578063d539139314610400578063d547741f14610415578063dd62ed3e1461042857600080fd5b806395d89b41116100de57806395d89b41146103b7578063a217fddf146103bf578063a9059cbb146103c7578063b413148e146103da57600080fd5b80634e7ceacb1461034657806370a082311461036e57806391d14854146103a457600080fd5b80632a0276f811610171578063313ce5671161014b578063313ce567146102e157806336568abe1461030d57806340c10f191461032057806342966c681461033357600080fd5b80632a0276f8146102925780632a4b2fd9146102b95780632f2ff15d146102ce57600080fd5b8063127e8e4d116101ad578063127e8e4d1461022457806318160ddd1461024557806323b872dd1461026c578063248a9ca31461027f57600080fd5b806301ffc9a7146101d457806306fdde03146101fc578063095ea7b314610211575b600080fd5b6101e76101e2366004611c09565b6104b1565b60405190151581526020015b60405180910390f35b6102046104e8565b6040516101f39190611c79565b6101e761021f366004611ca8565b6105ab565b610237610232366004611cd2565b6105c3565b6040519081526020016101f3565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0254610237565b6101e761027a366004611ceb565b6105ce565b61023761028d366004611cd2565b6105f2565b6102377f13b7ad447453d194d272cdda9bb09d7d357cda1ab7de80d865b4c1cbefc3cf2881565b6102cc6102c7366004611dd2565b610614565b005b6102cc6102dc366004611e4f565b61072e565b6000805160206122c783398151915254600160b01b900460ff1660405160ff90911681526020016101f3565b6102cc61031b366004611e4f565b610750565b6102cc61032e366004611ca8565b610788565b6102cc610341366004611cd2565b6107aa565b6000805160206122c7833981519152546040516001600160a01b0390911681526020016101f3565b61023761037c366004611e7b565b6001600160a01b031660009081526000805160206122e7833981519152602052604090205490565b6101e76103b2366004611e4f565b6107d0565b610204610808565b610237600081565b6101e76103d5366004611ca8565b610847565b6102cc6103e8366004611e96565b610855565b6102cc6103fb366004611e7b565b610945565b61023760008051602061234783398151915281565b6102cc610423366004611e4f565b610978565b610237610436366004611ef0565b610994565b61044e610449366004611f1a565b6109de565b6040516101f393929190611f83565b6102cc61046b366004612021565b610e4d565b6102cc61047e366004611f1a565b610fc8565b6000805160206122c783398151915254600160a01b900461ffff1660405161ffff90911681526020016101f3565b60006001600160e01b03198216637965db0b60e01b14806104e257506301ffc9a760e01b6001600160e01b03198316145b92915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060916000805160206122e78339815191529161052790612064565b80601f016020809104026020016040519081016040528092919081815260200182805461055390612064565b80156105a05780601f10610575576101008083540402835291602001916105a0565b820191906000526020600020905b81548152906001019060200180831161058357829003601f168201915b505050505091505090565b6000336105b9818585610ffb565b5060019392505050565b60006104e282611008565b6000336105dc8582856110ec565b6105e785858561114c565b506001949350505050565b6000908152600080516020612327833981519152602052604090206001015490565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156106595750825b90506000826001600160401b031660011480156106755750303b155b905081158015610683575080155b156106a15760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156106cb57845460ff60401b1916600160401b1785555b6106d36111d6565b6106de8888886111e8565b831561072457845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b610737826105f2565b6107408161120b565b61074a8383611218565b50505050565b6001600160a01b03811633146107795760405163334bd91960e11b815260040160405180910390fd5b61078382826112c4565b505050565b6000805160206123478339815191526107a08161120b565b6107838383611340565b6000805160206123478339815191526107c28161120b565b6107cc3383611376565b5050565b6000918252600080516020612327833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0580546060916000805160206122e78339815191529161052790612064565b6000336105b981858561114c565b6000805160206122e78339815191526001600160401b038311156108b25760405162461bcd60e51b815260206004820152600f60248201526e56616c756520746f6f206c6172676560881b60448201526064015b60405180910390fd5b60006108bd84611008565b90506108c93385611376565b60038201546108e1906001600160a01b031682611340565b60006108ed82866120b4565b90506108f981856113ac565b336001600160a01b03167f4c971c8b2abb197a17896b2fc57f597830db78d3556831e3faa337a596150f22828685604051610936939291906120c7565b60405180910390a25050505050565b7f13b7ad447453d194d272cdda9bb09d7d357cda1ab7de80d865b4c1cbefc3cf2861096f8161120b565b6107cc82611589565b610981826105f2565b61098a8161120b565b61074a83836112c4565b6001600160a01b0391821660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b60608080606461ffff85161115610a295760405162461bcd60e51b815260206004820152600f60248201526e52616e676520746f6f206c6172676560881b60448201526064016108a9565b6000805160206123078339815191526000805b8661ffff168161ffff161015610ad357825460009060649061ffff908116849003606301160690506000846001018261ffff1681548110610a7f57610a7f61210f565b6000918252602090912060088204015460079091166004026101000a900463ffffffff1690508015610ac95763ffffffff8116600090815260028601602052604090205493909301925b5050600101610a3c565b50806001600160401b03811115610aec57610aec611d27565b604051908082528060200260200182016040528015610b15578160200160208202803683370190505b509450806001600160401b03811115610b3057610b30611d27565b604051908082528060200260200182016040528015610b59578160200160208202803683370190505b509350806001600160401b03811115610b7457610b74611d27565b604051908082528060200260200182016040528015610ba757816020015b6060815260200190600190039081610b925790505b5092506000805b8761ffff168161ffff161015610e4257835460009060649061ffff908116849003606301160690506000856001018261ffff1681548110610bf157610bf161210f565b6000918252602090912060088204015460079091166004026101000a900463ffffffff1690508015610e385763ffffffff81166000908152600287016020908152604080832080548251818502810185019093528083529192909190849084015b82821015610d3c576000848152602090819020604080516060810182526002860290920180546001600160401b038082168552600160401b909104169383019390935260018301805492939291840191610cab90612064565b80601f0160208091040260200160405190810160405280929190818152602001828054610cd790612064565b8015610d245780601f10610cf957610100808354040283529160200191610d24565b820191906000526020600020905b815481529060010190602001808311610d0757829003601f168201915b50505050508152505081526020019060010190610c52565b50505050905060005b8151811015610e3557818181518110610d6057610d6061210f565b6020026020010151600001518b8781518110610d7e57610d7e61210f565b60200260200101906001600160401b031690816001600160401b031681525050818181518110610db057610db061210f565b6020026020010151602001518a8781518110610dce57610dce61210f565b60200260200101906001600160401b031690816001600160401b031681525050818181518110610e0057610e0061210f565b602002602001015160400151898781518110610e1e57610e1e61210f565b602090810291909101015260019586019501610d45565b50505b5050600101610bae565b505050509193909250565b600080516020612347833981519152610e658161120b565b6001600160401b038381161115610eb05760405162461bcd60e51b815260206004820152600f60248201526e56616c756520746f6f206c6172676560881b60448201526064016108a9565b826001600160401b0316826001600160401b031610610f205760405162461bcd60e51b815260206004820152602660248201527f5a656e4254433a2046656520657863656564732074686520616d6f756e7420746044820152651bc81b5a5b9d60d21b60648201526084016108a9565b6000805160206122e78339815191526000610f3b8486612125565b6003830154909150610f5f906001600160a01b03166001600160401b038616611340565b610f7286826001600160401b0316611340565b604080516001600160401b038084168252861660208201526001600160a01b038816917f890f2fd0578a1ba6f8735fc60c94d1ec453ad71501d82ba4d5a9041f9ac14f2e910160405180910390a2505050505050565b7f13b7ad447453d194d272cdda9bb09d7d357cda1ab7de80d865b4c1cbefc3cf28610ff28161120b565b6107cc826115f2565b61078383838360016116d1565b6000612710821161104d5760405162461bcd60e51b815260206004820152600f60248201526e15985b1d59481d1bdbc81cdb585b1b608a1b60448201526064016108a9565b6000805160206122c7833981519152546000805160206122e783398151915290600090611086908590600160a01b900461ffff166117b9565b90508381106110e55760405162461bcd60e51b815260206004820152602560248201527f5a656e4254433a2072656465656d2066656520657863656564732074686520616044820152641b5bdd5b9d60da1b60648201526084016108a9565b9392505050565b60006110f88484610994565b9050600019811461074a578181101561113d57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016108a9565b61074a848484840360006116d1565b6001600160a01b03831661117657604051634b637e8f60e11b8152600060048201526024016108a9565b6001600160a01b0382166111a05760405163ec442f0560e01b8152600060048201526024016108a9565b306001600160a01b038316036111cb5760405163ec442f0560e01b81523060048201526024016108a9565b6107838383836117d6565b6111de611914565b6111e661195d565b565b6111f0611914565b6111f86119a8565b6112006119b0565b6107838383836119c0565b6112158133611a35565b50565b600060008051602061232783398151915261123384846107d0565b6112b3576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556112693390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506104e2565b60009150506104e2565b5092915050565b60006000805160206123278339815191526112df84846107d0565b156112b3576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506104e2565b6001600160a01b03821661136a5760405163ec442f0560e01b8152600060048201526024016108a9565b6107cc600083836117d6565b6001600160a01b0382166113a057604051634b637e8f60e11b8152600060048201526024016108a9565b6107cc826000836117d6565b4363ffffffff811660009081527fc87a15b711c3ce3a8db8c4c853ae9c01bdbe93da3d2e98c03c16ae2fd0346f026020526040812080546000805160206123078339815191529392036114d25782546001840180546114439261ffff169081106114185761141861210f565b90600052602060002090600891828204019190066004029054906101000a900463ffffffff16611a6e565b8254600184018054849261ffff169081106114605761146061210f565b90600052602060002090600891828204019190066004026101000a81548163ffffffff021916908363ffffffff160217905550606461ffff168360000160009054906101000a900461ffff1660010161ffff16816114c0576114c06120f9565b845461ffff191691900661ffff161783555b82546201000090046001600160401b03168360026114ef83612145565b82546001600160401b039182166101009390930a928302928202191691909117909155604080516060810182528654620100009004831681528883166020808301918252928201898152865460018181018955600089815295909520845160029290920201805493518716600160401b026001600160801b03199094169190961617919091178455519093509082019061072490826121bb565b6000805160206122c783398151915280546001600160a01b0319166001600160a01b0383169081179091556040516000805160206122e783398151915291907f446e39bcf1b47cfadfaa23442cb4b34682cfe6bd9220da084894e3b1f834e4f390600090a25050565b61271061ffff8216111561165c5760405162461bcd60e51b815260206004820152602b60248201527f626173697320706f696e747320726174652063616e6e6f74206578636565642060448201526a31303030302075696e747360a81b60648201526084016108a9565b6000805160206122c7833981519152805461ffff8316600160a01b810261ffff60a01b199092169190911790915560408051918252516000805160206122e7833981519152917f79d2a06c43b232cf9d1835b5e915efe74621561aaf75fecec91fd75a940f9d70919081900360200190a15050565b6000805160206122e78339815191526001600160a01b03851661170a5760405163e602df0560e01b8152600060048201526024016108a9565b6001600160a01b03841661173457604051634a1406b160e11b8152600060048201526024016108a9565b6001600160a01b038086166000908152600183016020908152604080832093881683529290522083905581156117b257836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925856040516117a991815260200190565b60405180910390a35b5050505050565b60006127106117cc61ffff84168561227a565b6110e59190612291565b6000805160206122e78339815191526001600160a01b038416611812578181600201600082825461180791906122b3565b909155506118849050565b6001600160a01b038416600090815260208290526040902054828110156118655760405163391434e360e21b81526001600160a01b038616600482015260248101829052604481018490526064016108a9565b6001600160a01b03851660009081526020839052604090209083900390555b6001600160a01b0383166118a25760028101805483900390556118c1565b6001600160a01b03831660009081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161190691815260200190565b60405180910390a350505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166111e657604051631afcd79f60e31b815260040160405180910390fd5b611965611914565b604080516064808252610ca082019092526000805160206123078339815191529160208201610c8080368337505081516107cc9260018501925060200190611abc565b6111e6611914565b6119b8611914565b6111e6611aa9565b6119c8611914565b6000805160206122e78339815191527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace04611a0285826121bb565b5060058101611a1184826121bb565b50600301805460ff909216600160b01b0260ff60b01b199092169190911790555050565b611a3f82826107d0565b6107cc5760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016108a9565b60008051602061230783398151915263ffffffff8216156107cc5763ffffffff8216600090815260028201602052604081206107cc91611b6b565b611ab1611914565b611215600033611218565b82805482825590600052602060002090600701600890048101928215611b5b5791602002820160005b83821115611b2957835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302611ae5565b8015611b595782816101000a81549063ffffffff0219169055600401602081600301049283019260010302611b29565b505b50611b67929150611b8c565b5090565b50805460008255600202906000526020600020908101906112159190611ba1565b5b80821115611b675760008155600101611b8d565b80821115611b675780546001600160801b03191681556000611bc66001830182611bcf565b50600201611ba1565b508054611bdb90612064565b6000825580601f10611beb575050565b601f0160209004906000526020600020908101906112159190611b8c565b600060208284031215611c1b57600080fd5b81356001600160e01b0319811681146110e557600080fd5b6000815180845260005b81811015611c5957602081850181015186830182015201611c3d565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006110e56020830184611c33565b80356001600160a01b0381168114611ca357600080fd5b919050565b60008060408385031215611cbb57600080fd5b611cc483611c8c565b946020939093013593505050565b600060208284031215611ce457600080fd5b5035919050565b600080600060608486031215611d0057600080fd5b611d0984611c8c565b9250611d1760208501611c8c565b9150604084013590509250925092565b634e487b7160e01b600052604160045260246000fd5b60006001600160401b0380841115611d5757611d57611d27565b604051601f8501601f19908116603f01168101908282118183101715611d7f57611d7f611d27565b81604052809350858152868686011115611d9857600080fd5b858560208301376000602087830101525050509392505050565b600082601f830112611dc357600080fd5b6110e583833560208501611d3d565b600080600060608486031215611de757600080fd5b83356001600160401b0380821115611dfe57600080fd5b611e0a87838801611db2565b94506020860135915080821115611e2057600080fd5b50611e2d86828701611db2565b925050604084013560ff81168114611e4457600080fd5b809150509250925092565b60008060408385031215611e6257600080fd5b82359150611e7260208401611c8c565b90509250929050565b600060208284031215611e8d57600080fd5b6110e582611c8c565b60008060408385031215611ea957600080fd5b8235915060208301356001600160401b03811115611ec657600080fd5b8301601f81018513611ed757600080fd5b611ee685823560208401611d3d565b9150509250929050565b60008060408385031215611f0357600080fd5b611f0c83611c8c565b9150611e7260208401611c8c565b600060208284031215611f2c57600080fd5b813561ffff811681146110e557600080fd5b60008151808452602080850194506020840160005b83811015611f785781516001600160401b031687529582019590820190600101611f53565b509495945050505050565b606081526000611f966060830186611f3e565b602083820381850152611fa98287611f3e565b915083820360408501528185518084528284019150828160051b85010183880160005b83811015611ffa57601f19878403018552611fe8838351611c33565b94860194925090850190600101611fcc565b50909a9950505050505050505050565b80356001600160401b0381168114611ca357600080fd5b60008060006060848603121561203657600080fd5b61203f84611c8c565b925061204d6020850161200a565b915061205b6040850161200a565b90509250925092565b600181811c9082168061207857607f821691505b60208210810361209857634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b818103818111156104e2576104e261209e565b6001600160401b03841681526060602082015260006120e96060830185611c33565b9050826040830152949350505050565b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6001600160401b038281168282160390808211156112bd576112bd61209e565b60006001600160401b038083168181036121615761216161209e565b6001019392505050565b601f821115610783576000816000526020600020601f850160051c810160208610156121945750805b601f850160051c820191505b818110156121b3578281556001016121a0565b505050505050565b81516001600160401b038111156121d4576121d4611d27565b6121e8816121e28454612064565b8461216b565b602080601f83116001811461221d57600084156122055750858301515b600019600386901b1c1916600185901b1785556121b3565b600085815260208120601f198616915b8281101561224c5788860151825594840194600190910190840161222d565b508582101561226a5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b80820281158282048414176104e2576104e261209e565b6000826122ae57634e487b7160e01b600052601260045260246000fd5b500490565b808201808211156104e2576104e261209e56fe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0352c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00c87a15b711c3ce3a8db8c4c853ae9c01bdbe93da3d2e98c03c16ae2fd0346f0002dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800b458d13b0f4ce9e6aa65d297a27b10f75fdc6d0957bb29e1f2a30c8766b35415a2646970667358221220430110f7638d24ef0d13851dc96d2331be8b4619d666224773eafdb5a9fd111264736f6c63430008180033",
}

// ZenBTCABI is the input ABI used to generate the binding from.
// Deprecated: Use ZenBTCMetaData.ABI instead.
var ZenBTCABI = ZenBTCMetaData.ABI

// ZenBTCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZenBTCMetaData.Bin instead.
var ZenBTCBin = ZenBTCMetaData.Bin

// DeployZenBTC deploys a new Ethereum contract, binding an instance of ZenBTC to it.
func DeployZenBTC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZenBTC, error) {
	parsed, err := ZenBTCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZenBTCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZenBTC{ZenBTCCaller: ZenBTCCaller{contract: contract}, ZenBTCTransactor: ZenBTCTransactor{contract: contract}, ZenBTCFilterer: ZenBTCFilterer{contract: contract}}, nil
}

// ZenBTC is an auto generated Go binding around an Ethereum contract.
type ZenBTC struct {
	ZenBTCCaller     // Read-only binding to the contract
	ZenBTCTransactor // Write-only binding to the contract
	ZenBTCFilterer   // Log filterer for contract events
}

// ZenBTCCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZenBTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZenBTCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZenBTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZenBTCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZenBTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZenBTCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZenBTCSession struct {
	Contract     *ZenBTC           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZenBTCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZenBTCCallerSession struct {
	Contract *ZenBTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZenBTCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZenBTCTransactorSession struct {
	Contract     *ZenBTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZenBTCRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZenBTCRaw struct {
	Contract *ZenBTC // Generic contract binding to access the raw methods on
}

// ZenBTCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZenBTCCallerRaw struct {
	Contract *ZenBTCCaller // Generic read-only contract binding to access the raw methods on
}

// ZenBTCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZenBTCTransactorRaw struct {
	Contract *ZenBTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZenBTC creates a new instance of ZenBTC, bound to a specific deployed contract.
func NewZenBTC(address common.Address, backend bind.ContractBackend) (*ZenBTC, error) {
	contract, err := bindZenBTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZenBTC{ZenBTCCaller: ZenBTCCaller{contract: contract}, ZenBTCTransactor: ZenBTCTransactor{contract: contract}, ZenBTCFilterer: ZenBTCFilterer{contract: contract}}, nil
}

// NewZenBTCCaller creates a new read-only instance of ZenBTC, bound to a specific deployed contract.
func NewZenBTCCaller(address common.Address, caller bind.ContractCaller) (*ZenBTCCaller, error) {
	contract, err := bindZenBTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZenBTCCaller{contract: contract}, nil
}

// NewZenBTCTransactor creates a new write-only instance of ZenBTC, bound to a specific deployed contract.
func NewZenBTCTransactor(address common.Address, transactor bind.ContractTransactor) (*ZenBTCTransactor, error) {
	contract, err := bindZenBTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZenBTCTransactor{contract: contract}, nil
}

// NewZenBTCFilterer creates a new log filterer instance of ZenBTC, bound to a specific deployed contract.
func NewZenBTCFilterer(address common.Address, filterer bind.ContractFilterer) (*ZenBTCFilterer, error) {
	contract, err := bindZenBTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZenBTCFilterer{contract: contract}, nil
}

// bindZenBTC binds a generic wrapper to an already deployed contract.
func bindZenBTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZenBTCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZenBTC *ZenBTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZenBTC.Contract.ZenBTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZenBTC *ZenBTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZenBTC.Contract.ZenBTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZenBTC *ZenBTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZenBTC.Contract.ZenBTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZenBTC *ZenBTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZenBTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZenBTC *ZenBTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZenBTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZenBTC *ZenBTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZenBTC.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZenBTC *ZenBTCCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZenBTC *ZenBTCSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZenBTC.Contract.DEFAULTADMINROLE(&_ZenBTC.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZenBTC *ZenBTCCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZenBTC.Contract.DEFAULTADMINROLE(&_ZenBTC.CallOpts)
}

// FEEROLE is a free data retrieval call binding the contract method 0x2a0276f8.
//
// Solidity: function FEE_ROLE() view returns(bytes32)
func (_ZenBTC *ZenBTCCaller) FEEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "FEE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FEEROLE is a free data retrieval call binding the contract method 0x2a0276f8.
//
// Solidity: function FEE_ROLE() view returns(bytes32)
func (_ZenBTC *ZenBTCSession) FEEROLE() ([32]byte, error) {
	return _ZenBTC.Contract.FEEROLE(&_ZenBTC.CallOpts)
}

// FEEROLE is a free data retrieval call binding the contract method 0x2a0276f8.
//
// Solidity: function FEE_ROLE() view returns(bytes32)
func (_ZenBTC *ZenBTCCallerSession) FEEROLE() ([32]byte, error) {
	return _ZenBTC.Contract.FEEROLE(&_ZenBTC.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ZenBTC *ZenBTCCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ZenBTC *ZenBTCSession) MINTERROLE() ([32]byte, error) {
	return _ZenBTC.Contract.MINTERROLE(&_ZenBTC.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ZenBTC *ZenBTCCallerSession) MINTERROLE() ([32]byte, error) {
	return _ZenBTC.Contract.MINTERROLE(&_ZenBTC.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZenBTC *ZenBTCCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZenBTC *ZenBTCSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZenBTC.Contract.Allowance(&_ZenBTC.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZenBTC *ZenBTCCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZenBTC.Contract.Allowance(&_ZenBTC.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZenBTC *ZenBTCCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZenBTC *ZenBTCSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZenBTC.Contract.BalanceOf(&_ZenBTC.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZenBTC *ZenBTCCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZenBTC.Contract.BalanceOf(&_ZenBTC.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZenBTC *ZenBTCCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZenBTC *ZenBTCSession) Decimals() (uint8, error) {
	return _ZenBTC.Contract.Decimals(&_ZenBTC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZenBTC *ZenBTCCallerSession) Decimals() (uint8, error) {
	return _ZenBTC.Contract.Decimals(&_ZenBTC.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 value) view returns(uint256)
func (_ZenBTC *ZenBTCCaller) EstimateFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "estimateFee", value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 value) view returns(uint256)
func (_ZenBTC *ZenBTCSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _ZenBTC.Contract.EstimateFee(&_ZenBTC.CallOpts, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 value) view returns(uint256)
func (_ZenBTC *ZenBTCCallerSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _ZenBTC.Contract.EstimateFee(&_ZenBTC.CallOpts, value)
}

// GetBasisPointsRate is a free data retrieval call binding the contract method 0xfded2529.
//
// Solidity: function getBasisPointsRate() view returns(uint16)
func (_ZenBTC *ZenBTCCaller) GetBasisPointsRate(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "getBasisPointsRate")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetBasisPointsRate is a free data retrieval call binding the contract method 0xfded2529.
//
// Solidity: function getBasisPointsRate() view returns(uint16)
func (_ZenBTC *ZenBTCSession) GetBasisPointsRate() (uint16, error) {
	return _ZenBTC.Contract.GetBasisPointsRate(&_ZenBTC.CallOpts)
}

// GetBasisPointsRate is a free data retrieval call binding the contract method 0xfded2529.
//
// Solidity: function getBasisPointsRate() view returns(uint16)
func (_ZenBTC *ZenBTCCallerSession) GetBasisPointsRate() (uint16, error) {
	return _ZenBTC.Contract.GetBasisPointsRate(&_ZenBTC.CallOpts)
}

// GetFeeAddress is a free data retrieval call binding the contract method 0x4e7ceacb.
//
// Solidity: function getFeeAddress() view returns(address)
func (_ZenBTC *ZenBTCCaller) GetFeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "getFeeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeAddress is a free data retrieval call binding the contract method 0x4e7ceacb.
//
// Solidity: function getFeeAddress() view returns(address)
func (_ZenBTC *ZenBTCSession) GetFeeAddress() (common.Address, error) {
	return _ZenBTC.Contract.GetFeeAddress(&_ZenBTC.CallOpts)
}

// GetFeeAddress is a free data retrieval call binding the contract method 0x4e7ceacb.
//
// Solidity: function getFeeAddress() view returns(address)
func (_ZenBTC *ZenBTCCallerSession) GetFeeAddress() (common.Address, error) {
	return _ZenBTC.Contract.GetFeeAddress(&_ZenBTC.CallOpts)
}

// GetRecentRedemptionData is a free data retrieval call binding the contract method 0xf1690dbe.
//
// Solidity: function getRecentRedemptionData(uint16 numBlocks) view returns(uint64[] ids, uint64[] amounts, bytes[] destination)
func (_ZenBTC *ZenBTCCaller) GetRecentRedemptionData(opts *bind.CallOpts, numBlocks uint16) (struct {
	Ids         []uint64
	Amounts     []uint64
	Destination [][]byte
}, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "getRecentRedemptionData", numBlocks)

	outstruct := new(struct {
		Ids         []uint64
		Amounts     []uint64
		Destination [][]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ids = *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)
	outstruct.Destination = *abi.ConvertType(out[2], new([][]byte)).(*[][]byte)

	return *outstruct, err

}

// GetRecentRedemptionData is a free data retrieval call binding the contract method 0xf1690dbe.
//
// Solidity: function getRecentRedemptionData(uint16 numBlocks) view returns(uint64[] ids, uint64[] amounts, bytes[] destination)
func (_ZenBTC *ZenBTCSession) GetRecentRedemptionData(numBlocks uint16) (struct {
	Ids         []uint64
	Amounts     []uint64
	Destination [][]byte
}, error) {
	return _ZenBTC.Contract.GetRecentRedemptionData(&_ZenBTC.CallOpts, numBlocks)
}

// GetRecentRedemptionData is a free data retrieval call binding the contract method 0xf1690dbe.
//
// Solidity: function getRecentRedemptionData(uint16 numBlocks) view returns(uint64[] ids, uint64[] amounts, bytes[] destination)
func (_ZenBTC *ZenBTCCallerSession) GetRecentRedemptionData(numBlocks uint16) (struct {
	Ids         []uint64
	Amounts     []uint64
	Destination [][]byte
}, error) {
	return _ZenBTC.Contract.GetRecentRedemptionData(&_ZenBTC.CallOpts, numBlocks)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZenBTC *ZenBTCCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZenBTC *ZenBTCSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZenBTC.Contract.GetRoleAdmin(&_ZenBTC.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZenBTC *ZenBTCCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZenBTC.Contract.GetRoleAdmin(&_ZenBTC.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZenBTC *ZenBTCCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZenBTC *ZenBTCSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZenBTC.Contract.HasRole(&_ZenBTC.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZenBTC *ZenBTCCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZenBTC.Contract.HasRole(&_ZenBTC.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZenBTC *ZenBTCCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZenBTC *ZenBTCSession) Name() (string, error) {
	return _ZenBTC.Contract.Name(&_ZenBTC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZenBTC *ZenBTCCallerSession) Name() (string, error) {
	return _ZenBTC.Contract.Name(&_ZenBTC.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZenBTC *ZenBTCCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZenBTC *ZenBTCSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZenBTC.Contract.SupportsInterface(&_ZenBTC.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZenBTC *ZenBTCCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZenBTC.Contract.SupportsInterface(&_ZenBTC.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZenBTC *ZenBTCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZenBTC *ZenBTCSession) Symbol() (string, error) {
	return _ZenBTC.Contract.Symbol(&_ZenBTC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZenBTC *ZenBTCCallerSession) Symbol() (string, error) {
	return _ZenBTC.Contract.Symbol(&_ZenBTC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZenBTC *ZenBTCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZenBTC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZenBTC *ZenBTCSession) TotalSupply() (*big.Int, error) {
	return _ZenBTC.Contract.TotalSupply(&_ZenBTC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZenBTC *ZenBTCCallerSession) TotalSupply() (*big.Int, error) {
	return _ZenBTC.Contract.TotalSupply(&_ZenBTC.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ZenBTC *ZenBTCTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ZenBTC *ZenBTCSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.Contract.Approve(&_ZenBTC.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ZenBTC *ZenBTCTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.Contract.Approve(&_ZenBTC.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ZenBTC *ZenBTCTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ZenBTC *ZenBTCSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.Contract.Burn(&_ZenBTC.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ZenBTC *ZenBTCTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.Contract.Burn(&_ZenBTC.TransactOpts, value)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZenBTC *ZenBTCTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZenBTC.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZenBTC *ZenBTCSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZenBTC.Contract.GrantRole(&_ZenBTC.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZenBTC *ZenBTCTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZenBTC.Contract.GrantRole(&_ZenBTC.TransactOpts, role, account)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0x2a4b2fd9.
//
// Solidity: function initializeV1(string name_, string symbol_, uint8 decimals_) returns()
func (_ZenBTC *ZenBTCTransactor) InitializeV1(opts *bind.TransactOpts, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _ZenBTC.contract.Transact(opts, "initializeV1", name_, symbol_, decimals_)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0x2a4b2fd9.
//
// Solidity: function initializeV1(string name_, string symbol_, uint8 decimals_) returns()
func (_ZenBTC *ZenBTCSession) InitializeV1(name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _ZenBTC.Contract.InitializeV1(&_ZenBTC.TransactOpts, name_, symbol_, decimals_)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0x2a4b2fd9.
//
// Solidity: function initializeV1(string name_, string symbol_, uint8 decimals_) returns()
func (_ZenBTC *ZenBTCTransactorSession) InitializeV1(name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _ZenBTC.Contract.InitializeV1(&_ZenBTC.TransactOpts, name_, symbol_, decimals_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 value) returns()
func (_ZenBTC *ZenBTCTransactor) Mint(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.contract.Transact(opts, "mint", account, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 value) returns()
func (_ZenBTC *ZenBTCSession) Mint(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.Contract.Mint(&_ZenBTC.TransactOpts, account, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 value) returns()
func (_ZenBTC *ZenBTCTransactorSession) Mint(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.Contract.Mint(&_ZenBTC.TransactOpts, account, value)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZenBTC *ZenBTCTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZenBTC.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZenBTC *ZenBTCSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZenBTC.Contract.RenounceRole(&_ZenBTC.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZenBTC *ZenBTCTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZenBTC.Contract.RenounceRole(&_ZenBTC.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZenBTC *ZenBTCTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZenBTC.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZenBTC *ZenBTCSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZenBTC.Contract.RevokeRole(&_ZenBTC.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZenBTC *ZenBTCTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZenBTC.Contract.RevokeRole(&_ZenBTC.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ZenBTC *ZenBTCTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ZenBTC *ZenBTCSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.Contract.Transfer(&_ZenBTC.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ZenBTC *ZenBTCTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.Contract.Transfer(&_ZenBTC.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ZenBTC *ZenBTCTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ZenBTC *ZenBTCSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.Contract.TransferFrom(&_ZenBTC.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ZenBTC *ZenBTCTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZenBTC.Contract.TransferFrom(&_ZenBTC.TransactOpts, from, to, value)
}

// Unwrap is a paid mutator transaction binding the contract method 0xb413148e.
//
// Solidity: function unwrap(uint256 value, bytes destAddr) returns()
func (_ZenBTC *ZenBTCTransactor) Unwrap(opts *bind.TransactOpts, value *big.Int, destAddr []byte) (*types.Transaction, error) {
	return _ZenBTC.contract.Transact(opts, "unwrap", value, destAddr)
}

// Unwrap is a paid mutator transaction binding the contract method 0xb413148e.
//
// Solidity: function unwrap(uint256 value, bytes destAddr) returns()
func (_ZenBTC *ZenBTCSession) Unwrap(value *big.Int, destAddr []byte) (*types.Transaction, error) {
	return _ZenBTC.Contract.Unwrap(&_ZenBTC.TransactOpts, value, destAddr)
}

// Unwrap is a paid mutator transaction binding the contract method 0xb413148e.
//
// Solidity: function unwrap(uint256 value, bytes destAddr) returns()
func (_ZenBTC *ZenBTCTransactorSession) Unwrap(value *big.Int, destAddr []byte) (*types.Transaction, error) {
	return _ZenBTC.Contract.Unwrap(&_ZenBTC.TransactOpts, value, destAddr)
}

// UpdateBasisPointsRate is a paid mutator transaction binding the contract method 0xfd154cc7.
//
// Solidity: function updateBasisPointsRate(uint16 basisPointsRate_) returns()
func (_ZenBTC *ZenBTCTransactor) UpdateBasisPointsRate(opts *bind.TransactOpts, basisPointsRate_ uint16) (*types.Transaction, error) {
	return _ZenBTC.contract.Transact(opts, "updateBasisPointsRate", basisPointsRate_)
}

// UpdateBasisPointsRate is a paid mutator transaction binding the contract method 0xfd154cc7.
//
// Solidity: function updateBasisPointsRate(uint16 basisPointsRate_) returns()
func (_ZenBTC *ZenBTCSession) UpdateBasisPointsRate(basisPointsRate_ uint16) (*types.Transaction, error) {
	return _ZenBTC.Contract.UpdateBasisPointsRate(&_ZenBTC.TransactOpts, basisPointsRate_)
}

// UpdateBasisPointsRate is a paid mutator transaction binding the contract method 0xfd154cc7.
//
// Solidity: function updateBasisPointsRate(uint16 basisPointsRate_) returns()
func (_ZenBTC *ZenBTCTransactorSession) UpdateBasisPointsRate(basisPointsRate_ uint16) (*types.Transaction, error) {
	return _ZenBTC.Contract.UpdateBasisPointsRate(&_ZenBTC.TransactOpts, basisPointsRate_)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address feeAddress_) returns()
func (_ZenBTC *ZenBTCTransactor) UpdateFeeAddress(opts *bind.TransactOpts, feeAddress_ common.Address) (*types.Transaction, error) {
	return _ZenBTC.contract.Transact(opts, "updateFeeAddress", feeAddress_)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address feeAddress_) returns()
func (_ZenBTC *ZenBTCSession) UpdateFeeAddress(feeAddress_ common.Address) (*types.Transaction, error) {
	return _ZenBTC.Contract.UpdateFeeAddress(&_ZenBTC.TransactOpts, feeAddress_)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address feeAddress_) returns()
func (_ZenBTC *ZenBTCTransactorSession) UpdateFeeAddress(feeAddress_ common.Address) (*types.Transaction, error) {
	return _ZenBTC.Contract.UpdateFeeAddress(&_ZenBTC.TransactOpts, feeAddress_)
}

// Wrap is a paid mutator transaction binding the contract method 0xf1d96dd3.
//
// Solidity: function wrap(address account, uint64 value, uint64 fee) returns()
func (_ZenBTC *ZenBTCTransactor) Wrap(opts *bind.TransactOpts, account common.Address, value uint64, fee uint64) (*types.Transaction, error) {
	return _ZenBTC.contract.Transact(opts, "wrap", account, value, fee)
}

// Wrap is a paid mutator transaction binding the contract method 0xf1d96dd3.
//
// Solidity: function wrap(address account, uint64 value, uint64 fee) returns()
func (_ZenBTC *ZenBTCSession) Wrap(account common.Address, value uint64, fee uint64) (*types.Transaction, error) {
	return _ZenBTC.Contract.Wrap(&_ZenBTC.TransactOpts, account, value, fee)
}

// Wrap is a paid mutator transaction binding the contract method 0xf1d96dd3.
//
// Solidity: function wrap(address account, uint64 value, uint64 fee) returns()
func (_ZenBTC *ZenBTCTransactorSession) Wrap(account common.Address, value uint64, fee uint64) (*types.Transaction, error) {
	return _ZenBTC.Contract.Wrap(&_ZenBTC.TransactOpts, account, value, fee)
}

// ZenBTCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZenBTC contract.
type ZenBTCApprovalIterator struct {
	Event *ZenBTCApproval // Event containing the contract specifics and raw log

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
func (it *ZenBTCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTCApproval)
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
		it.Event = new(ZenBTCApproval)
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
func (it *ZenBTCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTCApproval represents a Approval event raised by the ZenBTC contract.
type ZenBTCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZenBTC *ZenBTCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ZenBTCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZenBTC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZenBTCApprovalIterator{contract: _ZenBTC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZenBTC *ZenBTCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZenBTCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZenBTC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTCApproval)
				if err := _ZenBTC.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZenBTC *ZenBTCFilterer) ParseApproval(log types.Log) (*ZenBTCApproval, error) {
	event := new(ZenBTCApproval)
	if err := _ZenBTC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTCBasisPointsRateUpdatedIterator is returned from FilterBasisPointsRateUpdated and is used to iterate over the raw logs and unpacked data for BasisPointsRateUpdated events raised by the ZenBTC contract.
type ZenBTCBasisPointsRateUpdatedIterator struct {
	Event *ZenBTCBasisPointsRateUpdated // Event containing the contract specifics and raw log

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
func (it *ZenBTCBasisPointsRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTCBasisPointsRateUpdated)
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
		it.Event = new(ZenBTCBasisPointsRateUpdated)
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
func (it *ZenBTCBasisPointsRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTCBasisPointsRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTCBasisPointsRateUpdated represents a BasisPointsRateUpdated event raised by the ZenBTC contract.
type ZenBTCBasisPointsRateUpdated struct {
	BasisPointsRateUpdated uint16
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterBasisPointsRateUpdated is a free log retrieval operation binding the contract event 0x79d2a06c43b232cf9d1835b5e915efe74621561aaf75fecec91fd75a940f9d70.
//
// Solidity: event BasisPointsRateUpdated(uint16 basisPointsRateUpdated)
func (_ZenBTC *ZenBTCFilterer) FilterBasisPointsRateUpdated(opts *bind.FilterOpts) (*ZenBTCBasisPointsRateUpdatedIterator, error) {

	logs, sub, err := _ZenBTC.contract.FilterLogs(opts, "BasisPointsRateUpdated")
	if err != nil {
		return nil, err
	}
	return &ZenBTCBasisPointsRateUpdatedIterator{contract: _ZenBTC.contract, event: "BasisPointsRateUpdated", logs: logs, sub: sub}, nil
}

// WatchBasisPointsRateUpdated is a free log subscription operation binding the contract event 0x79d2a06c43b232cf9d1835b5e915efe74621561aaf75fecec91fd75a940f9d70.
//
// Solidity: event BasisPointsRateUpdated(uint16 basisPointsRateUpdated)
func (_ZenBTC *ZenBTCFilterer) WatchBasisPointsRateUpdated(opts *bind.WatchOpts, sink chan<- *ZenBTCBasisPointsRateUpdated) (event.Subscription, error) {

	logs, sub, err := _ZenBTC.contract.WatchLogs(opts, "BasisPointsRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTCBasisPointsRateUpdated)
				if err := _ZenBTC.contract.UnpackLog(event, "BasisPointsRateUpdated", log); err != nil {
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

// ParseBasisPointsRateUpdated is a log parse operation binding the contract event 0x79d2a06c43b232cf9d1835b5e915efe74621561aaf75fecec91fd75a940f9d70.
//
// Solidity: event BasisPointsRateUpdated(uint16 basisPointsRateUpdated)
func (_ZenBTC *ZenBTCFilterer) ParseBasisPointsRateUpdated(log types.Log) (*ZenBTCBasisPointsRateUpdated, error) {
	event := new(ZenBTCBasisPointsRateUpdated)
	if err := _ZenBTC.contract.UnpackLog(event, "BasisPointsRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTCFeeAddressUpdatedIterator is returned from FilterFeeAddressUpdated and is used to iterate over the raw logs and unpacked data for FeeAddressUpdated events raised by the ZenBTC contract.
type ZenBTCFeeAddressUpdatedIterator struct {
	Event *ZenBTCFeeAddressUpdated // Event containing the contract specifics and raw log

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
func (it *ZenBTCFeeAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTCFeeAddressUpdated)
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
		it.Event = new(ZenBTCFeeAddressUpdated)
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
func (it *ZenBTCFeeAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTCFeeAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTCFeeAddressUpdated represents a FeeAddressUpdated event raised by the ZenBTC contract.
type ZenBTCFeeAddressUpdated struct {
	NewFeeAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeAddressUpdated is a free log retrieval operation binding the contract event 0x446e39bcf1b47cfadfaa23442cb4b34682cfe6bd9220da084894e3b1f834e4f3.
//
// Solidity: event FeeAddressUpdated(address indexed newFeeAddress)
func (_ZenBTC *ZenBTCFilterer) FilterFeeAddressUpdated(opts *bind.FilterOpts, newFeeAddress []common.Address) (*ZenBTCFeeAddressUpdatedIterator, error) {

	var newFeeAddressRule []interface{}
	for _, newFeeAddressItem := range newFeeAddress {
		newFeeAddressRule = append(newFeeAddressRule, newFeeAddressItem)
	}

	logs, sub, err := _ZenBTC.contract.FilterLogs(opts, "FeeAddressUpdated", newFeeAddressRule)
	if err != nil {
		return nil, err
	}
	return &ZenBTCFeeAddressUpdatedIterator{contract: _ZenBTC.contract, event: "FeeAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeAddressUpdated is a free log subscription operation binding the contract event 0x446e39bcf1b47cfadfaa23442cb4b34682cfe6bd9220da084894e3b1f834e4f3.
//
// Solidity: event FeeAddressUpdated(address indexed newFeeAddress)
func (_ZenBTC *ZenBTCFilterer) WatchFeeAddressUpdated(opts *bind.WatchOpts, sink chan<- *ZenBTCFeeAddressUpdated, newFeeAddress []common.Address) (event.Subscription, error) {

	var newFeeAddressRule []interface{}
	for _, newFeeAddressItem := range newFeeAddress {
		newFeeAddressRule = append(newFeeAddressRule, newFeeAddressItem)
	}

	logs, sub, err := _ZenBTC.contract.WatchLogs(opts, "FeeAddressUpdated", newFeeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTCFeeAddressUpdated)
				if err := _ZenBTC.contract.UnpackLog(event, "FeeAddressUpdated", log); err != nil {
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

// ParseFeeAddressUpdated is a log parse operation binding the contract event 0x446e39bcf1b47cfadfaa23442cb4b34682cfe6bd9220da084894e3b1f834e4f3.
//
// Solidity: event FeeAddressUpdated(address indexed newFeeAddress)
func (_ZenBTC *ZenBTCFilterer) ParseFeeAddressUpdated(log types.Log) (*ZenBTCFeeAddressUpdated, error) {
	event := new(ZenBTCFeeAddressUpdated)
	if err := _ZenBTC.contract.UnpackLog(event, "FeeAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTCInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZenBTC contract.
type ZenBTCInitializedIterator struct {
	Event *ZenBTCInitialized // Event containing the contract specifics and raw log

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
func (it *ZenBTCInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTCInitialized)
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
		it.Event = new(ZenBTCInitialized)
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
func (it *ZenBTCInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTCInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTCInitialized represents a Initialized event raised by the ZenBTC contract.
type ZenBTCInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZenBTC *ZenBTCFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZenBTCInitializedIterator, error) {

	logs, sub, err := _ZenBTC.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZenBTCInitializedIterator{contract: _ZenBTC.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZenBTC *ZenBTCFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZenBTCInitialized) (event.Subscription, error) {

	logs, sub, err := _ZenBTC.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTCInitialized)
				if err := _ZenBTC.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ZenBTC *ZenBTCFilterer) ParseInitialized(log types.Log) (*ZenBTCInitialized, error) {
	event := new(ZenBTCInitialized)
	if err := _ZenBTC.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTCRedeemFeeUpdatedIterator is returned from FilterRedeemFeeUpdated and is used to iterate over the raw logs and unpacked data for RedeemFeeUpdated events raised by the ZenBTC contract.
type ZenBTCRedeemFeeUpdatedIterator struct {
	Event *ZenBTCRedeemFeeUpdated // Event containing the contract specifics and raw log

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
func (it *ZenBTCRedeemFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTCRedeemFeeUpdated)
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
		it.Event = new(ZenBTCRedeemFeeUpdated)
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
func (it *ZenBTCRedeemFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTCRedeemFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTCRedeemFeeUpdated represents a RedeemFeeUpdated event raised by the ZenBTC contract.
type ZenBTCRedeemFeeUpdated struct {
	NewRedeemFee *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeemFeeUpdated is a free log retrieval operation binding the contract event 0xd6c7508d6658ccee36b7b7d7fd72e5cbaeefb40c64eff24e9ae7470e846304ee.
//
// Solidity: event RedeemFeeUpdated(uint256 newRedeemFee)
func (_ZenBTC *ZenBTCFilterer) FilterRedeemFeeUpdated(opts *bind.FilterOpts) (*ZenBTCRedeemFeeUpdatedIterator, error) {

	logs, sub, err := _ZenBTC.contract.FilterLogs(opts, "RedeemFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &ZenBTCRedeemFeeUpdatedIterator{contract: _ZenBTC.contract, event: "RedeemFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchRedeemFeeUpdated is a free log subscription operation binding the contract event 0xd6c7508d6658ccee36b7b7d7fd72e5cbaeefb40c64eff24e9ae7470e846304ee.
//
// Solidity: event RedeemFeeUpdated(uint256 newRedeemFee)
func (_ZenBTC *ZenBTCFilterer) WatchRedeemFeeUpdated(opts *bind.WatchOpts, sink chan<- *ZenBTCRedeemFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _ZenBTC.contract.WatchLogs(opts, "RedeemFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTCRedeemFeeUpdated)
				if err := _ZenBTC.contract.UnpackLog(event, "RedeemFeeUpdated", log); err != nil {
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

// ParseRedeemFeeUpdated is a log parse operation binding the contract event 0xd6c7508d6658ccee36b7b7d7fd72e5cbaeefb40c64eff24e9ae7470e846304ee.
//
// Solidity: event RedeemFeeUpdated(uint256 newRedeemFee)
func (_ZenBTC *ZenBTCFilterer) ParseRedeemFeeUpdated(log types.Log) (*ZenBTCRedeemFeeUpdated, error) {
	event := new(ZenBTCRedeemFeeUpdated)
	if err := _ZenBTC.contract.UnpackLog(event, "RedeemFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTCRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZenBTC contract.
type ZenBTCRoleAdminChangedIterator struct {
	Event *ZenBTCRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZenBTCRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTCRoleAdminChanged)
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
		it.Event = new(ZenBTCRoleAdminChanged)
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
func (it *ZenBTCRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTCRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTCRoleAdminChanged represents a RoleAdminChanged event raised by the ZenBTC contract.
type ZenBTCRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZenBTC *ZenBTCFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZenBTCRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ZenBTC.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZenBTCRoleAdminChangedIterator{contract: _ZenBTC.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZenBTC *ZenBTCFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZenBTCRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ZenBTC.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTCRoleAdminChanged)
				if err := _ZenBTC.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ZenBTC *ZenBTCFilterer) ParseRoleAdminChanged(log types.Log) (*ZenBTCRoleAdminChanged, error) {
	event := new(ZenBTCRoleAdminChanged)
	if err := _ZenBTC.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTCRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZenBTC contract.
type ZenBTCRoleGrantedIterator struct {
	Event *ZenBTCRoleGranted // Event containing the contract specifics and raw log

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
func (it *ZenBTCRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTCRoleGranted)
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
		it.Event = new(ZenBTCRoleGranted)
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
func (it *ZenBTCRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTCRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTCRoleGranted represents a RoleGranted event raised by the ZenBTC contract.
type ZenBTCRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZenBTC *ZenBTCFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZenBTCRoleGrantedIterator, error) {

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

	logs, sub, err := _ZenBTC.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZenBTCRoleGrantedIterator{contract: _ZenBTC.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZenBTC *ZenBTCFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZenBTCRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ZenBTC.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTCRoleGranted)
				if err := _ZenBTC.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ZenBTC *ZenBTCFilterer) ParseRoleGranted(log types.Log) (*ZenBTCRoleGranted, error) {
	event := new(ZenBTCRoleGranted)
	if err := _ZenBTC.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTCRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZenBTC contract.
type ZenBTCRoleRevokedIterator struct {
	Event *ZenBTCRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ZenBTCRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTCRoleRevoked)
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
		it.Event = new(ZenBTCRoleRevoked)
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
func (it *ZenBTCRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTCRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTCRoleRevoked represents a RoleRevoked event raised by the ZenBTC contract.
type ZenBTCRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZenBTC *ZenBTCFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZenBTCRoleRevokedIterator, error) {

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

	logs, sub, err := _ZenBTC.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZenBTCRoleRevokedIterator{contract: _ZenBTC.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZenBTC *ZenBTCFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZenBTCRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ZenBTC.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTCRoleRevoked)
				if err := _ZenBTC.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ZenBTC *ZenBTCFilterer) ParseRoleRevoked(log types.Log) (*ZenBTCRoleRevoked, error) {
	event := new(ZenBTCRoleRevoked)
	if err := _ZenBTC.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTCTokenRedemptionIterator is returned from FilterTokenRedemption and is used to iterate over the raw logs and unpacked data for TokenRedemption events raised by the ZenBTC contract.
type ZenBTCTokenRedemptionIterator struct {
	Event *ZenBTCTokenRedemption // Event containing the contract specifics and raw log

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
func (it *ZenBTCTokenRedemptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTCTokenRedemption)
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
		it.Event = new(ZenBTCTokenRedemption)
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
func (it *ZenBTCTokenRedemptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTCTokenRedemptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTCTokenRedemption represents a TokenRedemption event raised by the ZenBTC contract.
type ZenBTCTokenRedemption struct {
	Redeemer common.Address
	Value    uint64
	DestAddr []byte
	Fee      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenRedemption is a free log retrieval operation binding the contract event 0x4c971c8b2abb197a17896b2fc57f597830db78d3556831e3faa337a596150f22.
//
// Solidity: event TokenRedemption(address indexed redeemer, uint64 value, bytes destAddr, uint256 fee)
func (_ZenBTC *ZenBTCFilterer) FilterTokenRedemption(opts *bind.FilterOpts, redeemer []common.Address) (*ZenBTCTokenRedemptionIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}

	logs, sub, err := _ZenBTC.contract.FilterLogs(opts, "TokenRedemption", redeemerRule)
	if err != nil {
		return nil, err
	}
	return &ZenBTCTokenRedemptionIterator{contract: _ZenBTC.contract, event: "TokenRedemption", logs: logs, sub: sub}, nil
}

// WatchTokenRedemption is a free log subscription operation binding the contract event 0x4c971c8b2abb197a17896b2fc57f597830db78d3556831e3faa337a596150f22.
//
// Solidity: event TokenRedemption(address indexed redeemer, uint64 value, bytes destAddr, uint256 fee)
func (_ZenBTC *ZenBTCFilterer) WatchTokenRedemption(opts *bind.WatchOpts, sink chan<- *ZenBTCTokenRedemption, redeemer []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}

	logs, sub, err := _ZenBTC.contract.WatchLogs(opts, "TokenRedemption", redeemerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTCTokenRedemption)
				if err := _ZenBTC.contract.UnpackLog(event, "TokenRedemption", log); err != nil {
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

// ParseTokenRedemption is a log parse operation binding the contract event 0x4c971c8b2abb197a17896b2fc57f597830db78d3556831e3faa337a596150f22.
//
// Solidity: event TokenRedemption(address indexed redeemer, uint64 value, bytes destAddr, uint256 fee)
func (_ZenBTC *ZenBTCFilterer) ParseTokenRedemption(log types.Log) (*ZenBTCTokenRedemption, error) {
	event := new(ZenBTCTokenRedemption)
	if err := _ZenBTC.contract.UnpackLog(event, "TokenRedemption", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTCTokensMintedWithFeeIterator is returned from FilterTokensMintedWithFee and is used to iterate over the raw logs and unpacked data for TokensMintedWithFee events raised by the ZenBTC contract.
type ZenBTCTokensMintedWithFeeIterator struct {
	Event *ZenBTCTokensMintedWithFee // Event containing the contract specifics and raw log

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
func (it *ZenBTCTokensMintedWithFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTCTokensMintedWithFee)
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
		it.Event = new(ZenBTCTokensMintedWithFee)
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
func (it *ZenBTCTokensMintedWithFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTCTokensMintedWithFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTCTokensMintedWithFee represents a TokensMintedWithFee event raised by the ZenBTC contract.
type ZenBTCTokensMintedWithFee struct {
	Recipient common.Address
	Value     uint64
	Fee       uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensMintedWithFee is a free log retrieval operation binding the contract event 0x890f2fd0578a1ba6f8735fc60c94d1ec453ad71501d82ba4d5a9041f9ac14f2e.
//
// Solidity: event TokensMintedWithFee(address indexed recipient, uint64 value, uint64 fee)
func (_ZenBTC *ZenBTCFilterer) FilterTokensMintedWithFee(opts *bind.FilterOpts, recipient []common.Address) (*ZenBTCTokensMintedWithFeeIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ZenBTC.contract.FilterLogs(opts, "TokensMintedWithFee", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ZenBTCTokensMintedWithFeeIterator{contract: _ZenBTC.contract, event: "TokensMintedWithFee", logs: logs, sub: sub}, nil
}

// WatchTokensMintedWithFee is a free log subscription operation binding the contract event 0x890f2fd0578a1ba6f8735fc60c94d1ec453ad71501d82ba4d5a9041f9ac14f2e.
//
// Solidity: event TokensMintedWithFee(address indexed recipient, uint64 value, uint64 fee)
func (_ZenBTC *ZenBTCFilterer) WatchTokensMintedWithFee(opts *bind.WatchOpts, sink chan<- *ZenBTCTokensMintedWithFee, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ZenBTC.contract.WatchLogs(opts, "TokensMintedWithFee", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTCTokensMintedWithFee)
				if err := _ZenBTC.contract.UnpackLog(event, "TokensMintedWithFee", log); err != nil {
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

// ParseTokensMintedWithFee is a log parse operation binding the contract event 0x890f2fd0578a1ba6f8735fc60c94d1ec453ad71501d82ba4d5a9041f9ac14f2e.
//
// Solidity: event TokensMintedWithFee(address indexed recipient, uint64 value, uint64 fee)
func (_ZenBTC *ZenBTCFilterer) ParseTokensMintedWithFee(log types.Log) (*ZenBTCTokensMintedWithFee, error) {
	event := new(ZenBTCTokensMintedWithFee)
	if err := _ZenBTC.contract.UnpackLog(event, "TokensMintedWithFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZenBTCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZenBTC contract.
type ZenBTCTransferIterator struct {
	Event *ZenBTCTransfer // Event containing the contract specifics and raw log

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
func (it *ZenBTCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZenBTCTransfer)
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
		it.Event = new(ZenBTCTransfer)
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
func (it *ZenBTCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZenBTCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZenBTCTransfer represents a Transfer event raised by the ZenBTC contract.
type ZenBTCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZenBTC *ZenBTCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZenBTCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZenBTC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZenBTCTransferIterator{contract: _ZenBTC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZenBTC *ZenBTCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZenBTCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZenBTC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZenBTCTransfer)
				if err := _ZenBTC.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZenBTC *ZenBTCFilterer) ParseTransfer(log types.Log) (*ZenBTCTransfer, error) {
	event := new(ZenBTCTransfer)
	if err := _ZenBTC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}