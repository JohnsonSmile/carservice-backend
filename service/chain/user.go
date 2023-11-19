package chain

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

type UserInfo struct {
	Id    *big.Int
	Phone *big.Int
	Score *big.Int
}

func (cc *ChainClient) CreateUser(userId int, phone int, score int) error {

	methodName := "CreateUser"
	dataByte, err := cc.userABI.Pack(methodName, userId, phone, score)
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

	result, err := invokeContractWithResult(cc.client, cc.userContract.Name, methodName, "", kvs, true)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("result: %+v\n", result)
	return nil
}

func (cc *ChainClient) GetUserInfo(userId int) (*UserInfo, error) {

	methodName := "GetUser"

	dataByte, err := cc.userABI.Pack(methodName, big.NewInt(int64(userId)))
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

	result, err := invokeContractWithResult(cc.client, cc.userContract.Name, methodName, "", kvs, true)
	if err != nil {
		log.Fatalln(err)
	}

	userInfos, err := cc.userABI.Unpack(methodName, result)
	if err != nil {
		log.Fatalln(err)
	}
	if len(userInfos) > 0 {
		dataBytes, err := json.Marshal(userInfos[0])
		if err != nil {
			return nil, err
		}
		var userInfo UserInfo
		err = json.Unmarshal(dataBytes, &userInfo)
		if err != nil {
			return nil, err
		}
		return &userInfo, nil
	}
	return nil, nil
}
