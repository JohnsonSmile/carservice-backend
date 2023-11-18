package request

type RegisterRequest struct {
	Phone      string `json:"phone" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required"`
}

type LoginRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PreviewRequest struct {
	Data string `json:"data" binding:"required"`
}

type StartRequest struct {
	PositionID int `json:"position_id" binding:"required"`
}

type EndRequest struct {
	OrderSn    string `json:"order_sn" binding:"required"`
	PositionID int    `json:"position_id" binding:"required"` // end positionID
}

type ChargeEndRequest struct {
	ID int `json:"id" binding:"required"`
}

type HighwayOrdersRequest struct {
	Page int `form:"page" binding:"required"`
	Size int `form:"size" binding:"required"`
}
