package database_test

import (
	"carservice/infra/config"
	"carservice/infra/database"
	"carservice/infra/logger"
	"carservice/model"
	"testing"
)

func TestMain(m *testing.M) {
	logger.Initialize(true)
	config.Initialize(true)
	database.Initialize(true)
	m.Run()
}

func TestPrePareHighwayData(t *testing.T) {

	// order types
	// 1-highway;2-charge;3-park
	orders := []model.OrderType{
		{
			Name:       "高速公路",
			FeePerUnit: 500, // 5.00
			Unit:       "公里",
		},
		{
			Name:       "充电桩",
			FeePerUnit: 400, // 4.00
			Unit:       "度",
		},
		{
			Name:       "停车场",
			FeePerUnit: 1000, // 10.00
			Unit:       "小时",
		},
	}

	for idx, order := range orders {
		order.ID = idx + 1 // 1, 2, 3
		if err := database.CreateOrderType(&order); err != nil {
			t.Error(err)
			return
		}
	}

	// highStartCity: 1, region: 1, place: 1001
	highStartCity := model.City{
		Name: "深圳",
	}
	highStartCity.ID = 1

	if err := database.CreateCity(&highStartCity); err != nil {
		t.Error(err)
		return
	}

	highStartRegion := model.Region{
		Name:   "南山区",
		CityID: 1,
	}

	if err := database.CreateRegion(&highStartRegion); err != nil {
		t.Error(err)
		return
	}

	highStartPlace := model.Place{
		Name:      "宝安中心-001",
		RegionID:  1,
		Longitude: 24.15,
		Latitude:  113.23,
		PlaceType: 1,
	}
	highStartPlace.ID = 1001

	if err := database.CreatePlace(&highStartPlace); err != nil {
		t.Error(err)
		return
	}

	// high start position
	highStartPosition := model.Position{
		PlaceID:  1001,
		CityID:   1,
		RegionID: 1,
		Name:     "宝安中心-金城路A",
		// 1-highway;2-charge;3-park
		OrderTypeID: 1,
	}

	if err := database.CreatePosition(&highStartPosition); err != nil {
		t.Error(err)
		return
	}

	// high end position
	// highEndCity: 2, region: 2, place: 1011
	// 香港， 九龙， 尖沙咀-A
	highEndCity := model.City{
		Name: "香港",
	}
	highEndCity.ID = 2
	if err := database.CreateCity(&highEndCity); err != nil {
		t.Error(err)
		return
	}

	highEndRegion := model.Region{
		Name:   "九龙",
		CityID: 2,
	}

	if err := database.CreateRegion(&highEndRegion); err != nil {
		t.Error(err)
		return
	}

	highEndPlace := model.Place{
		Name:      "尖沙咀-A",
		RegionID:  2,
		Longitude: 23.15,
		Latitude:  123.23,
	}
	highStartPlace.ID = 1001

	if err := database.CreatePlace(&highEndPlace); err != nil {
		t.Error(err)
		return
	}

	// high end position
	highEndPosition := model.Position{
		PlaceID:  1011,
		CityID:   2,
		RegionID: 2,
		Name:     "九龙-尖沙咀-A",
		// 1-highway;2-charge;3-park
		OrderTypeID: 1,
	}

	if err := database.CreatePosition(&highEndPosition); err != nil {
		t.Error(err)
		return
	}

	// TODO: fee 可以按照起始的位置来收费，也可以通过地图sdk中地理位置计算的路段总长度来按公里收费，这里只按照不同点之间的收费来处理。
	fee := model.HighWayFee{
		StartPositionID: 1,
		EndPositionID:   2,
		Fee:             1004, // 10.04元
	}
	if err := database.CreateHighWayFee(&fee); err != nil {
		t.Error(err)
		return
	}

}

func TestPrePareChargeData(t *testing.T) {

	chargePlace := model.Place{
		Name:      "万象城-001",
		RegionID:  1,
		Longitude: 24.15,
		Latitude:  113.23,
		PlaceType: 2, // charge
	}
	chargePlace.ID = 11

	if err := database.CreatePlace(&chargePlace); err != nil {
		t.Error(err)
		return
	}

	// charge position
	chargePosition := model.Position{
		PlaceID:  11,
		CityID:   1,
		RegionID: 1,
		Name:     "宝安区-万象城-001",
		// 1-highway;2-charge;3-park
		OrderTypeID: 2,
	}

	if err := database.CreatePosition(&chargePosition); err != nil {
		t.Error(err)
		return
	}

	// TODO: fee 按时间收费
	fee := model.ChargeFee{
		PositionID:   3,
		FeePerDegree: 205,
	}
	if err := database.CreateChargeFee(&fee); err != nil {
		t.Error(err)
		return
	}
}
