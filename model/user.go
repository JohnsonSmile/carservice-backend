package model

type User struct {
	BaseModel
	Phone    string `gorm:"column:phone;uniqueIndex;type:varchar(50);not null;default '' comment '用户电话号码'" json:"phone"`
	Username string `gorm:"column:username;type:varchar(80);not null;default '' comment '用户名'" json:"username"`
	Password string `gorm:"column:password;type:varchar(80);not null;default '' comment '密码'" json:"password"`
	Avatar   string `gorm:"column:avatar;type:varchar(2046);not null;default '' comment '头像'" json:"avatar"`
	Bio      string `gorm:"column:bio;type:varchar(2046);not null;default '' comment '个人简介'" json:"bio"`
}
