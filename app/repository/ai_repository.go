package repository

import (
	"phase3_platform/app/models"
	"phase3_platform/database"
)

// CreateDetectionInfo 创建检测信息
func CreateDetectionInfo(detInfo *models.DetectionInfo) (err error) {
	if err = database.DB.Create(&detInfo).Error; err != nil {

		return err
	}
	return nil
}

// GetDetectionInfoByID 根据ID获取检测信息
func GetDetectionInfoByID(id int) (detInfo models.DetectionInfo, err error) {
	if err = database.DB.Where("id = ?", id).First(&detInfo).Error; err != nil {
		return detInfo, err
	}
	return detInfo, nil
}

// UpdateDetectionInfo 更新检测信息
func UpdateDetectionInfo(id int, updatedDetInfo *models.DetectionInfo) (err error) {
	if err = database.DB.Model(&models.DetectionInfo{}).Where("id = ?", id).Updates(updatedDetInfo).Error; err != nil {
		return err
	}
	return nil
}

// DeleteDetectionInfo 删除检测信息
func DeleteDetectionInfo(id int) (err error) {
	if err = database.DB.Where("id = ?", id).Delete(&models.DetectionInfo{}).Error; err != nil {
		return err
	}
	return nil
}

// GetAllDetectionInfos 获取所有的检测信息
func GetAllDetectionInfos() (detInfos []models.DetectionInfo, err error) {
	if err = database.DB.Find(&detInfos).Error; err != nil {
		return nil, err
	}
	return detInfos, nil
}
