package chain

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"math/big"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	cm "github.com/ethereum/go-ethereum/common"
)

type UserInfo struct {
	Id    *big.Int
	Phone *big.Int
	Score *big.Int
}

func (cc *ChainClient) IsManager(address string) (bool, error) {

	managerRoleBytes := []byte("[175,41,13,134,128,130,10,173,146,40,85,243,155,48,96,151,178,14,40,119,77,108,26,211,90,32,50,86,48,195,160,44]")
	var roleBytes32 [32]byte
	err := json.Unmarshal(managerRoleBytes, &roleBytes32)
	if err != nil {
		return false, err
	}

	// grant role
	hasRoleMethodName := "hasRole"

	grantDataByte, err := cc.userABI.Pack(hasRoleMethodName, roleBytes32, cm.HexToAddress(address))
	if err != nil {
		return false, err
	}

	hasRoleDataString := hex.EncodeToString(grantDataByte)

	hasRoleKvs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(hasRoleDataString),
		},
	}

	hasRoleResult, err := invokeContractWithResult(cc.client, cc.userContract.Name, hasRoleMethodName, "", hasRoleKvs, true)
	if err != nil {
		return false, err
	}

	results, err := cc.userABI.Unpack(hasRoleMethodName, hasRoleResult)
	if err != nil {
		return false, err
	}

	if len(results) > 0 {
		dataBytes, err := json.Marshal(results[0])
		if err != nil {
			return false, err
		}
		var hasRoleResult bool
		err = json.Unmarshal(dataBytes, &hasRoleResult)
		if err != nil {
			return false, err
		}

		log.Printf("isManager result: %+v\n", hasRoleResult)
		return hasRoleResult, nil
	}
	return false, errors.New("get is manager error")
}

func (cc *ChainClient) GrantManager(address string) error {

	// managerRoleMethodName := "MANAGER"

	// dataByte, err := cc.userABI.Pack(managerRoleMethodName)
	// if err != nil {
	// 	return err
	// }

	// dataString := hex.EncodeToString(dataByte)

	// kvs := []*common.KeyValuePair{
	// 	{
	// 		Key:   "data",
	// 		Value: []byte(dataString),
	// 	},
	// }

	// result, err := invokeContractWithResult(cc.client, cc.userContract.Name, managerRoleMethodName, "", kvs, true)
	// if err != nil {
	//	return err
	// }

	// roleNames, err := cc.userABI.Unpack(managerRoleMethodName, result)
	// if err != nil {
	//	return err
	// }

	// log.Printf("role name: %+v\n", roleNames)
	// if len(roleNames) == 0 {
	// 	return errors.New("get role name error")
	// }
	// managerRoleBytes, err := json.Marshal(roleNames[0])
	// if err != nil {
	//	return err
	// }
	// log.Printf("managerRoleBytes: %+v\n", string(managerRoleBytes))
	managerRoleBytes := []byte("[175,41,13,134,128,130,10,173,146,40,85,243,155,48,96,151,178,14,40,119,77,108,26,211,90,32,50,86,48,195,160,44]")
	var roleBytes32 [32]byte
	err := json.Unmarshal(managerRoleBytes, &roleBytes32)
	if err != nil {
		return err
	}

	// grant role
	grantMethodName := "grantRole"

	grantDataByte, err := cc.userABI.Pack(grantMethodName, roleBytes32, cm.HexToAddress(address))
	if err != nil {
		return err
	}

	grantDataString := hex.EncodeToString(grantDataByte)

	grantKvs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(grantDataString),
		},
	}

	grantResult, err := invokeContractWithResult(cc.client, cc.userContract.Name, grantMethodName, "", grantKvs, true)
	if err != nil {
		return err
	}

	granted, err := cc.userABI.Unpack(grantMethodName, grantResult)
	if err != nil {
		return err
	}

	log.Printf("grant result: %+v\n", granted)

	return nil
}

func (cc *ChainClient) CreateUser(userId int, phone int, score int) error {

	methodName := "CreateUser"
	dataByte, err := cc.userABI.Pack(methodName, big.NewInt(int64(userId)), big.NewInt(int64(phone)), big.NewInt(int64(score)))
	if err != nil {
		return err
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
		return err
	}

	log.Printf("result: %+v\n", result)
	return nil
}

func (cc *ChainClient) GetUserInfo(userId int) (*UserInfo, error) {

	methodName := "GetUser"

	dataByte, err := cc.userABI.Pack(methodName, big.NewInt(int64(userId)))
	if err != nil {
		return nil, err
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
		return nil, err
	}

	userInfos, err := cc.userABI.Unpack(methodName, result)
	if err != nil {
		return nil, err
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
	return nil, errors.New("get user info error")
}

func (cc *ChainClient) ChargeScore(userId int, score int) error {

	methodName := "ChargeScore"

	dataByte, err := cc.userABI.Pack(methodName, big.NewInt(int64(userId)), big.NewInt(int64(score)))
	if err != nil {
		return err
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
		return err
	}

	_, err = cc.userABI.Unpack(methodName, result)
	if err != nil {
		return err
	}

	return nil
}

func (cc *ChainClient) PayScore(userId int, score int) error {

	methodName := "PayScore"

	dataByte, err := cc.userABI.Pack(methodName, big.NewInt(int64(userId)), big.NewInt(int64(score)))
	if err != nil {
		return err
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
		return err
	}

	_, err = cc.userABI.Unpack(methodName, result)
	if err != nil {
		return err
	}

	return nil
}
