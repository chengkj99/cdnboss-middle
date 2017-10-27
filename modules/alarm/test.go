package alarm

import (
	"github.com/labstack/echo"
)

// Test Demo 路由拦截处理
func Test(rg *echo.Group) {
	rg.GET("/test", getTest)
	rg.POST("/test", postTest)
	rg.PUT("/test", putTest)
	rg.DELETE("/test", deleteTest)
}

func getTest(c echo.Context) error {
	c.String(200, "hello GET")
	return nil
}
func postTest(c echo.Context) error {
	c.String(200, "hello POST")
	return nil
}
func putTest(c echo.Context) error {
	c.String(200, "hello PUT")
	return nil
}
func deleteTest(c echo.Context) error {
	c.String(200, "hello DELETE")
	return nil
}
