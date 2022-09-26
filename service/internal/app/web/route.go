package web

import "github.com/gin-gonic/gin"

func (s *restService) setApiRouters(parentRouteGroup *gin.RouterGroup) {
	privateRouteGroup := parentRouteGroup.Group("")

	s.setMemberAPIRoutes(privateRouteGroup)
	s.setProductAPIRoutes(privateRouteGroup)
	s.setOrderAPIRoutes(privateRouteGroup)
}
