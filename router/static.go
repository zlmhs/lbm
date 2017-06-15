package router

import (
	"github.com/gin-gonic/gin"
)

func initStatic(router *gin.Engine) {

	router.Static("/device/css", "./static/css")
	router.Static("/device/images", "./static/images")
	router.Static("/device/js", "./static/js")
	router.Static("/device/ueditor", "./static/ueditor")
}
