package main

import (
	"fmt"
	"net/http"
	"net/url"
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

	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			ctx.JSON(code, he)
			fmt.Println("code, he.Message..", code, he.Message)
		}
	}

	// public.ProxyParse(e)
	url1, _ := url.Parse("http://100.100.84.92:9090")
	url2, _ := url.Parse("http://100.100.84.92:9090")
	ProxyConfig := middleware.Proxy(&middleware.RoundRobinBalancer{
		Targets: []*middleware.ProxyTarget{
			&middleware.ProxyTarget{
				URL: url1,
			},
			&middleware.ProxyTarget{
				URL: url2,
			},
		},
	})
	fmt.Println("ProxyConfig...", &ProxyConfig)
	// e.Use()
	e.Logger.Fatal(e.Start(":1323"))
}
