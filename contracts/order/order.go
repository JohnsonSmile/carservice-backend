// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package order

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

// OrderOrderEntity is an auto generated low-level Go binding around an user-defined struct.
type OrderOrderEntity struct {
	Id             *big.Int
	OrderID        *big.Int
	OrderSn        string
	OrderType      uint8
	StartAt        *big.Int
	EndAt          *big.Int
	UserID         *big.Int
	Fee            *big.Int
	UniteCount     *big.Int
	IsPayed        bool
	BlockHeight    *big.Int
	BlockTimestamp *big.Int
	StartPosition  string
	EndPosition    string
}

// OrderMetaData contains all meta data concerning the Order contract.
var OrderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Error_OrderAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Error_OrderAlreadyPayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Error_OrderNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Error_OrderUserNotEnoughScoreToPay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Error_ShouldInitialized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"orderSn\",\"type\":\"string\"},{\"internalType\":\"enumOrder.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"uniteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPayed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"startPosition\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endPosition\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structOrder.OrderEntity\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"orderSn\",\"type\":\"string\"},{\"internalType\":\"enumOrder.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"uniteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPayed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"startPosition\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endPosition\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structOrder.OrderEntity\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"OrderPayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"orderSn\",\"type\":\"string\"},{\"internalType\":\"enumOrder.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"uniteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPayed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"startPosition\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endPosition\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structOrder.OrderEntity\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"OrderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_orderSn\",\"type\":\"string\"},{\"internalType\":\"enumOrder.OrderType\",\"name\":\"_orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_startAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_userID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_uniteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isPayed\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_startPosition\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_endPosition\",\"type\":\"string\"}],\"name\":\"CreateOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderID\",\"type\":\"uint256\"}],\"name\":\"GetOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"orderSn\",\"type\":\"string\"},{\"internalType\":\"enumOrder.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"uniteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPayed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"startPosition\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endPosition\",\"type\":\"string\"}],\"internalType\":\"structOrder.OrderEntity\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"Initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_userID\",\"type\":\"uint256\"}],\"name\":\"PayOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_orderSn\",\"type\":\"string\"},{\"internalType\":\"enumOrder.OrderType\",\"name\":\"_orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_startAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_userID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_uniteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isPayed\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_startPosition\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_endPosition\",\"type\":\"string\"}],\"name\":\"UpdateOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608080604052346100645760008054336001600160a01b0319821681178355916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3600160025401600255611a80908161006a8239f35b600080fdfe608080604052600436101561001357600080fd5b600090813560e01c9081631c75cfb41461161557508063305db5d01461107b57806336b1453514611020578063715018a614610fc65780637830fed3146108e75780638da5cb5b146108c05780639de44fb0146102c4578063d59a1fce146101575763f2fde38b1461008457600080fd5b34610154576020366003190112610154576004356001600160a01b0381811691829003610150576100b36119f2565b81156100fc57600054826bffffffffffffffffffffffff60a01b821617600055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a380f35b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b8280fd5b80fd5b5034610154576020366003190112610154576040516101758161182a565b818152816020820152606060408201528160608201528160808201528160a08201528160c08201528160e0820152610293600d61010093808585015261012090808286015261014081818701526101609082828801526040610180936060858a015260606101a0809a0152600435815260016020522093604051986101f98a61182a565b85548a52600186015460208b01526102136002870161194c565b60408b015261022c60ff60038801541660608c01611906565b600486015460808b0152600586015460a08b0152600686015460c08b0152600786015460e08b01526008860154908a015260ff600986015416151590890152600a84015490880152600b83015490870152610289600c830161194c565b908601520161194c565b908201528051156102b2576102ae906040519182918261173f565b0390f35b60405163b537178960e01b8152600490fd5b5034610154578060ff6102d636611663565b9a9d909b9198939c9e6102f09a939a9895989796976119f2565b5460a01c16156108ae5760408f8f815260016020522054156102b2578e8e81526001602052604090209c604051809e6103288261182a565b80548252600181015460208301526103426002820161194c565b6040830152600381015461035c9060ff1660608401611906565b60048101546080830152600581015460a0830152600681015460c0830152600781015460e08301526008810154610100830152600981015460ff161515610120830152600a810154610140830152600b8101546101608301526103c1600c820161194c565b610180830152600d016103d39061194c565b906101a0015236906103e4926118c0565b60408d01526103f69060608d01611906565b60808b015260a08a015260c089015260e0880152610100870152369061041b926118c0565b610180850152369061042c926118c0565b6101a08301521515610120820152818352600160205260408320815181556020820151600182015560408201518051906001600160401b0382116107005781906104796002850154611912565b601f811161085e575b50602090601f83116001146107ef5787926107e4575b50508160011b916000199060031b1c19161760028201555b6003810160608301519060038210156107d05760ff199160ff8383541691161790556080830151600483015560a0830151600583015560c0830151600683015560e083015160078301556101008301516008830155600982019060ff6101208501511515918354169116179055610140820151600a820155610160820151600b8201556101808201518051906001600160401b03821161070057610557600c840154611912565b601f811161078c575b50602090601f831160011461071f57600d939291879183610714575b50508160011b916000199060031b1c191617600c8201555b016101a08201518051906001600160401b038211610700576105b68354611912565b601f81116106bb575b50602090601f831160011461063157918061062094927f363832d3b3f06348c3120f3513a621cf078922944f359dafeb37f4d4ebbb802196948992610626575b50508160011b916000199060031b1c19161790555b6040519182918261173f565b0390a280f35b0151905038806105ff565b838752602087209190601f198416885b8181106106a357509260019285927f363832d3b3f06348c3120f3513a621cf078922944f359dafeb37f4d4ebbb8021989661062098961061068a575b505050811b019055610614565b015160001960f88460031b161c1916905538808061067d565b92936020600181928786015181550195019301610641565b83875260208720601f840160051c810191602085106106f6575b601f0160051c01905b8181106106eb57506105bf565b8781556001016106de565b90915081906106d5565b634e487b7160e01b86526041600452602486fd5b01519050388061057c565b600c84018752602087209190601f198416885b8181106107745750916001939185600d9796941061075b575b505050811b01600c820155610594565b015160001960f88460031b161c1916905538808061074b565b92936020600181928786015181550195019301610732565b600c8401875260208720601f840160051c8101602085106107c9575b601f830160051c820181106107be575050610560565b8881556001016107a8565b50806107a8565b634e487b7160e01b86526021600452602486fd5b015190503880610498565b9250600284018752602087209087935b601f1984168510610843576001945083601f1981161061082a575b505050811b0160028201556104b0565b015160001960f88460031b161c1916905538808061081a565b818101518355602094850194600190930192909101906107ff565b90915060028401875260208720601f840160051c8101602085106108a7575b90849392915b601f830160051c82018110610899575050610482565b898155859450600101610883565b508061087d565b60405163b340155d60e01b8152600490fd5b5034610154578060031936011261015457546040516001600160a01b039091168152602090f35b5034610154576040366003190112610154576109016119f2565b60ff815460a01c16156108ae576004358152600160205260408120906109d3600d6040519361092f8561182a565b80548552600181015460208601526109496002820161194c565b604086015261096260ff60038301541660608701611906565b60048101546080860152600581015460a0860152600681015460c0860152600781015460e0860152600881015461010086015260ff6009820154161515610120860152600a810154610140860152600b8101546101608601526109c7600c820161194c565b6101808601520161194c565b6101a08301528151156102b257610120820151610fb45760e0820151610100830151908082810204821481151715610fa05790606491020460608301516003811015610f8c5715610f81575b6003546040516303878ea960e11b8152602480356004830152909291606091849182906001600160a01b03165afa918215610f76578392610f46575b5080604083015110610f345760016101208501526004358352600160205260408320845181556020850151600182015560408501518051906001600160401b038211610700578190610ab06002850154611912565b601f8111610ee4575b50602090601f8311600114610e75578792610e6a575b50508160011b916000199060031b1c19161760028201555b60608501516003811015610e56576003820160ff199160ff8383541691161790556080860151600483015560a0860151600583015560c0860151600683015560e086015160078301556101008601516008830155600982019060ff6101208801511515918354169116179055610140850151600a820155610160850151600b8201556101808501518051906001600160401b038211610700578190610b8f600c850154611912565b601f8111610e06575b50602090601f8311600114610d9c578792610d91575b50508160011b916000199060031b1c191617600c8201555b6101a08501518051906001600160401b03821161070057610bea600d840154611912565b601f8111610d4d575b50602090601f8311600114610ce057600d9291879183610cd5575b50508160011b916000199060031b1c1916179101555b60018060a01b03600354169151823b15610cd157906044849283604051958694859363148c32b760e01b8552600485015260248401525af18015610cc657610c9b575b50907fe3fbc07850e69ef4022dfb43bddb3544c98b02b6f914f34d8e7895609b38409a60405180610620600435948261173f565b6001600160401b038111610cb25760405238610c67565b634e487b7160e01b82526041600452602482fd5b6040513d84823e3d90fd5b8380fd5b015190503880610c0e565b90600d840187526020872091875b601f1985168110610d355750918391600193600d95601f19811610610d1c575b505050811b01910155610c24565b015160001960f88460031b161c19169055388080610d0e565b91926020600181928685015181550194019201610cee565b600d8401875260208720601f840160051c810160208510610d8a575b601f830160051c82018110610d7f575050610bf3565b888155600101610d69565b5080610d69565b015190503880610bae565b600c85018852602088209250601f198416885b818110610dee5750908460019594939210610dd5575b505050811b01600c820155610bc6565b015160001960f88460031b161c19169055388080610dc5565b92936020600181928786015181550195019301610daf565b909150600c8401875260208720601f840160051c810160208510610e4f575b90849392915b601f830160051c82018110610e41575050610b98565b898155859450600101610e2b565b5080610e25565b634e487b7160e01b85526021600452602485fd5b015190503880610acf565b9250600284018752602087209087935b601f1984168510610ec9576001945083601f19811610610eb0575b505050811b016002820155610ae7565b015160001960f88460031b161c19169055388080610ea0565b81810151835560209485019460019093019290910190610e85565b90915060028401875260208720601f840160051c810160208510610f2d575b90849392915b601f830160051c82018110610f1f575050610ab9565b898155859450600101610f09565b5080610f03565b60405163d60de76f60e01b8152600490fd5b610f6891925060603d8111610f6f575b610f60818361185c565b81019061187d565b9038610a5b565b503d610f56565b6040513d85823e3d90fd5b5060e0820151610a1f565b634e487b7160e01b83526021600452602483fd5b634e487b7160e01b83526011600452602483fd5b60405163682a017760e11b8152600490fd5b5034610154578060031936011261015457610fdf6119f2565b80546001600160a01b03198116825581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b5034610154576020366003190112610154576004356001600160a01b038116908190036110775761104f6119f2565b815460ff60a01b1916600160a01b178255600380546001600160a01b03191691909117905580f35b5080fd5b5034610154578060ff61108d36611663565b9b9199929a9d909c9e6110a79994999895989796976119f2565b5460a01c16156108ae5760408f8f815260016020522054611603576003546040516303878ea960e11b8152600481018790529d908e9081906001600160a01b03165a92602491606094fa9d8e156115f5578f9e6115d7575b506002549d6040519e8f906111138261182a565b8152602001523690611124926118c0565b60408d01526111369060608d01611906565b60808b015260a08a015260c089015260e0880152610100870152151561012086015243610140860152426101608601523690611171926118c0565b6101808401523690611182926118c0565b6101a0820152818352600160205260408320815181556020820151600182015560408201518051906001600160401b0382116107005781906111c76002850154611912565b601f8111611587575b50602090601f831160011461151857879261150d575b50508160011b916000199060031b1c19161760028201555b6003810160608301519060038210156107d05760ff199160ff8383541691161790556080830151600483015560a0830151600583015560c0830151600683015560e083015160078301556101008301516008830155600982019060ff6101208501511515918354169116179055610140820151600a820155610160820151600b8201556101808201518051906001600160401b038211610700576112a5600c840154611912565b601f81116114c9575b50602090601f831160011461145c57600d939291879183611451575b50508160011b916000199060031b1c191617600c8201555b016101a08201518051906001600160401b038211610700576113048354611912565b601f811161140c575b50602090601f831160011461138257918061062094927fe0066053412f6b53f46598c1b74b2ca24cc9cfcbf4364c620ec75a496c2d226d96948992611377575b50508160011b916000199060031b1c19161790555b6001600254016002556040519182918261173f565b01519050388061134d565b838752602087209190601f198416885b8181106113f457509260019285927fe0066053412f6b53f46598c1b74b2ca24cc9cfcbf4364c620ec75a496c2d226d98966106209896106113db575b505050811b019055611362565b015160001960f88460031b161c191690553880806113ce565b92936020600181928786015181550195019301611392565b83875260208720601f840160051c81019160208510611447575b601f0160051c01905b81811061143c575061130d565b87815560010161142f565b9091508190611426565b0151905038806112ca565b600c84018752602087209190601f198416885b8181106114b15750916001939185600d97969410611498575b505050811b01600c8201556112e2565b015160001960f88460031b161c19169055388080611488565b9293602060018192878601518155019501930161146f565b600c8401875260208720601f840160051c810160208510611506575b601f830160051c820181106114fb5750506112ae565b8881556001016114e5565b50806114e5565b0151905038806111e6565b9250600284018752602087209087935b601f198416851061156c576001945083601f19811610611553575b505050811b0160028201556111fe565b015160001960f88460031b161c19169055388080611543565b81810151835560209485019460019093019290910190611528565b90915060028401875260208720601f840160051c8101602085106115d0575b90849392915b601f830160051c820181106115c25750506111d0565b8981558594506001016115ac565b50806115a6565b6115ee9060603d8111610f6f57610f60818361185c565b50386110ff565b508f604051903d90823e3d90fd5b604051630f91c39160e01b8152600490fd5b9050346110775781600319360112611077576020906002548152f35b9181601f8401121561165e578235916001600160401b03831161165e576020838186019501011161165e57565b600080fd5b9061016060031983011261165e57600435916001600160401b039160243583811161165e578261169591600401611631565b93909392604435600381101561165e5792606435926084359260a4359260c4359260e4359261010435801515810361165e57926101243583811161165e57826116e091600401611631565b939093926101443591821161165e576116fb91600401611631565b9091565b919082519283825260005b84811061172b575050826000602080949584010152601f8019910116010190565b60208183018101518483018201520161170a565b906020825280516020830152602081015160408301526040810151916117736101c0938460608401526101e08301906116ff565b90606083015193600385101561181457611811946080830152608084015160a083015260a084015160c083015260c084015160e083015260e0840151610100908184015284015161012090818401528401511515610140908184015284015161016090818401528401516101809081840152840151611801601f19946101a0928686830301848701526116ff565b94015192828503019101526116ff565b90565b634e487b7160e01b600052602160045260246000fd5b6101c081019081106001600160401b0382111761184657604052565b634e487b7160e01b600052604160045260246000fd5b90601f801991011681019081106001600160401b0382111761184657604052565b9081606091031261165e576040519060608201908282106001600160401b0383111761184657604091825280518352602081015160208401520151604082015290565b9291926001600160401b03821161184657604051916118e9601f8201601f19166020018461185c565b82948184528183011161165e578281602093846000960137010152565b60038210156118145752565b90600182811c92168015611942575b602083101461192c57565b634e487b7160e01b600052602260045260246000fd5b91607f1691611921565b906040519182600082549261196084611912565b9081845260019485811690816000146119cf575060011461198c575b505061198a9250038361185c565b565b9093915060005260209081600020936000915b8183106119b757505061198a9350820101388061197c565b8554888401850152948501948794509183019161199f565b91505061198a94506020925060ff191682840152151560051b820101388061197c565b6000546001600160a01b03163303611a0657565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fdfea26469706673582212202763778258d16cc7d68548623c58e060c2254cad91e60f856bd97243df8c12d864736f6c63430008130033",
}

