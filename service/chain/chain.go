package chain

import (
	"carservice/contracts/order"
	"carservice/contracts/user"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"github.com/ethereum/go-ethereum/accounts/abi"
	cm "github.com/ethereum/go-ethereum/common"
)

type ChainClient struct {
	client        *sdk.ChainClient
	userContract  *common.Contract
	userABI       *abi.ABI
	orderContract *common.Contract
	orderABI      *abi.ABI
}

type Option = func(*ChainClient)

func NewChainClient(opts ...Option) (*ChainClient, error) {
	cli, err := createClientWithConfig()
	if err != nil {
		return nil, err
	}
	client := &ChainClient{
		client: cli,
	}

	for _, opt := range opts {
		opt(client)
	}
	return client, nil
}

// options

func createClientWithConfig() (*sdk.ChainClient, error) {
	chainClient, err := sdk.NewChainClient(
		sdk.WithConfPath("./sdk_config.yml"),
	)

	if err != nil {
		return nil, err
	}

	//启用证书压缩（开启证书压缩可以减小交易包大小，提升处理性能）
	err = chainClient.EnableCertHash()
	if err != nil {
		return nil, err
	}

	return chainClient, nil
}

func (cc *ChainClient) InitContracts() {
	var (
		userContractName     = "user1"
		orderContractName    = "order1"
		userContractVersion  = "1.0.0"
		orderContractVersion = "1.0.0"
		userContractAddress  = ""
		orderContractAddress = ""
	)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	userContract, err := cc.client.GetContractInfo(userContractName)
	if err != nil {
		if strings.Contains(err.Error(), "contract not exist") || userContract == nil || userContract.Name == "" {
			log.Printf("合约[%s]不存在\n", userContractName)
			log.Println("====================== 创建合约 ======================")
		}
		// read bytecode
		byteCode, err := os.ReadFile(path.Join(dir, "contracts/user/User.bin"))
		if err != nil {
			log.Fatalln(err)
		}
		createPayload, err := cc.client.CreateContractCreatePayload(userContractName, userContractVersion, string(byteCode), common.RuntimeType_EVM, []*common.KeyValuePair{})
		if err != nil {
			log.Fatalln(err)
		}

		usernames := []string{UserNameOrg1Admin1, UserNameOrg2Admin1, UserNameOrg3Admin1, UserNameOrg4Admin1}

		endorsers, err := GetEndorsersWithAuthType(cc.client.GetHashType(),
			cc.client.GetAuthType(), createPayload, usernames...)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := cc.client.SendContractManageRequest(createPayload, endorsers, 30000, true)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%+v\n", resp)

		userContract, err := cc.client.GetContractInfo(userContractName)
		if err != nil {
			log.Fatalln(err)
		}
		userContractAddress = userContract.Address
		cc.userContract = userContract
	} else {
		userContractAddress = userContract.Address
		cc.userContract = userContract
	}

	log.Printf("user contract address is %s\n", userContractAddress)

	orderContract, err := cc.client.GetContractInfo(orderContractName)
	if err != nil {
		if strings.Contains(err.Error(), "contract not exist") || orderContract == nil || orderContract.Name == "" {
			log.Printf("合约[%s]不存在\n", orderContractName)
			log.Println("====================== 创建合约 ======================")
		}
		// read bytecode
		byteCode, err := os.ReadFile(path.Join(dir, "contracts/order/Order.bin"))
		if err != nil {
			log.Fatalln(err)
		}
		if err != nil {
			log.Fatalln(err)
		}
		createPayload, err := cc.client.CreateContractCreatePayload(orderContractName, orderContractVersion, string(byteCode), common.RuntimeType_EVM, []*common.KeyValuePair{})
		if err != nil {
			log.Fatalln(err)
		}

		usernames := []string{UserNameOrg1Admin1, UserNameOrg2Admin1, UserNameOrg3Admin1, UserNameOrg4Admin1}

		endorsers, err := GetEndorsersWithAuthType(cc.client.GetHashType(),
			cc.client.GetAuthType(), createPayload, usernames...)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := cc.client.SendContractManageRequest(createPayload, endorsers, 30000, true)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%+v\n", resp)

		orderContract, err := cc.client.GetContractInfo(orderContractName)
		if err != nil {
			log.Fatalln(err)
		}
		cc.orderContract = orderContract
		orderContractAddress = orderContract.Address
	} else {
		cc.orderContract = orderContract
		orderContractAddress = orderContract.Address
	}

	log.Printf("order contract address is %s\n", orderContractAddress)

	// init abi
	// 使用 go-ethereum的abi包
	userABI, err := user.UserMetaData.GetAbi()
	if err != nil {
		log.Fatalln(err)
	}
	cc.userABI = userABI

	orderABI, err := order.OrderMetaData.GetAbi()
	if err != nil {
		log.Fatalln(err)
	}
	cc.orderABI = orderABI

	// order initialize
	initializeMethodName := "Initialize"
	dataByte, err := orderABI.Pack(initializeMethodName, cm.HexToAddress(userContractAddress))
	if err != nil {
		log.Fatalf("initialize err: %+v\n", err)
	}
	dataString := hex.EncodeToString(dataByte)
	kvs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(dataString),
		},
	}
	result, err := invokeContractWithResult(cc.client, orderContractName, initializeMethodName, "", kvs, true)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("result: %+v\n", result)

	// grant manager for order
	isManager, err := cc.IsManager(orderContractAddress)
	if err != nil {
		log.Fatalln(err)
	}
	if !isManager {
		err = cc.GrantManager(orderContractAddress)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// FIXME: test create user

	// err = cc.CreateUser(1, 18539265600, 100000)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// FIXME: test get user
	// userInfo, err := cc.GetUserInfo(1)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Printf("user info: %+v\n", userInfo)

	// FIXME: test charge score
	// err = cc.ChargeScore(1, 1000)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// FIXME: test pay score
	// err = cc.PayScore(1, 1000)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

}

func (cc *ChainClient) Shutdown() {
	cc.client.Stop()
}

func invokeContractWithResult(client *sdk.ChainClient, contractName, method, txId string,
	kvs []*common.KeyValuePair, withSyncResult bool) ([]byte, error) {

	resp, err := client.InvokeContract(contractName, method, txId, kvs, -1, withSyncResult)
	if err != nil {
		return nil, err
	}

	if resp.Code != common.TxStatusCode_SUCCESS {
		return nil, fmt.Errorf("invoke contract failed, [code:%d]/[msg:%s]", resp.Code, resp.Message)
	}

	return resp.ContractResult.Result, nil
}
