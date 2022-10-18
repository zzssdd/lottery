package route

import (
	"github.com/gin-gonic/gin"
	"micro-web/control/activity"
	"micro-web/control/choose"
	"micro-web/control/order"
	"micro-web/control/prize"
	"micro-web/control/user"
	"micro-web/middleware"
)

func InitRouter(router *gin.Engine) {
	router.Use(middleware.Cors())
	router_user := router.Group("/user")
	router_prize := router.Group("/prize")
	router_order := router.Group("/order")
	router_choose := router.Group("/choose")
	router_activity := router.Group("/activity")

	user.InitRoute(router_user)
	prize.InitRoute(router_prize)
	order.InitRoute(router_order)
	choose.InitRoute(router_choose)
	activity.InitRoute(router_activity)

	router.Static("/upload", "./upload")
}