// OrderABI is the input ABI used to generate the binding from.
// Deprecated: Use OrderMetaData.ABI instead.
var OrderABI = OrderMetaData.ABI

// OrderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OrderMetaData.Bin instead.
var OrderBin = OrderMetaData.Bin

// DeployOrder deploys a new Ethereum contract, binding an instance of Order to it.
func DeployOrder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Order, error) {
	parsed, err := OrderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OrderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Order{OrderCaller: OrderCaller{contract: contract}, OrderTransactor: OrderTransactor{contract: contract}, OrderFilterer: OrderFilterer{contract: contract}}, nil
}

// Order is an auto generated Go binding around an Ethereum contract.
type Order struct {
	OrderCaller     // Read-only binding to the contract
	OrderTransactor // Write-only binding to the contract
	OrderFilterer   // Log filterer for contract events
}

// OrderCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderSession struct {
	Contract     *Order            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderCallerSession struct {
	Contract *OrderCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OrderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderTransactorSession struct {
	Contract     *OrderTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderRaw struct {
	Contract *Order // Generic contract binding to access the raw methods on
}

// OrderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderCallerRaw struct {
	Contract *OrderCaller // Generic read-only contract binding to access the raw methods on
}

// OrderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderTransactorRaw struct {
	Contract *OrderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrder creates a new instance of Order, bound to a specific deployed contract.
func NewOrder(address common.Address, backend bind.ContractBackend) (*Order, error) {
	contract, err := bindOrder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Order{OrderCaller: OrderCaller{contract: contract}, OrderTransactor: OrderTransactor{contract: contract}, OrderFilterer: OrderFilterer{contract: contract}}, nil
}

// NewOrderCaller creates a new read-only instance of Order, bound to a specific deployed contract.
func NewOrderCaller(address common.Address, caller bind.ContractCaller) (*OrderCaller, error) {
	contract, err := bindOrder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderCaller{contract: contract}, nil
}

// NewOrderTransactor creates a new write-only instance of Order, bound to a specific deployed contract.
func NewOrderTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderTransactor, error) {
	contract, err := bindOrder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderTransactor{contract: contract}, nil
}

