package chain

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"chainmaker.org/chainmaker/common/v2/evmutils"
	"chainmaker.org/chainmaker/common/v2/evmutils/abi"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
)

type ChainClient struct {
	client *sdk.ChainClient
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

func (cc *ChainClient) TokenContractBalanceOf(address string) error {
	// kvs := common.
	contractName := "token"

	version := "1.0.0"

	contract, err := cc.client.GetContractInfo(contractName)
	if err != nil {
		if strings.Contains(err.Error(), "contract not exist") || contract == nil || contract.Name == "" {
			fmt.Printf("合约[%s]不存在\n", contractName)
			fmt.Println("====================== 创建合约 ======================")
			byteCode, err := os.ReadFile("./contracts/token/token.bin")
			if err != nil {
				log.Fatalln(err)
			}
			createPayload, err := cc.client.CreateContractCreatePayload(contractName, version, string(byteCode), common.RuntimeType_EVM, []*common.KeyValuePair{})
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
			// fmt.Printf("blockHeight:%d, txId:%s, result:%s, msg:%s\n\n", resp.TxBlockHeight, resp.TxId, resp.ContractResult.Result, resp.ContractResult.Message)
		} else {
			if err != nil {
				log.Fatalln(err)
			}
		}
	} else {
		fmt.Printf("合约已存在 %+v \n\n", contract)
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	abiJson, err := os.ReadFile(path.Join(dir, "contracts/token/token.abi"))
	if err != nil {
		log.Fatalln(err)
	}

	myAbi, err := abi.JSON(strings.NewReader(string(abiJson)))
	if err != nil {
		log.Fatalln(err)
	}

	//addr := evmutils.StringToAddress(address)
	addr := evmutils.BigToAddress(evmutils.FromDecimalString(address))

	methodName := "balanceOf"
	dataByte, err := myAbi.Pack(methodName, addr)
	if err != nil {
		log.Fatalln(err)
	}

	dataString := hex.EncodeToString(dataByte)

	kvs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(dataString),
		},
	}

	result, err := invokeTokenContractWithResult(cc.client, contractName, methodName, "", kvs, true)
	if err != nil {
		log.Fatalln(err)
	}

	balance, err := myAbi.Unpack(methodName, result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("addr [%s] => %+v\n", address, balance)
	return nil
}

func invokeTokenContractWithResult(client *sdk.ChainClient, contractName, method, txId string,
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
