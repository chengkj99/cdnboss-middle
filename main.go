package main

import (
	"cdnboss-middle/modules/public"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	accCdnboss, err := os.OpenFile("acc_cdnboss.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer accCdnboss.Close()
	e.Logger.SetOutput(accCdnboss)
	e.Logger.SetLevel(log.DEBUG)

	// 日志输出格式
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: accCdnboss,
	}))

	// 异常处理
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			ctx.JSON(code, he)
		}
	}

	// 实现代理：透传和自定义处理
	public.ProxyParse(e)
	fmt.Println("server runing: localhost:1323")
	e.Logger.Fatal(e.Start(":1323"))
}