// NewOrderFilterer creates a new log filterer instance of Order, bound to a specific deployed contract.
func NewOrderFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderFilterer, error) {
	contract, err := bindOrder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderFilterer{contract: contract}, nil
}

// bindOrder binds a generic wrapper to an already deployed contract.
func bindOrder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Order *OrderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Order.Contract.OrderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Order *OrderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Order.Contract.OrderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Order *OrderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Order.Contract.OrderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Order *OrderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Order.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Order *OrderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Order.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Order *OrderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Order.Contract.contract.Transact(opts, method, params...)
}

// GetOrder is a free data retrieval call binding the contract method 0xd59a1fce.
//
// Solidity: function GetOrder(uint256 _orderID) view returns((uint256,uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,string))
func (_Order *OrderCaller) GetOrder(opts *bind.CallOpts, _orderID *big.Int) (OrderOrderEntity, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "GetOrder", _orderID)

	if err != nil {
		return *new(OrderOrderEntity), err
	}

	out0 := *abi.ConvertType(out[0], new(OrderOrderEntity)).(*OrderOrderEntity)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0xd59a1fce.
//
// Solidity: function GetOrder(uint256 _orderID) view returns((uint256,uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,string))
func (_Order *OrderSession) GetOrder(_orderID *big.Int) (OrderOrderEntity, error) {
	return _Order.Contract.GetOrder(&_Order.CallOpts, _orderID)
}

