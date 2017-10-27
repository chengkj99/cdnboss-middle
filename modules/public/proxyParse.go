package public

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type proxyTargetInfo struct {
	Target      []string `json:"target"`
	ProxyMode   string   `json:"ProxyMode"`
	PathRewrite bool     `json:"pathRewrite"`
}

// ProxyParse 将ReadFile(proxy.xx.json)内容解析，生成路由配置
func ProxyParse(e *echo.Echo) {
	proxyInfo := make(map[string]proxyTargetInfo)
	//getEnvFileName() 获得配置文件路径
	envFileName := GetEnvFileName()
	err := ReadFile(envFileName, &proxyInfo)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return
	}

	// fmt.Printf("proxyInfo: %#v\n", proxyInfo)
	for k, v := range proxyInfo {
		api := k
		target := v.Target
		PathRewrite := v.PathRewrite
		ProxyMode := v.ProxyMode
		routerGroup := e.Group(api)
		var proxyTargets []*middleware.ProxyTarget
		var ProxyConfig echo.MiddlewareFunc

		// 获取负载均衡所需的URL数组切片
		for _, value := range target {
			url, _ := url.Parse(value)
			proxyTargets = append(proxyTargets, &middleware.ProxyTarget{
				URL: url,
			})
		}
		if ProxyMode == "RandomBalancer" {
			// 随机
			ProxyConfig = middleware.Proxy(&middleware.RandomBalancer{
				Targets: proxyTargets,
			})
		} else {
			// 轮询
			ProxyConfig = middleware.Proxy(&middleware.RoundRobinBalancer{
				Targets: proxyTargets,
			})
		}

		// 透传代理
		routerGroup.GET("/*", nil, handleRequestURIMiddleware(PathRewrite, api), ProxyConfig)
		routerGroup.POST("/*", nil, handleRequestURIMiddleware(PathRewrite, api), ProxyConfig)
		routerGroup.DELETE("/*", nil, handleRequestURIMiddleware(PathRewrite, api), ProxyConfig)
		routerGroup.PUT("/*", nil, handleRequestURIMiddleware(PathRewrite, api), ProxyConfig)
		// 自主拦截处理
		Intercept(api, routerGroup)
	}
}

// handleRequestURIMiddleware 处理RequestURI，是否被重写
func handleRequestURIMiddleware(PathRewrite bool, api string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			url := ctx.Request().RequestURI
			if PathRewrite {
				r, _ := regexp.Compile(api)
				ctx.Request().URL.Path = r.ReplaceAllString(url, "")
			}
			return next(ctx)
		}
	}
}
