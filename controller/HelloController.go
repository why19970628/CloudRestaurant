package controller

import (
	"github.com/gin-gonic/gin"
)

type HelloController struct {
	
}

func (hello *HelloController)Router(engine *gin.Engine)  {
	engine.GET("/hello", hello.Hello)
}
func (hello *HelloController) Hello(c *gin.Context)  {
	c.JSON(200, map[string]interface{}{
		"meeage": "hello cloudrestaurant",
	})
}