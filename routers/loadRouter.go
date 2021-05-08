package routers

import "github.com/gin-gonic/gin"

type Router struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(c *gin.Context)
}

type Routers []Router

func LoadRouter(r *gin.Engine) {
	loadWebRouter(r)
}
