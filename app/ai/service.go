package ai

import (
	"github.com/sirupsen/logrus"
	"phase3_platform/app/models"
	"phase3_platform/app/repository"
)

// AddDetectionInfo 添加一条检测信息
func AddDetectionInfo(detInfo *models.DetectionInfo) (err error) {
	if err = repository.CreateDetectionInfo(detInfo); err != nil {
		logrus.Errorf("添加检测信息失败 %v", err)
		return err
	}
	// 输出日志
	logrus.Infof("添加检测信息成功, ID: %d", detInfo.ID)
	return nil
}

// GetDetectionInfoByID 根据ID获取检测信息
func GetDetectionInfoByID(id int) (detInfo models.DetectionInfo, err error) {
	if detInfo, err = repository.GetDetectionInfoByID(id); err != nil {
		logrus.Errorf("获取检测信息失败 %v", err)
		return detInfo, err
	}
	// 输出日志
	logrus.Infof("获取检测信息成功, ID: %d", detInfo.ID)
	return detInfo, nil
}

// UpdateDetectionInfo 更新检测信息
func UpdateDetectionInfo(id int, updatedDetInfo *models.DetectionInfo) (err error) {
	if err = repository.UpdateDetectionInfo(id, updatedDetInfo); err != nil {
		logrus.Errorf("更新检测信息失败 %v", err)
		return err
	}
	// 输出日志
	logrus.Infof("更新检测信息成功, ID: %d", id)
	return nil
}

// DeleteDetectionInfo 删除检测信息
func DeleteDetectionInfo(id int) (err error) {
	var detInfo models.DetectionInfo
	if detInfo, err = repository.GetDetectionInfoByID(id); err != nil {
		logrus.Errorf("待删除的检测信息ID不存在 %v", err)
	}
	if err = repository.DeleteDetectionInfo(id); err != nil {
		logrus.Errorf("删除检测信息失败 %v", err)
		return err
	}
	// 输出日志
	// 我认为这是一个比较敏感的操作，所以输出了删除的检测信息的内容
	logrus.Infof("删除检测信息成功, 信息内容为：%v", detInfo)
	return nil
}
