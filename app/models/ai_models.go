package models

import (
	"time"
)

// DetectionRequest ai小组传递来的数据
type DetectionRequest struct {
	ImagePaths      []string  `gorm:"column:image_path" json:"image_paths"`            // 检测图片在文件夹中的位置
	FieldID         int       `gorm:"column:field_id" json:"field_id"`                 // 对应农田的编号
	DroneID         int       `gorm:"column:drone_id" json:"drone_id"`                 // 对应的无人机编号
	DetectionTime   time.Time `gorm:"column:detection_time" json:"detection_time"`     // 图片检测时间
	PestTypeID      int       `gorm:"column:pest_type_id" json:"pest_type_id"`         // 病虫害的类型编号
	PestProbability float64   `gorm:"column:pest_probability" json:"pest_probability"` // 该片农田发生病虫害的概率
}

// PestInfo pest_info表中的数据实体
type PestInfo struct {
	ID              int       `gorm:"primaryKey;column:id" json:"id"`                  // 病虫害类型编号
	PestName        string    `gorm:"column:pest_name" json:"pest_name"`               // 病虫害名称
	PestDescription string    `gorm:"column:pest_description" json:"pest_description"` // 病虫害描述
	PestPrevention  string    `gorm:"column:pest_prevention" json:"pest_prevention"`   // 病虫害防治方法
	PestImage       string    `gorm:"column:pest_image" json:"pest_image"`             // 病虫害图片
	CreatedAt       time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"` // 创建时间
	UpdatedAt       time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"` // 更新时间
}

// DetectionInfo detection_info表中的数据实体
type DetectionInfo struct {
	ID              int       `gorm:"primaryKey;column:id" json:"id"`                  // 检测信息编号
	DetectionRequest
	CreatedAt       time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"` // 创建时间
	UpdatedAt       time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"` // 更新时间
}
