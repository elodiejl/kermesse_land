package routes

import (
	"back/controllers"
	middleware "back/middlewares"
	"back/repositories"

	"back/config"
	"github.com/gin-gonic/gin"
)

func ParentRoutes(router *gin.Engine, parentRepo repositories.ParentRepository) {
	parentController := controllers.NewParentController(parentRepo)

	parentRoutes := router.Group("/parents")
	{
		//parentRoutes.GET("", middleware.AuthMiddleware(config.RoleAdmin), parentController.GetAllParents)
		parentRoutes.POST("", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent), parentController.CreateParent)
		parentRoutes.GET("/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent, config.RoleStudent, config.RoleOrganizer, config.RoleStandLeader), parentController.GetParentByID)
		parentRoutes.DELETE("/:id", middleware.AuthMiddleware(config.RoleAdmin), parentController.DeleteParent)
	}
}