// GetOrder is a free data retrieval call binding the contract method 0xd59a1fce.
//
// Solidity: function GetOrder(uint256 _orderID) view returns((uint256,uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,string))
func (_Order *OrderCallerSession) GetOrder(_orderID *big.Int) (OrderOrderEntity, error) {
	return _Order.Contract.GetOrder(&_Order.CallOpts, _orderID)
}

// CurrentID is a free data retrieval call binding the contract method 0x1c75cfb4.
//
// Solidity: function currentID() view returns(uint256 _value)
func (_Order *OrderCaller) CurrentID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "currentID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentID is a free data retrieval call binding the contract method 0x1c75cfb4.
//
// Solidity: function currentID() view returns(uint256 _value)
func (_Order *OrderSession) CurrentID() (*big.Int, error) {
	return _Order.Contract.CurrentID(&_Order.CallOpts)
}

// CurrentID is a free data retrieval call binding the contract method 0x1c75cfb4.
//
// Solidity: function currentID() view returns(uint256 _value)
func (_Order *OrderCallerSession) CurrentID() (*big.Int, error) {
	return _Order.Contract.CurrentID(&_Order.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Order *OrderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Order.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Order *OrderSession) Owner() (common.Address, error) {
	return _Order.Contract.Owner(&_Order.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Order *OrderCallerSession) Owner() (common.Address, error) {
	return _Order.Contract.Owner(&_Order.CallOpts)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x305db5d0.
//
// Solidity: function CreateOrder(uint256 _orderID, string _orderSn, uint8 _orderType, uint256 _startAt, uint256 _endAt, uint256 _userID, uint256 _fee, uint256 _uniteCount, bool _isPayed, string _startPosition, string _endPosition) returns()
func (_Order *OrderTransactor) CreateOrder(opts *bind.TransactOpts, _orderID *big.Int, _orderSn string, _orderType uint8, _startAt *big.Int, _endAt *big.Int, _userID *big.Int, _fee *big.Int, _uniteCount *big.Int, _isPayed bool, _startPosition string, _endPosition string) (*types.Transaction, error) {
	return _Order.contract.Transact(opts, "CreateOrder", _orderID, _orderSn, _orderType, _startAt, _endAt, _userID, _fee, _uniteCount, _isPayed, _startPosition, _endPosition)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x305db5d0.
//
// Solidity: function CreateOrder(uint256 _orderID, string _orderSn, uint8 _orderType, uint256 _startAt, uint256 _endAt, uint256 _userID, uint256 _fee, uint256 _uniteCount, bool _isPayed, string _startPosition, string _endPosition) returns()
func (_Order *OrderSession) CreateOrder(_orderID *big.Int, _orderSn string, _orderType uint8, _startAt *big.Int, _endAt *big.Int, _userID *big.Int, _fee *big.Int, _uniteCount *big.Int, _isPayed bool, _startPosition string, _endPosition string) (*types.Transaction, error) {
	return _Order.Contract.CreateOrder(&_Order.TransactOpts, _orderID, _orderSn, _orderType, _startAt, _endAt, _userID, _fee, _uniteCount, _isPayed, _startPosition, _endPosition)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x305db5d0.
//
// Solidity: function CreateOrder(uint256 _orderID, string _orderSn, uint8 _orderType, uint256 _startAt, uint256 _endAt, uint256 _userID, uint256 _fee, uint256 _uniteCount, bool _isPayed, string _startPosition, string _endPosition) returns()
func (_Order *OrderTransactorSession) CreateOrder(_orderID *big.Int, _orderSn string, _orderType uint8, _startAt *big.Int, _endAt *big.Int, _userID *big.Int, _fee *big.Int, _uniteCount *big.Int, _isPayed bool, _startPosition string, _endPosition string) (*types.Transaction, error) {
	return _Order.Contract.CreateOrder(&_Order.TransactOpts, _orderID, _orderSn, _orderType, _startAt, _endAt, _userID, _fee, _uniteCount, _isPayed, _startPosition, _endPosition)
}

// Initialize is a paid mutator transaction binding the contract method 0x36b14535.
//
// Solidity: function Initialize(address _user) returns()
func (_Order *OrderTransactor) Initialize(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Order.contract.Transact(opts, "Initialize", _user)
}

// Initialize is a paid mutator transaction binding the contract method 0x36b14535.
//
// Solidity: function Initialize(address _user) returns()
func (_Order *OrderSession) Initialize(_user common.Address) (*types.Transaction, error) {
	return _Order.Contract.Initialize(&_Order.TransactOpts, _user)
}

// Initialize is a paid mutator transaction binding the contract method 0x36b14535.
//
// Solidity: function Initialize(address _user) returns()
func (_Order *OrderTransactorSession) Initialize(_user common.Address) (*types.Transaction, error) {
	return _Order.Contract.Initialize(&_Order.TransactOpts, _user)
}

// PayOrder is a paid mutator transaction binding the contract method 0x7830fed3.
//
// Solidity: function PayOrder(uint256 _orderID, uint256 _userID) returns()
func (_Order *OrderTransactor) PayOrder(opts *bind.TransactOpts, _orderID *big.Int, _userID *big.Int) (*types.Transaction, error) {
	return _Order.contract.Transact(opts, "PayOrder", _orderID, _userID)
}

// PayOrder is a paid mutator transaction binding the contract method 0x7830fed3.
//
// Solidity: function PayOrder(uint256 _orderID, uint256 _userID) returns()
func (_Order *OrderSession) PayOrder(_orderID *big.Int, _userID *big.Int) (*types.Transaction, error) {
	return _Order.Contract.PayOrder(&_Order.TransactOpts, _orderID, _userID)
}

// PayOrder is a paid mutator transaction binding the contract method 0x7830fed3.
//
// Solidity: function PayOrder(uint256 _orderID, uint256 _userID) returns()
func (_Order *OrderTransactorSession) PayOrder(_orderID *big.Int, _userID *big.Int) (*types.Transaction, error) {
	return _Order.Contract.PayOrder(&_Order.TransactOpts, _orderID, _userID)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0x9de44fb0.
//
// Solidity: function UpdateOrder(uint256 _orderID, string _orderSn, uint8 _orderType, uint256 _startAt, uint256 _endAt, uint256 _userID, uint256 _fee, uint256 _uniteCount, bool _isPayed, string _startPosition, string _endPosition) returns()
func (_Order *OrderTransactor) UpdateOrder(opts *bind.TransactOpts, _orderID *big.Int, _orderSn string, _orderType uint8, _startAt *big.Int, _endAt *big.Int, _userID *big.Int, _fee *big.Int, _uniteCount *big.Int, _isPayed bool, _startPosition string, _endPosition string) (*types.Transaction, error) {
	return _Order.contract.Transact(opts, "UpdateOrder", _orderID, _orderSn, _orderType, _startAt, _endAt, _userID, _fee, _uniteCount, _isPayed, _startPosition, _endPosition)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0x9de44fb0.
//
// Solidity: function UpdateOrder(uint256 _orderID, string _orderSn, uint8 _orderType, uint256 _startAt, uint256 _endAt, uint256 _userID, uint256 _fee, uint256 _uniteCount, bool _isPayed, string _startPosition, string _endPosition) returns()
func (_Order *OrderSession) UpdateOrder(_orderID *big.Int, _orderSn string, _orderType uint8, _startAt *big.Int, _endAt *big.Int, _userID *big.Int, _fee *big.Int, _uniteCount *big.Int, _isPayed bool, _startPosition string, _endPosition string) (*types.Transaction, error) {
	return _Order.Contract.UpdateOrder(&_Order.TransactOpts, _orderID, _orderSn, _orderType, _startAt, _endAt, _userID, _fee, _uniteCount, _isPayed, _startPosition, _endPosition)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0x9de44fb0.
//
// Solidity: function UpdateOrder(uint256 _orderID, string _orderSn, uint8 _orderType, uint256 _startAt, uint256 _endAt, uint256 _userID, uint256 _fee, uint256 _uniteCount, bool _isPayed, string _startPosition, string _endPosition) returns()
func (_Order *OrderTransactorSession) UpdateOrder(_orderID *big.Int, _orderSn string, _orderType uint8, _startAt *big.Int, _endAt *big.Int, _userID *big.Int, _fee *big.Int, _uniteCount *big.Int, _isPayed bool, _startPosition string, _endPosition string) (*types.Transaction, error) {
	return _Order.Contract.UpdateOrder(&_Order.TransactOpts, _orderID, _orderSn, _orderType, _startAt, _endAt, _userID, _fee, _uniteCount, _isPayed, _startPosition, _endPosition)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Order *OrderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Order.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Order *OrderSession) RenounceOwnership() (*types.Transaction, error) {
	return _Order.Contract.RenounceOwnership(&_Order.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Order *OrderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Order.Contract.RenounceOwnership(&_Order.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Order *OrderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Order.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Order *OrderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Order.Contract.TransferOwnership(&_Order.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Order *OrderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Order.Contract.TransferOwnership(&_Order.TransactOpts, newOwner)
}

// OrderOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the Order contract.
type OrderOrderCreatedIterator struct {
	Event *OrderOrderCreated // Event containing the contract specifics and raw log

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
func (it *OrderOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderOrderCreated)
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
		it.Event = new(OrderOrderCreated)
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
func (it *OrderOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderOrderCreated represents a OrderCreated event raised by the Order contract.
type OrderOrderCreated struct {
	OrderID *big.Int
	Order   OrderOrderEntity
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0xe0066053412f6b53f46598c1b74b2ca24cc9cfcbf4364c620ec75a496c2d226d.
//
// Solidity: event OrderCreated(uint256 indexed orderID, (uint256,uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,string) order)
func (_Order *OrderFilterer) FilterOrderCreated(opts *bind.FilterOpts, orderID []*big.Int) (*OrderOrderCreatedIterator, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}

	logs, sub, err := _Order.contract.FilterLogs(opts, "OrderCreated", orderIDRule)
	if err != nil {
		return nil, err
	}
	return &OrderOrderCreatedIterator{contract: _Order.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0xe0066053412f6b53f46598c1b74b2ca24cc9cfcbf4364c620ec75a496c2d226d.
//
// Solidity: event OrderCreated(uint256 indexed orderID, (uint256,uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,string) order)
func (_Order *OrderFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *OrderOrderCreated, orderID []*big.Int) (event.Subscription, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}

	logs, sub, err := _Order.contract.WatchLogs(opts, "OrderCreated", orderIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderOrderCreated)
				if err := _Order.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0xe0066053412f6b53f46598c1b74b2ca24cc9cfcbf4364c620ec75a496c2d226d.
//
// Solidity: event OrderCreated(uint256 indexed orderID, (uint256,uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,string) order)
func (_Order *OrderFilterer) ParseOrderCreated(log types.Log) (*OrderOrderCreated, error) {
	event := new(OrderOrderCreated)
	if err := _Order.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderOrderPayedIterator is returned from FilterOrderPayed and is used to iterate over the raw logs and unpacked data for OrderPayed events raised by the Order contract.
type OrderOrderPayedIterator struct {
	Event *OrderOrderPayed // Event containing the contract specifics and raw log

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
func (it *OrderOrderPayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderOrderPayed)
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
		it.Event = new(OrderOrderPayed)
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
func (it *OrderOrderPayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderOrderPayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderOrderPayed represents a OrderPayed event raised by the Order contract.
type OrderOrderPayed struct {
	OrderID *big.Int
	Order   OrderOrderEntity
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderPayed is a free log retrieval operation binding the contract event 0xe3fbc07850e69ef4022dfb43bddb3544c98b02b6f914f34d8e7895609b38409a.
//
// Solidity: event OrderPayed(uint256 indexed orderID, (uint256,uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,string) order)
func (_Order *OrderFilterer) FilterOrderPayed(opts *bind.FilterOpts, orderID []*big.Int) (*OrderOrderPayedIterator, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}

	logs, sub, err := _Order.contract.FilterLogs(opts, "OrderPayed", orderIDRule)
	if err != nil {
		return nil, err
	}
	return &OrderOrderPayedIterator{contract: _Order.contract, event: "OrderPayed", logs: logs, sub: sub}, nil
}

// WatchOrderPayed is a free log subscription operation binding the contract event 0xe3fbc07850e69ef4022dfb43bddb3544c98b02b6f914f34d8e7895609b38409a.
//
// Solidity: event OrderPayed(uint256 indexed orderID, (uint256,uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,string) order)
func (_Order *OrderFilterer) WatchOrderPayed(opts *bind.WatchOpts, sink chan<- *OrderOrderPayed, orderID []*big.Int) (event.Subscription, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}

	logs, sub, err := _Order.contract.WatchLogs(opts, "OrderPayed", orderIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderOrderPayed)
				if err := _Order.contract.UnpackLog(event, "OrderPayed", log); err != nil {
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

// ParseOrderPayed is a log parse operation binding the contract event 0xe3fbc07850e69ef4022dfb43bddb3544c98b02b6f914f34d8e7895609b38409a.
//
// Solidity: event OrderPayed(uint256 indexed orderID, (uint256,uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,string) order)
func (_Order *OrderFilterer) ParseOrderPayed(log types.Log) (*OrderOrderPayed, error) {
	event := new(OrderOrderPayed)
	if err := _Order.contract.UnpackLog(event, "OrderPayed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderOrderUpdatedIterator is returned from FilterOrderUpdated and is used to iterate over the raw logs and unpacked data for OrderUpdated events raised by the Order contract.
type OrderOrderUpdatedIterator struct {
	Event *OrderOrderUpdated // Event containing the contract specifics and raw log

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
func (it *OrderOrderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderOrderUpdated)
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
		it.Event = new(OrderOrderUpdated)
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
func (it *OrderOrderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderOrderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderOrderUpdated represents a OrderUpdated event raised by the Order contract.
type OrderOrderUpdated struct {
	OrderID *big.Int
	Order   OrderOrderEntity
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderUpdated is a free log retrieval operation binding the contract event 0x363832d3b3f06348c3120f3513a621cf078922944f359dafeb37f4d4ebbb8021.
//
// Solidity: event OrderUpdated(uint256 indexed orderID, (uint256,uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,string) order)
func (_Order *OrderFilterer) FilterOrderUpdated(opts *bind.FilterOpts, orderID []*big.Int) (*OrderOrderUpdatedIterator, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}

	logs, sub, err := _Order.contract.FilterLogs(opts, "OrderUpdated", orderIDRule)
	if err != nil {
		return nil, err
	}
	return &OrderOrderUpdatedIterator{contract: _Order.contract, event: "OrderUpdated", logs: logs, sub: sub}, nil
}

// WatchOrderUpdated is a free log subscription operation binding the contract event 0x363832d3b3f06348c3120f3513a621cf078922944f359dafeb37f4d4ebbb8021.
//
// Solidity: event OrderUpdated(uint256 indexed orderID, (uint256,uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,string) order)
func (_Order *OrderFilterer) WatchOrderUpdated(opts *bind.WatchOpts, sink chan<- *OrderOrderUpdated, orderID []*big.Int) (event.Subscription, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}

	logs, sub, err := _Order.contract.WatchLogs(opts, "OrderUpdated", orderIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderOrderUpdated)
				if err := _Order.contract.UnpackLog(event, "OrderUpdated", log); err != nil {
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

// ParseOrderUpdated is a log parse operation binding the contract event 0x363832d3b3f06348c3120f3513a621cf078922944f359dafeb37f4d4ebbb8021.
//
// Solidity: event OrderUpdated(uint256 indexed orderID, (uint256,uint256,string,uint8,uint256,uint256,uint256,uint256,uint256,bool,uint256,uint256,string,string) order)
func (_Order *OrderFilterer) ParseOrderUpdated(log types.Log) (*OrderOrderUpdated, error) {
	event := new(OrderOrderUpdated)
	if err := _Order.contract.UnpackLog(event, "OrderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Order contract.
type OrderOwnershipTransferredIterator struct {
	Event *OrderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OrderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderOwnershipTransferred)
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
		it.Event = new(OrderOwnershipTransferred)
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
func (it *OrderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderOwnershipTransferred represents a OwnershipTransferred event raised by the Order contract.
type OrderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Order *OrderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OrderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Order.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OrderOwnershipTransferredIterator{contract: _Order.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Order *OrderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OrderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Order.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderOwnershipTransferred)
				if err := _Order.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Order *OrderFilterer) ParseOwnershipTransferred(log types.Log) (*OrderOwnershipTransferred, error) {
	event := new(OrderOwnershipTransferred)
	if err := _Order.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
