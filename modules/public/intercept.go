package public

import (
	"cdnboss-middle/modules/alarm"

	"github.com/labstack/echo"
)

// Intercept 拦截路由处理配置
func Intercept(k string, g *echo.Group) {
	if k == "/alarm" {
		alarm.User(g)
	}
}
