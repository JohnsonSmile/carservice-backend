package chain

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"math/big"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

// [{Id:+1 OrderID:+1 OrderSn:TESTSZ1000230011 OrderType:0 StartAt:+1700377074 EndAt:+1700378074 UserID:+1 Fee:+1000 UniteCount:+1 IsPayed:false BlockHeight:+138 BlockTimestamp:+1700454270 StartPosition:火星1号高速A EndPosition:兰桂坊高速B}]
type OrderInfo struct {
	Id             *big.Int
	OrderSn        string
	OrderType      int
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

// orderType 从0开始的，应该是数据库里面的-1
func (cc *ChainClient) CreateOrder(orderId int, orderSn string, orderType uint8, startAt int, endAt int, userID int, fee int, uniteCount int, isPayed bool, startPosition string, endPosition string) error {

	methodName := "CreateOrder"

	dataByte, err := cc.orderABI.Pack(methodName, big.NewInt(int64(orderId)), orderSn, orderType, big.NewInt(int64(startAt)), big.NewInt(int64(endAt)), big.NewInt(int64(userID)), big.NewInt(int64(fee)), big.NewInt(int64(uniteCount)), isPayed, startPosition, endPosition)
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

	result, err := invokeContractWithResult(cc.client, cc.orderContract.Name, methodName, "", kvs, true)
	if err != nil {
		return err
	}

	_, err = cc.orderABI.Unpack(methodName, result)
	if err != nil {
		return err
	}

	return nil
}

func (cc *ChainClient) PayOrder(orderId int, userId int) error {

	methodName := "PayOrder"

	dataByte, err := cc.orderABI.Pack(methodName, big.NewInt(int64(orderId)), big.NewInt(int64(userId)))
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

	result, err := invokeContractWithResult(cc.client, cc.orderContract.Name, methodName, "", kvs, true)
	if err != nil {
		return err
	}

	_, err = cc.orderABI.Unpack(methodName, result)
	if err != nil {
		return err
	}

	return nil
}

// orderType 从0开始的，应该是数据库里面的-1
func (cc *ChainClient) GetOrder(orderId int) (*OrderInfo, error) {

	methodName := "GetOrder"

	dataByte, err := cc.orderABI.Pack(methodName, big.NewInt(int64(orderId)))
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

	result, err := invokeContractWithResult(cc.client, cc.orderContract.Name, methodName, "", kvs, true)
	if err != nil {
		return nil, err
	}

	// [{Id:+1 OrderID:+1 OrderSn:TESTSZ1000230011 OrderType:0 StartAt:+1700377074 EndAt:+1700378074 UserID:+1 Fee:+1000 UniteCount:+1 IsPayed:false BlockHeight:+138 BlockTimestamp:+1700454270 StartPosition:火星1号高速A EndPosition:兰桂坊高速B}]
	orderInfos, err := cc.orderABI.Unpack(methodName, result)
	if err != nil {
		return nil, err
	}
	log.Printf("orderInfo: %+v\n", orderInfos)

	if len(orderInfos) > 0 {
		dataBytes, err := json.Marshal(orderInfos[0])
		if err != nil {
			return nil, err
		}
		var orderInfo OrderInfo
		err = json.Unmarshal(dataBytes, &orderInfo)
		if err != nil {
			return nil, err
		}
		return &orderInfo, nil
	}
	return nil, errors.New("get order info error")
}
