package public

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/labstack/echo"
)

// ReqRelay 对前端请求透传处理
// host string, rewriteURL string
func ReqRelay(arguments ...string) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		url := handleArguments(ctx, arguments)
		req, err := http.NewRequest(ctx.Request().Method, url, ctx.Request().Body)
		if err != nil {
			ctx.Logger().Errorf("error: %v", err)
			return nil
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			ctx.Logger().Errorf("error: %v", err)
			return nil
		}
		defer res.Body.Close()
		ctx.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx.Response().Status = res.StatusCode
		_, err = io.Copy(ctx.Response().Writer, res.Body)
		if err != nil {
			ctx.Logger().Errorf("error: %v", err)
			return nil
		}
		return nil
	}
}

// 根据arguments的长度，判断是否含有第二个rewriteURL参数，如果有就删掉；
func handleArguments(ctx echo.Context, arguments []string) string {
	ArgLen := len(arguments)
	target := arguments[0]
	url := target + ctx.Request().RequestURI
	fmt.Println("url before:", url)
	if ArgLen > 1 {
		rewriteURL := arguments[1]
		r, _ := regexp.Compile(rewriteURL)
		url = r.ReplaceAllString(url, "")
	}
	fmt.Println("url after:", url)
	return url
}
