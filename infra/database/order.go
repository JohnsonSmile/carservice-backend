package database

import (
	"carservice/model"
	"errors"

	"gorm.io/gorm"
)

func GetLatestUnFinishedHighwayOrderByUserId(userId uint) (*model.Order, error) {
	var order model.Order
	// order_type_id: 	1-highway;2-charge;3-park
	// order_status: 	0-start;1-end;2-payed;
	err := db.Model(&model.Order{}).Preload("StartPosition").Preload("EndPosition").Where("user_id = ? AND order_type_id = ? AND order_status != ?", userId, 1, 2).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &order, nil
}

func GetHighwayOrderByOrderSn(orderSn string) (*model.Order, error) {
	var order model.Order
	err := db.Model(&model.Order{}).Preload("StartPosition").Preload("EndPosition").Where("order_sn = ?", orderSn).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func CreateOrder(order *model.Order) error {
	if rowsAffected := db.Create(order).RowsAffected; rowsAffected == 0 {
		return errors.New("insert failed")
	}
	// TODO:...
	return db.Model(&model.Order{}).Preload("StartPosition").Where("id = ?", order.ID).First(order).Error
}

func EndOrder(order *model.Order) error {
	if rowsAffected := db.Model(order).Updates(model.Order{
		OrderStatus:   1,
		EndAt:         order.EndAt,
		EndPositionID: order.EndPositionID,
		Fee:           order.Fee,
	}).RowsAffected; rowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

func CreateOrderType(orderType *model.OrderType) error {
	if rowsAffected := db.Create(orderType).RowsAffected; rowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func GetHighwayOrders(page, size int) (orders []*model.Order, total int64, err error) {
	orders = make([]*model.Order, 0)
	if err := db.Scopes(model.Paginate(page, size)).Model(&model.Order{}).Preload("StartPosition").Preload("EndPosition").Where("order_type_id = ?", 1).Order("start_at desc").Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Model(&model.Order{}).Where("order_type_id = ?", 1).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return orders, total, nil
}
