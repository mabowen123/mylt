package routers

import (
	"github.com/gin-gonic/gin"
	. "mylt/controllers"
)

var webRouters = Routers{
	Router{
		Name:        "主页",
		Pattern:     "/",
		HandlerFunc: WebController{}.Index,
	},
}

func loadWebRouter(r *gin.Engine) {
	for _, router := range webRouters {
		r.GET(router.Pattern, router.HandlerFunc)
	}
}
