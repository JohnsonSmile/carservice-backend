package database

import (
	"carservice/model"
	"errors"
)

func GetPosition(cityID string, regionID string, placeID string) (*model.Position, error) {
	position := &model.Position{}
	if rowsAffected := db.Model(&model.Position{}).Preload("City").Preload("Region").Preload("Place").Where("city_id = ? AND region_id = ? AND place_id = ?", cityID, regionID, placeID).RowsAffected; rowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	return position, nil
}
