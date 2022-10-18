package choose

import (
	"github.com/gin-gonic/gin"
	"micro-web/middleware"
)

func InitRoute(route *gin.RouterGroup) {
	route.POST("/do", middleware.UserJwt, Choose)
	route.GET("/result", middleware.UserJwt, ChooseResult)
}
