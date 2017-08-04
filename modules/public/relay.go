package public

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReqRelay 对前端请求不做处理直接转发的方法
func ReqRelay(host string) func(e *gin.Context) {
	return func(e *gin.Context) {
		url := host + e.Request.RequestURI
		req, err := http.NewRequest(e.Request.Method, url, e.Request.Body)
		if err != nil {
			fmt.Println(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		e.Header("Content-Type", "application/json; charset=utf-8")
		_, err = io.Copy(e.Writer, res.Body)
		if err != nil {
			fmt.Println(err)
		}
	}
}
