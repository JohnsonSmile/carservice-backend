package database

import (
	"carservice/model"
	"errors"

	"gorm.io/gorm"
)

func GetLatestUnFinishedHighwayOrderByUserId(userId uint) (*model.Order, error) {
	var order model.Order
	// order_type_id: 	0-highway;1-charge;3-park
	// order_status: 	0-start;1-end;2-payed;
	err := db.Where("user_id = ? AND order_type_id = ? AND order_status != ?", userId, 0, 2).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &order, nil
}
