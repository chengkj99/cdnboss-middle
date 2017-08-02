package main

import (
	"cdnboss-middle/modules/dispatch"

	"github.com/gin-gonic/gin"
)

var DB = make(map[string]string)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	// // Ping test
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.String(200, "hello wolrd")
	// })

	// // Get user value
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	// 	value, ok := DB[user]
	// 	if ok {
	// 		c.JSON(200, gin.H{"user": user, "value": value})
	// 	} else {
	// 		c.JSON(200, gin.H{"user": user, "status": "no value"})
	// 	}
	// })

	modules.Dispatch(r)
	// Listen and Server in 0.0.0.0:6060
	r.Run(":6060")
}
