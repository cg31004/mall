package web

import "github.com/gin-gonic/gin"

func (s *restService) setMemberAPIRoutes(parentRouteGroup *gin.RouterGroup) {
	// 登入功能
	loginGroup := parentRouteGroup.Group("")
	loginGroup.POST("/member/login", s.in.MemberCtrl.Login)
	loginGroup.POST("/member/logout", s.in.MemberAuthMiddleware.Authentication, s.in.MemberCtrl.Logout)

	// 購物車
	chartGroup := parentRouteGroup.Group("")
	chartGroup.Use(s.in.MemberAuthMiddleware.Authentication)

	chartGroup.GET("/member/chart", s.in.MemberCtrl.GetChart)
	chartGroup.POST("/member/chart", s.in.MemberCtrl.CreateChart)
	chartGroup.PUT("/member/chart", s.in.MemberCtrl.UpdateChart)
	chartGroup.DELETE("/member/chart", s.in.MemberCtrl.DeleteChart)
}
