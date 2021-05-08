package main

import (
	"github.com/gin-gonic/gin"
	"mylt/routers"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/html/*")
	r.Static("/static/js", "views/js")
	r.Static("/static/css", "views/css")
	r.Static("/image", "views/static")
	r.StaticFile("/favicon.ico", "views/static/favicon.ico")
	routers.LoadRouter(r)
	_ = r.Run(":9090")
}
