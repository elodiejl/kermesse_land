package routes

import (
	"back/config"
	"back/controllers"
	middleware "back/middlewares"
	"back/repositories"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func TombolaRoutes(router *gin.Engine, tombolaRepo repositories.TombolaRepository, db *gorm.DB) {
	tombolaController := controllers.NewTombolaController(tombolaRepo)
	userRepo := repositories.NewUserRepository(db)
	tombolaController.UserRepo = userRepo

	// Routes pour les tombolas
	tombolaGroup := router.Group("/tombola")
	{
		tombolaGroup.POST("", middleware.AuthMiddleware(config.RoleAdmin, config.RoleOrganizer), tombolaController.CreateTombola)
		tombolaGroup.GET("/detail/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent, config.RoleStudent, config.RoleOrganizer, config.RoleStandLeader), tombolaController.GetTombolaByID)
		tombolaGroup.DELETE("/delete/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleOrganizer), tombolaController.DeleteTombola)
	}

	router.GET("/kermesse/:kermesse_id/tombolas", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent, config.RoleStudent, config.RoleStandLeader), tombolaController.GetTombolasByKermesse)
}
