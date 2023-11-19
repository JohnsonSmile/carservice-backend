// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package user

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

// UserUserEntity is an auto generated low-level Go binding around an user-defined struct.
type UserUserEntity struct {
	Id    *big.Int
	Phone *big.Int
	Score *big.Int
}

// UserMetaData contains all meta data concerning the User contract.
var UserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Error_UserAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Error_UserInfoParamsInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Error_UserNotExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"userId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"phone\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structUser.UserEntity\",\"name\":\"user\",\"type\":\"tuple\"}],\"name\":\"UserCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"userId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"phone\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structUser.UserEntity\",\"name\":\"user\",\"type\":\"tuple\"}],\"name\":\"UserUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_score\",\"type\":\"uint256\"}],\"name\":\"ChargeScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_phone\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_score\",\"type\":\"uint256\"}],\"name\":\"CreateUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_userId\",\"type\":\"uint256\"}],\"name\":\"GetUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"phone\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"internalType\":\"structUser.UserEntity\",\"name\":\"user\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_score\",\"type\":\"uint256\"}],\"name\":\"PayScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_phone\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_score\",\"type\":\"uint256\"}],\"name\":\"UpdateUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040818152346100e7577faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c90600091808352602090838252828420338552825260ff8385205416156100af575b50828052828152818320338452815260ff828420541615610076575b610cc084816100ed8239f35b82805282815281832090338452528120600160ff1982541617905533903390600080516020610dad8339815191528180a438808061006a565b8084528382528284203385528252828420600160ff1982541617905533903390600080516020610dad8339815191528680a43861004e565b600080fdfe608060408181526004918236101561001657600080fd5b600092833560e01c91826301ffc9a71461067357508163070f1d52146105d9578163148c32b7146105345781631b2df850146104f9578163210755c914610432578163248a9ca3146104085781632f2ff15d1461036a57816336568abe146102e25781633eeb4a871461022b57816391d14854146101f1578163a217fddf146101d6578163c57e13cf146100ea575063d547741f146100b457600080fd5b346100e65760016100e3916100de6100cb366106e1565b93909283875286602052862001546109c3565b610af1565b80f35b5080fd5b919050346101d2576100fb36610707565b929193610106610791565b841580156101ca575b6101bc578486526001602052828620546101ae575091817f17a87a2888b2cae7ae946db8c992fa2329cb5a2152c5578a293bcbd8bb2719a59360026101a894519361015985610721565b8785526020850192835283850190815287895260016020528389209285518455516001840155519101555191829182919091604080606083019480518452602081015160208501520151910152565b0390a280f35b82516344b5a60160e11b8152fd5b825163aba49a8d60e01b8152fd5b50811561010f565b8280fd5b5050346100e657816003193601126100e65751908152602090f35b5050346100e65760ff81602093610207366106e1565b9082528186528282206001600160a01b039091168252855220549151911615158152f35b919050346101d25761023c36610707565b90929193610248610791565b848652600160205282862054156102d45750916101a8917f9049d2447d6eb5adf4d038b979f34c6950f566b8bf3244d4247834bd9b1dad8293858752600160205280600283892085600182015501558151926102a384610721565b8684526020840152818301525191829182919091604080606083019480518452602081015160208501520151910152565b8251633e88ae7f60e11b8152fd5b839150346100e6576102f3366106e1565b91336001600160a01b0384160361030f5750906100e391610af1565b608490602086519162461bcd60e51b8352820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152fd5b5050346100e65761037a366106e1565b909182845283602052610392600182862001546109c3565b82845260208481528185206001600160a01b039093168086529290528084205460ff16156103be578380f35b828452836020528084208285526020528320600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8480a43880808380f35b9050346101d25760203660031901126101d257816020936001923581528085522001549051908152f35b919050346101d257610443366106c6565b91909261044e610791565b838552600160205281852090825161046581610721565b82548082528460026001860154956020850196875201549201918252156104ea57519384018094116104d75750916101a8917f9049d2447d6eb5adf4d038b979f34c6950f566b8bf3244d4247834bd9b1dad82938587526001602052826002838920015551918151926102a384610721565b634e487b7160e01b865260119052602485fd5b508251633e88ae7f60e11b8152fd5b5050346100e657816003193601126100e657602090517faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c8152f35b919050346101d257610545366106c6565b919092610550610791565b838552600160205281852090825161056781610721565b82548082528460026001860154956020850196875201549201918252156104ea57519384039384116104d75750916101a8917f9049d2447d6eb5adf4d038b979f34c6950f566b8bf3244d4247834bd9b1dad82938587526001602052826002838920015551918151926102a384610721565b9050346101d25760203660031901126101d257803590838380516105fc81610721565b8281528260208201520152818452600160205282842054156102d4575082829161066f945260016020522090600281519261063684610721565b80548452600181015460208501520154818301525191829182919091604080606083019480518452602081015160208501520151910152565b0390f35b8491346101d25760203660031901126101d2573563ffffffff60e01b81168091036101d25760209250637965db0b60e01b81149081156106b5575b5015158152f35b6301ffc9a760e01b149050836106ae565b60409060031901126106dc576004359060243590565b600080fd5b60409060031901126106dc57600435906024356001600160a01b03811681036106dc5790565b60609060031901126106dc57600435906024359060443590565b6060810190811067ffffffffffffffff82111761073d57604052565b634e487b7160e01b600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761073d57604052565b90601f8019910116810190811067ffffffffffffffff82111761073d57604052565b3360009081527f511473bfc0317b2ecebc3e7070288e5dde2957151ef598d1c2d01ac136128f326020908152604080832054909291907faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c9060ff16156107f75750505050565b61080033610b8c565b9084519061080d82610753565b604282528382019460603687378251156109af57603086538251906001918210156109af5790607860218501536041915b818311610941575050506108ff57846108cd60486108f19360449798519889916108be8984019876020b1b1b2b9b9a1b7b73a3937b61d1030b1b1b7bab73a1604d1b8a52610895815180928d603789019101610ace565b8401917001034b99036b4b9b9b4b733903937b6329607d1b603784015251809386840190610ace565b0103602881018952018761076f565b5194859362461bcd60e51b8552600485015251809281602486015285850190610ace565b601f01601f19168101030190fd5b60648386519062461bcd60e51b825280600483015260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152fd5b909192600f8116601081101561099b576f181899199a1a9b1b9c1cb0b131b232b360811b901a6109718587610b65565b5360041c9280156109875760001901919061083e565b634e487b7160e01b82526011600452602482fd5b634e487b7160e01b83526032600452602483fd5b634e487b7160e01b81526032600452602490fd5b60008181526020818152604092838320338452825260ff8484205416156109ea5750505050565b6109f333610b8c565b90845190610a0082610753565b604282528382019460603687378251156109af57603086538251906001918210156109af5790607860218501536041915b818311610a88575050506108ff57846108cd60486108f19360449798519889916108be8984019876020b1b1b2b9b9a1b7b73a3937b61d1030b1b1b7bab73a1604d1b8a52610895815180928d603789019101610ace565b909192600f8116601081101561099b576f181899199a1a9b1b9c1cb0b131b232b360811b901a610ab88587610b65565b5360041c92801561098757600019019190610a31565b60005b838110610ae15750506000910152565b8181015183820152602001610ad1565b9060009180835282602052604083209160018060a01b03169182845260205260ff604084205416610b2157505050565b80835282602052604083208284526020526040832060ff1981541690557ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b339380a4565b908151811015610b76570160200190565b634e487b7160e01b600052603260045260246000fd5b60405190610b9982610721565b602a8252602082016040368237825115610b7657603090538151600190811015610b7657607860218401536029905b808211610c1c575050610bd85790565b606460405162461bcd60e51b815260206004820152602060248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152fd5b9091600f81166010811015610c75576f181899199a1a9b1b9c1cb0b131b232b360811b901a610c4b8486610b65565b5360041c918015610c60576000190190610bc8565b60246000634e487b7160e01b81526011600452fd5b60246000634e487b7160e01b81526032600452fdfea2646970667358221220a856bdce3037a5dcb1479512134e0db8fc8c3ae4086be0b0f27612fdbc941a3264736f6c634300081300332f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d",
}

