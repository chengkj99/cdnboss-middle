package common

import (
	"cdnboss-middle/modules/public"

	"github.com/gin-gonic/gin"
)

// 平台信息接口
func Platform(rg *gin.RouterGroup) {

	// 获取信息平台
	rg.GET("/platforms/:platformId", public.ReqRelay(host))

	// 获取信息平台列表
	rg.GET("/platforms", public.ReqRelay(host))

	// 获取信息平台下缓存组列表
	rg.GET("/platforms/:platformId/cacheGroups", public.ReqRelay(host))

	// 新建信息平台
	rg.POST("/platforms", public.ReqRelay(host))

	// 修改信息平台
	rg.PUT("/platforms/:platformId", public.ReqRelay(host))

}
