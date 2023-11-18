package database

import "carservice/model"

// TODO: start id 和 end id 应该是联合唯一键
func CreateHighWayFee(fee *model.HighWayFee) error {
	return db.Create(fee).Error
}

func GetHighwayFee(startPositionID, endPositionID int) (int, error) {
	var fee int
	err := db.Model(&model.HighWayFee{}).Select("Fee").Where("start_position_id = ? and end_position_id = ?", startPositionID, endPositionID).First(&fee).Error
	return fee, err

}

func CreateChargeFee(fee *model.ChargeFee) error {
	return db.Create(fee).Error
}

func GetChargeFee(positionID int) (int, error) {
	var fee int
	err := db.Model(&model.ChargeFee{}).Select("FeePerDegree").Where("position_id = ?", positionID).First(&fee).Error
	return fee, err
}