// UserABI is the input ABI used to generate the binding from.
// Deprecated: Use UserMetaData.ABI instead.
var UserABI = UserMetaData.ABI

// UserBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UserMetaData.Bin instead.
var UserBin = UserMetaData.Bin

// DeployUser deploys a new Ethereum contract, binding an instance of User to it.
func DeployUser(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *User, error) {
	parsed, err := UserMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UserBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}, UserFilterer: UserFilterer{contract: contract}}, nil
}

// User is an auto generated Go binding around an Ethereum contract.
type User struct {
	UserCaller     // Read-only binding to the contract
	UserTransactor // Write-only binding to the contract
	UserFilterer   // Log filterer for contract events
}

// UserCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserSession struct {
	Contract     *User             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserCallerSession struct {
	Contract *UserCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserTransactorSession struct {
	Contract     *UserTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserRaw struct {
	Contract *User // Generic contract binding to access the raw methods on
}

// UserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserCallerRaw struct {
	Contract *UserCaller // Generic read-only contract binding to access the raw methods on
}

// UserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserTransactorRaw struct {
	Contract *UserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUser creates a new instance of User, bound to a specific deployed contract.
func NewUser(address common.Address, backend bind.ContractBackend) (*User, error) {
	contract, err := bindUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}, UserFilterer: UserFilterer{contract: contract}}, nil
}

