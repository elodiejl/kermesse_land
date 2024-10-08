package routes

import (
	"back/controllers"
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
		tombolaGroup.POST("", tombolaController.CreateTombola)
		tombolaGroup.GET("/:id", tombolaController.GetTombolaByID)
		tombolaGroup.DELETE("/:id", tombolaController.DeleteTombola)
	}

	router.GET("/kermesse/:kermesse_id/tombolas", tombolaController.GetTombolasByKermesse)
}
