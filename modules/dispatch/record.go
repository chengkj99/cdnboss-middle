package dispatch

import (
	"cdnboss-middle/modules/public"

	"github.com/gin-gonic/gin"
)

// Record 解析记录接口 doType: glb - dispatch
func Record(rg *gin.RouterGroup) {

	// 1. 通过id获取解析记录  recordId: recordId
	// 2. 获取解析记录详情 recordId: detail
	// 3. 获取dnspod,qasn记录对比接口 recordId: sync
	rg.GET("/:doType/records/:recordId", public.ReqRelay(host))

	// 获取缓存组涉及的解析
	rg.GET("/:doType/cacheGroups/:id/records", public.ReqRelay(host))

	// 新建解析记录
	rg.POST("/:doType/records", public.ReqRelay(host))

	// 获取符合条件的解析记录
	rg.POST("/:doType/records/search", public.ReqRelay(host))

	// 重试同步dnspod,qans接口
	rg.POST("/:doType/records/sync", public.ReqRelay(host))

	// 修改解析记录
	rg.PUT("/:doType/records/:recordId", public.ReqRelay(host))

	// 删除解析记录
	rg.DELETE("/:doType/records/:recordId", public.ReqRelay(host))

}
