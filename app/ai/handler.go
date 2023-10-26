package ai

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phase3_platform/app/models"

	"strconv"
)

// AddDetectionInfoHandler 添加一条检测信息
func AddDetectionInfoHandler(c *gin.Context) {
	var detInfo models.DetectionInfo
	if err := c.BindJSON(&detInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResultError("Invalid request data"))
		return
	}
	if err := AddDetectionInfo(&detInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResultError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.NewResultSuccessWithoutData())
}

// GetDetectionInfoByIDHandler 根据ID获取检测信息
func GetDetectionInfoByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewResultError("Invalid ID"))
		return
	}
	detInfo, err := GetDetectionInfoByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewResultError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.NewResultSuccessWithData(detInfo))
}

// UpdateDetectionInfoHandler 更新检测信息
func UpdateDetectionInfoHandler(c *gin.Context) {
	var updatedDetInfo models.DetectionInfo
	if err := c.BindJSON(&updatedDetInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResultError("Invalid request data"))
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewResultError("Invalid ID"))
		return
	}
	if err := UpdateDetectionInfo(id, &updatedDetInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResultError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.NewResultSuccessWithoutData())
}

// DeleteDetectionInfoHandler 删除检测信息
func DeleteDetectionInfoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewResultError("Invalid ID"))
		return
	}
	if err := DeleteDetectionInfo(id); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResultError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.NewResultSuccessWithoutData())
}
