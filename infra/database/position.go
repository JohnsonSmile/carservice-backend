package database

import (
	"carservice/model"
	"errors"
)

func GetPosition(orderTypeID int, cityID string, regionID string, placeID string) (*model.Position, error) {
	position := &model.Position{}
	if rowsAffected := db.Model(&model.Position{}).Preload("City").Preload("Region").Preload("Place").Where("order_type_id = ? AND city_id = ? AND region_id = ? AND place_id = ?", orderTypeID, cityID, regionID, placeID).First(position).RowsAffected; rowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	return position, nil
}

func GetPositionByID(id int) (*model.Position, error) {
	position := &model.Position{}
	if rowsAffected := db.Model(&model.Position{}).Preload("City").Preload("Region").Preload("Place").Where("id = ?", id).First(position).RowsAffected; rowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	return position, nil
}

// city: 1, region: 1, place: 1001
// 插入城市信息
func CreateCity(city *model.City) error {
	if rowsAffected := db.Create(city).RowsAffected; rowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// 插入区域信息
func CreateRegion(region *model.Region) error {
	if rowsAffected := db.Create(region).RowsAffected; rowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// 插入地点信息
func CreatePlace(place *model.Place) error {
	if rowsAffected := db.Create(place).RowsAffected; rowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func CreatePosition(position *model.Position) error {
	if rowsAffected := db.Create(position).RowsAffected; rowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}
