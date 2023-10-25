package models

// FieldInfo field_info表中的数据实体
type FieldInfo struct {
	FieldID   int     `gorm:"column:field_id" json:"field_id"`   // 对应农田的编号
	Latitude  float64 `gorm:"column:latitude" json:"latitude"`   // 纬度
	Longitude float64 `gorm:"column:longitude" json:"longitude"` // 经度
}
