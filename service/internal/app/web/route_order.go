package web

import "github.com/gin-gonic/gin"

func (s *restService) setOrderAPIRoutes(parentRouteGroup *gin.RouterGroup) {
	group := parentRouteGroup.Group("")
	group.Use(s.in.MemberAuthMiddleware.Authentication)

	group.POST("/order", s.in.OrderCtrl.Order)
	group.GET("/orders", s.in.OrderCtrl.GetOrderList)
}
