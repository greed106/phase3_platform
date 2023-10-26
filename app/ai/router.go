package ai

import (
	"github.com/gin-gonic/gin"
	// 其他必要的导入，例如你的handler包、models包等
)

func AiRouter(r *gin.Engine) {
	// 子路由组 "/detection-info"
	detection := r.Group("/detection-info")
	{
		detection.POST("/detInfo", AddDetectionInfoHandler)      // 添加一条检测信息
		detection.GET("/detInfo", GetDetectionInfoByIDHandler)   // 根据ID获取检测信息
		detection.PUT("/detInfo", UpdateDetectionInfoHandler)    // 更新检测信息
		detection.DELETE("/detInfo", DeleteDetectionInfoHandler) // 删除检测信息
	}
}
