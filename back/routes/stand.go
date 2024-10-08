package routes

import (
	"back/config"
	"back/controllers"
	middleware "back/middlewares"
	"back/repositories"

	"github.com/gin-gonic/gin"
)

func StandRoutes(router *gin.Engine, standRepo repositories.StandRepository) {
	standController := controllers.NewStandController(standRepo)

	standRoutes := router.Group("/stands")
	{
		standRoutes.POST("", middleware.AuthMiddleware(config.RoleAdmin, config.RoleStandLeader), standController.CreateStand)
		standRoutes.GET("/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent, config.RoleStudent, config.RoleOrganizer, config.RoleStandLeader), standController.GetStandByID)
		standRoutes.GET("/kermesse/:kermesse_id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent, config.RoleStudent, config.RoleOrganizer, config.RoleStandLeader), standController.GetStandsByKermesse)
		standRoutes.DELETE("/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleStandLeader), standController.DeleteStand)
	}
}
