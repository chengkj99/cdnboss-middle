package modules

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dispatch(r *gin.Engine) {
	dis := r.Group("/dis")

	dis.GET("/v1/admin/common/pops", func(c *gin.Context) {
		res, err := http.Get("http://cs.zjmanageplatform.qiniu.io:8090/api/v1/admin/common/pops")
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		c.Header("Content-Type", "application/json; charset=utf-8")
		_, err = io.Copy(c.Writer, res.Body)
		if err != nil {
			fmt.Println(err)
		}
	})
}
