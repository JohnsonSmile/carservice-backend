package model

import "time"

type Order struct {
	BaseModel
	OrderSn         string    `gorm:"column:order_sn;type:varchar(255);not null; default '' comment '订单号'" json:"order_sn"`
	OrderTypeID     int       `gorm:"column:order_type_id;type:tinyint(1);not null;default 0 comment '订单类型:0-highway;1-charge;3-park'" json:"order_type_id"`
	OrderStatus     int       `gorm:"column:order_status;type:tinyint(1);not null;default 0 comment '订单状态:0-start;1-end;2-payed;'" json:"order_status"`
	UserID          int       `gorm:"column:user_id;type:int(11);not null;default 0 comment '用户ID'" json:"user_id"`
	StartAt         time.Time `gorm:"column:created_at" json:"created_at"`
	EndAt           time.Time `gorm:"column:end_at" json:"end_at"`
	OrderType       OrderType `gorm:"foreignKey:OrderTypeID;references:ID" json:"order_type"`
	Fee             int       `gorm:"column:fee;type:int(11);not null;default 0 comment '订单费用'" json:"fee"` // highway根据起始位置id判断费用，不按照单位计费，其余则根据订单类型来同意收取，这里需要冗余设计，而不是查找，一般收费会动态调整，所以在订单中应该单独存储，其实充电和停车也可以设计一个单独计费的表格来通过地区或者单独的id来确定收费标准。
	StartPositionID int       `gorm:"column:start_position_id;type:int(11);not null;default 0 comment '起始位置id'" json:"start_position_id"`
	StartPosition   Position  `gorm:"foreignKey:StartPositionID;references:ID" json:"start_position"`
	EndPositionID   int       `gorm:"column:end_position_id;type:int(11);not null;default 0 comment '终止位置id'" json:"end_position_id"`
	EndPosition     Position  `gorm:"foreignKey:EndPositionID;references:ID" json:"end_position"`
}

type OrderType struct {
	BaseModel
	FeePerUnit int    `gorm:"column:fee_per_unit;type:int(11);not null;default 0 comment '单位费用:highway-0;charge-feePerDegree;park-feePerHour;'" json:"fee_per_unit"`
	Unit       string `gorm:"column:unit;type:varchar(10);not null;default 0 comment '单位:highway-km;charge-degree;park-hour;'" json:"unit"`
	Name       string `gorm:"column:name;type:varchar(10);not null;default 0 comment '名称:highway-高速;charge-充电;park-停车;'" json:"name"`
}
