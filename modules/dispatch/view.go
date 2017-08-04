package dispatch

import (
	"cdnboss-middle/modules/public"

	"github.com/gin-gonic/gin"
)

// View 获取所有view   doType: glb
func View(rg *gin.RouterGroup) {

	rg.GET("/:doType/views", public.ReqRelay(host))
}
