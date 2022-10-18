package prize

import (
	"github.com/gin-gonic/gin"
	"micro-web/middleware"
)

func InitRoute(route *gin.RouterGroup) {
	route.GET("/list", PrizeList)
	route.POST("/add", middleware.AdminJwt, PrizeAdd)
	route.DELETE("/del", middleware.AdminJwt, PrizeDel)
	route.PUT("/edit", middleware.AdminJwt, PrizeEdit)
	route.GET("/info", PrizeInfo)
}
