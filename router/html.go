package router
import (
	"github.com/gin-gonic/gin"
	"xsbPro/xsbDeviceH5/controllers"
)

func initHtml(router *gin.Engine)  {
	router.LoadHTMLGlob("views/*.html")

	//赛事
	router.GET("/device/matchNotice/:user", controllers.MatchNotice)
	router.GET("/device/matchAbstract/:user", controllers.MatchAbstract)
	router.GET("/device/matchCourseAbstract/:user", controllers.MatchCourseAbstract)

	//分享
	router.GET("/device/newShare/:user", controllers.NewShare)



	router.GET("/device/video/:user/:video", controllers.Video)

}
