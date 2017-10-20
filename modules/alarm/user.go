package alarm

import (
	"github.com/labstack/echo"
)

// Test 路由拦截处理Demo
func Test(rg *echo.Group) {
	rg.GET("/test", getUser)
}

func getUser(c echo.Context) error {
	c.String(200, "hello world")
	return nil
}
