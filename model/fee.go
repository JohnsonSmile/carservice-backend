package model

type HighWayFee struct {
	BaseModel
	StartPositionID int      `gorm:"column:start_position_id;type:int(11);not null;default 0 comment '起始位置id'" json:"start_position_id"`
	StartPosition   Position `gorm:"foreignKey:StartPositionID;references:ID" json:"start_position"`
	EndPositionID   int      `gorm:"column:end_position_id;type:int(11);not null;default 0 comment '终止位置id'" json:"end_position_id"`
	EndPosition     Position `gorm:"foreignKey:EndPositionID;references:ID" json:"end_position"`
	Fee             int      `gorm:"column:fee;type:int(11);not null;default 0 comment '费用'" json:"fee"`
}

type ChargeFee struct {
	BaseModel
	PositionID   int      `gorm:"column:position_id;type:int(11);not null;default 0 comment '位置id'" json:"position_id"`
	Position     Position `gorm:"foreignKey:PositionID;references:ID" json:"position"`
	FeePerDegree int      `gorm:"column:fee_per_degree;type:int(11);not null;default 0 comment '费用/度'" json:"fee_per_degree"`
}

type ParkFee struct {
	BaseModel
	FeePerHour int `gorm:"column:fee_per_hour;type:int(11);not null;default 0 comment '费用/小时'" json:"fee_per_hour"`
}
