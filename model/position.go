package model

type Position struct {
	BaseModel
	Name        string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	OrderTypeID int    `gorm:"column:order_type_id;type:tinyint(1);not null;default 0 comment '订单类型:0-highway;1-charge;3-park'" json:"order_type_id"`
	CityID      int    `gorm:"column:city_id;type:int(11);not null;default 0" json:"city_id"`
	City        City   `gorm:"foreignKey:CityID" json:"city"`
	RegionID    int    `gorm:"column:region_id;type:int(11);not null;default 0" json:"region_id"`
	Region      Region `gorm:"foreignKey:RegionID" json:"region"`
	PlaceID     int    `gorm:"column:place_id;type:int(11);not null;default 0" json:"place_id"`
	Place       Place  `gorm:"foreignKey:PlaceID" json:"place"`
}

type City struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
}

type Region struct {
	BaseModel
	Name   string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	City   City   `gorm:"foreignKey:CityID" json:"city"`
	CityID int    `gorm:"column:city_id;type:int(11);not null;default 0" json:"city_id"`
}

type Place struct {
	BaseModel
	PlaceType int     `gorm:"column:place_type;type:tinyint(1);not null;default 0 comment '地点类型:1-highway;2-charge;3-park'" json:"place_type"`
	Name      string  `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Region    Region  `gorm:"foreignKey:RegionID" json:"region"`
	RegionID  int     `gorm:"column:region_id;type:int(11);not null;default 0" json:"region_id"`
	Longitude float64 `gorm:"column:longitude;type:decimal(10,6);not null;default 0" json:"longitude"`
	Latitude  float64 `gorm:"column:latitude;type:decimal(10,6);not null;default 0" json:"latitude"`
}
