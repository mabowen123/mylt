package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebController struct {
}

func (WebController WebController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