// NewUserCaller creates a new read-only instance of User, bound to a specific deployed contract.
func NewUserCaller(address common.Address, caller bind.ContractCaller) (*UserCaller, error) {
	contract, err := bindUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserCaller{contract: contract}, nil
}

// NewUserTransactor creates a new write-only instance of User, bound to a specific deployed contract.
func NewUserTransactor(address common.Address, transactor bind.ContractTransactor) (*UserTransactor, error) {
	contract, err := bindUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserTransactor{contract: contract}, nil
}

// NewUserFilterer creates a new log filterer instance of User, bound to a specific deployed contract.
func NewUserFilterer(address common.Address, filterer bind.ContractFilterer) (*UserFilterer, error) {
	contract, err := bindUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserFilterer{contract: contract}, nil
}

// bindUser binds a generic wrapper to an already deployed contract.
func bindUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UserMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.UserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_User *UserCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_User *UserSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _User.Contract.DEFAULTADMINROLE(&_User.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_User *UserCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _User.Contract.DEFAULTADMINROLE(&_User.CallOpts)
}

// GetUser is a free data retrieval call binding the contract method 0x070f1d52.
//
// Solidity: function GetUser(uint256 _userId) view returns((uint256,uint256,uint256) user)
func (_User *UserCaller) GetUser(opts *bind.CallOpts, _userId *big.Int) (UserUserEntity, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "GetUser", _userId)

	if err != nil {
		return *new(UserUserEntity), err
	}

	out0 := *abi.ConvertType(out[0], new(UserUserEntity)).(*UserUserEntity)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x070f1d52.
//
// Solidity: function GetUser(uint256 _userId) view returns((uint256,uint256,uint256) user)
func (_User *UserSession) GetUser(_userId *big.Int) (UserUserEntity, error) {
	return _User.Contract.GetUser(&_User.CallOpts, _userId)
}

