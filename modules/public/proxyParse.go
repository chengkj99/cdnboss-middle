package public

import (
	"fmt"

	"github.com/labstack/echo"
)

type proxyTargetInfo struct {
	Target      []string `json:"target"`
	PathRewrite bool     `json:"pathRewrite"`
}

// ProxyParse 将ReadFile(proxy.xx.json)内容解析，生成路由配置
func ProxyParse(e *echo.Echo) {
	proxyInfo := make(map[string]proxyTargetInfo)
	err := ReadFile("./conf/proxy.dev.json", &proxyInfo)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return
	}

	// fmt.Printf("proxyInfo: %#v\n", proxyInfo)
	for k, v := range proxyInfo {
		routerGroup := e.Group(k)
		Intercept(k, routerGroup)
		if v.PathRewrite {
			routerGroup.GET("/*", ReqRelay(v.Target, k))
		} else {
			routerGroup.GET("/*", ReqRelay(v.Target))
		}
	}
}
