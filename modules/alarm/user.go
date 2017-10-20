package alarm

import (
	"github.com/labstack/echo"
)

func User(rg *echo.Group) {
	rg.GET("/user", getUser)
}

func getUser(c echo.Context) error {
	c.String(200, "hello world")
	return nil
}