// GetUser is a free data retrieval call binding the contract method 0x070f1d52.
//
// Solidity: function GetUser(uint256 _userId) view returns((uint256,uint256,uint256) user)
func (_User *UserCallerSession) GetUser(_userId *big.Int) (UserUserEntity, error) {
	return _User.Contract.GetUser(&_User.CallOpts, _userId)
}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_User *UserCaller) MANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_User *UserSession) MANAGER() ([32]byte, error) {
	return _User.Contract.MANAGER(&_User.CallOpts)
}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_User *UserCallerSession) MANAGER() ([32]byte, error) {
	return _User.Contract.MANAGER(&_User.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_User *UserCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_User *UserSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _User.Contract.GetRoleAdmin(&_User.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_User *UserCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _User.Contract.GetRoleAdmin(&_User.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_User *UserCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_User *UserSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _User.Contract.HasRole(&_User.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_User *UserCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _User.Contract.HasRole(&_User.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_User *UserCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_User *UserSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _User.Contract.SupportsInterface(&_User.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_User *UserCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _User.Contract.SupportsInterface(&_User.CallOpts, interfaceId)
}

// ChargeScore is a paid mutator transaction binding the contract method 0x210755c9.
//
// Solidity: function ChargeScore(uint256 _id, uint256 _score) returns()
func (_User *UserTransactor) ChargeScore(opts *bind.TransactOpts, _id *big.Int, _score *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "ChargeScore", _id, _score)
}

// ChargeScore is a paid mutator transaction binding the contract method 0x210755c9.
//
// Solidity: function ChargeScore(uint256 _id, uint256 _score) returns()
func (_User *UserSession) ChargeScore(_id *big.Int, _score *big.Int) (*types.Transaction, error) {
	return _User.Contract.ChargeScore(&_User.TransactOpts, _id, _score)
}

// ChargeScore is a paid mutator transaction binding the contract method 0x210755c9.
//
// Solidity: function ChargeScore(uint256 _id, uint256 _score) returns()
func (_User *UserTransactorSession) ChargeScore(_id *big.Int, _score *big.Int) (*types.Transaction, error) {
	return _User.Contract.ChargeScore(&_User.TransactOpts, _id, _score)
}

// CreateUser is a paid mutator transaction binding the contract method 0xc57e13cf.
//
// Solidity: function CreateUser(uint256 _id, uint256 _phone, uint256 _score) returns()
func (_User *UserTransactor) CreateUser(opts *bind.TransactOpts, _id *big.Int, _phone *big.Int, _score *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "CreateUser", _id, _phone, _score)
}

// CreateUser is a paid mutator transaction binding the contract method 0xc57e13cf.
//
// Solidity: function CreateUser(uint256 _id, uint256 _phone, uint256 _score) returns()
func (_User *UserSession) CreateUser(_id *big.Int, _phone *big.Int, _score *big.Int) (*types.Transaction, error) {
	return _User.Contract.CreateUser(&_User.TransactOpts, _id, _phone, _score)
}

// CreateUser is a paid mutator transaction binding the contract method 0xc57e13cf.
//
// Solidity: function CreateUser(uint256 _id, uint256 _phone, uint256 _score) returns()
func (_User *UserTransactorSession) CreateUser(_id *big.Int, _phone *big.Int, _score *big.Int) (*types.Transaction, error) {
	return _User.Contract.CreateUser(&_User.TransactOpts, _id, _phone, _score)
}

// PayScore is a paid mutator transaction binding the contract method 0x148c32b7.
//
// Solidity: function PayScore(uint256 _id, uint256 _score) returns()
func (_User *UserTransactor) PayScore(opts *bind.TransactOpts, _id *big.Int, _score *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "PayScore", _id, _score)
}

// PayScore is a paid mutator transaction binding the contract method 0x148c32b7.
//
// Solidity: function PayScore(uint256 _id, uint256 _score) returns()
func (_User *UserSession) PayScore(_id *big.Int, _score *big.Int) (*types.Transaction, error) {
	return _User.Contract.PayScore(&_User.TransactOpts, _id, _score)
}

// PayScore is a paid mutator transaction binding the contract method 0x148c32b7.
//
// Solidity: function PayScore(uint256 _id, uint256 _score) returns()
func (_User *UserTransactorSession) PayScore(_id *big.Int, _score *big.Int) (*types.Transaction, error) {
	return _User.Contract.PayScore(&_User.TransactOpts, _id, _score)
}

// UpdateUser is a paid mutator transaction binding the contract method 0x3eeb4a87.
//
// Solidity: function UpdateUser(uint256 _id, uint256 _phone, uint256 _score) returns()
func (_User *UserTransactor) UpdateUser(opts *bind.TransactOpts, _id *big.Int, _phone *big.Int, _score *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "UpdateUser", _id, _phone, _score)
}

// UpdateUser is a paid mutator transaction binding the contract method 0x3eeb4a87.
//
// Solidity: function UpdateUser(uint256 _id, uint256 _phone, uint256 _score) returns()
func (_User *UserSession) UpdateUser(_id *big.Int, _phone *big.Int, _score *big.Int) (*types.Transaction, error) {
	return _User.Contract.UpdateUser(&_User.TransactOpts, _id, _phone, _score)
}

// UpdateUser is a paid mutator transaction binding the contract method 0x3eeb4a87.
//
// Solidity: function UpdateUser(uint256 _id, uint256 _phone, uint256 _score) returns()
func (_User *UserTransactorSession) UpdateUser(_id *big.Int, _phone *big.Int, _score *big.Int) (*types.Transaction, error) {
	return _User.Contract.UpdateUser(&_User.TransactOpts, _id, _phone, _score)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_User *UserTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_User *UserSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _User.Contract.GrantRole(&_User.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_User *UserTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _User.Contract.GrantRole(&_User.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_User *UserTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_User *UserSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _User.Contract.RenounceRole(&_User.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_User *UserTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _User.Contract.RenounceRole(&_User.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_User *UserTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_User *UserSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _User.Contract.RevokeRole(&_User.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_User *UserTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _User.Contract.RevokeRole(&_User.TransactOpts, role, account)
}

// UserRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the User contract.
type UserRoleAdminChangedIterator struct {
	Event *UserRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *UserRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserRoleAdminChanged)
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
		it.Event = new(UserRoleAdminChanged)
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
func (it *UserRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserRoleAdminChanged represents a RoleAdminChanged event raised by the User contract.
type UserRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_User *UserFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*UserRoleAdminChangedIterator, error) {

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

	logs, sub, err := _User.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &UserRoleAdminChangedIterator{contract: _User.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_User *UserFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *UserRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _User.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserRoleAdminChanged)
				if err := _User.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_User *UserFilterer) ParseRoleAdminChanged(log types.Log) (*UserRoleAdminChanged, error) {
	event := new(UserRoleAdminChanged)
	if err := _User.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the User contract.
type UserRoleGrantedIterator struct {
	Event *UserRoleGranted // Event containing the contract specifics and raw log

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
func (it *UserRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserRoleGranted)
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
		it.Event = new(UserRoleGranted)
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
func (it *UserRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserRoleGranted represents a RoleGranted event raised by the User contract.
type UserRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_User *UserFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UserRoleGrantedIterator, error) {

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

	logs, sub, err := _User.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UserRoleGrantedIterator{contract: _User.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_User *UserFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *UserRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _User.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserRoleGranted)
				if err := _User.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_User *UserFilterer) ParseRoleGranted(log types.Log) (*UserRoleGranted, error) {
	event := new(UserRoleGranted)
	if err := _User.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the User contract.
type UserRoleRevokedIterator struct {
	Event *UserRoleRevoked // Event containing the contract specifics and raw log

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
func (it *UserRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserRoleRevoked)
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
		it.Event = new(UserRoleRevoked)
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
func (it *UserRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserRoleRevoked represents a RoleRevoked event raised by the User contract.
type UserRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_User *UserFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UserRoleRevokedIterator, error) {

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

	logs, sub, err := _User.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UserRoleRevokedIterator{contract: _User.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_User *UserFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *UserRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _User.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserRoleRevoked)
				if err := _User.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_User *UserFilterer) ParseRoleRevoked(log types.Log) (*UserRoleRevoked, error) {
	event := new(UserRoleRevoked)
	if err := _User.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserUserCreatedIterator is returned from FilterUserCreated and is used to iterate over the raw logs and unpacked data for UserCreated events raised by the User contract.
type UserUserCreatedIterator struct {
	Event *UserUserCreated // Event containing the contract specifics and raw log

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
func (it *UserUserCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserUserCreated)
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
		it.Event = new(UserUserCreated)
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
func (it *UserUserCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserUserCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserUserCreated represents a UserCreated event raised by the User contract.
type UserUserCreated struct {
	UserId *big.Int
	User   UserUserEntity
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUserCreated is a free log retrieval operation binding the contract event 0x17a87a2888b2cae7ae946db8c992fa2329cb5a2152c5578a293bcbd8bb2719a5.
//
// Solidity: event UserCreated(uint256 indexed userId, (uint256,uint256,uint256) user)
func (_User *UserFilterer) FilterUserCreated(opts *bind.FilterOpts, userId []*big.Int) (*UserUserCreatedIterator, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _User.contract.FilterLogs(opts, "UserCreated", userIdRule)
	if err != nil {
		return nil, err
	}
	return &UserUserCreatedIterator{contract: _User.contract, event: "UserCreated", logs: logs, sub: sub}, nil
}

// WatchUserCreated is a free log subscription operation binding the contract event 0x17a87a2888b2cae7ae946db8c992fa2329cb5a2152c5578a293bcbd8bb2719a5.
//
// Solidity: event UserCreated(uint256 indexed userId, (uint256,uint256,uint256) user)
func (_User *UserFilterer) WatchUserCreated(opts *bind.WatchOpts, sink chan<- *UserUserCreated, userId []*big.Int) (event.Subscription, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _User.contract.WatchLogs(opts, "UserCreated", userIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserUserCreated)
				if err := _User.contract.UnpackLog(event, "UserCreated", log); err != nil {
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

// ParseUserCreated is a log parse operation binding the contract event 0x17a87a2888b2cae7ae946db8c992fa2329cb5a2152c5578a293bcbd8bb2719a5.
//
// Solidity: event UserCreated(uint256 indexed userId, (uint256,uint256,uint256) user)
func (_User *UserFilterer) ParseUserCreated(log types.Log) (*UserUserCreated, error) {
	event := new(UserUserCreated)
	if err := _User.contract.UnpackLog(event, "UserCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserUserUpdatedIterator is returned from FilterUserUpdated and is used to iterate over the raw logs and unpacked data for UserUpdated events raised by the User contract.
type UserUserUpdatedIterator struct {
	Event *UserUserUpdated // Event containing the contract specifics and raw log

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
func (it *UserUserUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserUserUpdated)
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
		it.Event = new(UserUserUpdated)
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
func (it *UserUserUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserUserUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserUserUpdated represents a UserUpdated event raised by the User contract.
type UserUserUpdated struct {
	UserId *big.Int
	User   UserUserEntity
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUserUpdated is a free log retrieval operation binding the contract event 0x9049d2447d6eb5adf4d038b979f34c6950f566b8bf3244d4247834bd9b1dad82.
//
// Solidity: event UserUpdated(uint256 indexed userId, (uint256,uint256,uint256) user)
func (_User *UserFilterer) FilterUserUpdated(opts *bind.FilterOpts, userId []*big.Int) (*UserUserUpdatedIterator, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _User.contract.FilterLogs(opts, "UserUpdated", userIdRule)
	if err != nil {
		return nil, err
	}
	return &UserUserUpdatedIterator{contract: _User.contract, event: "UserUpdated", logs: logs, sub: sub}, nil
}

// WatchUserUpdated is a free log subscription operation binding the contract event 0x9049d2447d6eb5adf4d038b979f34c6950f566b8bf3244d4247834bd9b1dad82.
//
// Solidity: event UserUpdated(uint256 indexed userId, (uint256,uint256,uint256) user)
func (_User *UserFilterer) WatchUserUpdated(opts *bind.WatchOpts, sink chan<- *UserUserUpdated, userId []*big.Int) (event.Subscription, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	logs, sub, err := _User.contract.WatchLogs(opts, "UserUpdated", userIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserUserUpdated)
				if err := _User.contract.UnpackLog(event, "UserUpdated", log); err != nil {
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

// ParseUserUpdated is a log parse operation binding the contract event 0x9049d2447d6eb5adf4d038b979f34c6950f566b8bf3244d4247834bd9b1dad82.
//
// Solidity: event UserUpdated(uint256 indexed userId, (uint256,uint256,uint256) user)
func (_User *UserFilterer) ParseUserUpdated(log types.Log) (*UserUserUpdated, error) {
	event := new(UserUserUpdated)
	if err := _User.contract.UnpackLog(event, "UserUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
