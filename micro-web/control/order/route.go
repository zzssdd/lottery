package order

import (
	"github.com/gin-gonic/gin"
	"micro-web/middleware"
)

func InitRoute(route *gin.RouterGroup) {
	route.GET("/list", middleware.UserJwt, OrderList)
	route.GET("/userList", middleware.UserJwt, OrderByUser)
	route.DELETE("/del", middleware.UserJwt, OrderDel)
}
