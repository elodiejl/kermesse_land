package routes

import (
	"back/controllers"
	middleware "back/middlewares"
	"back/repositories"

	"back/config"
	"github.com/gin-gonic/gin"
)

func KermesseRoutes(router *gin.Engine, kermesseRepo repositories.KermesseRepository) {
	kermesseController := controllers.NewKermesseController(kermesseRepo)

	kermesseRoutes := router.Group("/kermesses")
	{
		// Route pour récupérer toutes les kermesses (accessible à tout le monde)
		kermesseRoutes.GET("", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent, config.RoleOrganizer, config.RoleStandLeader, config.RoleStudent), kermesseController.GetAllKermesses)

		// Route pour créer une nouvelle kermesse (accessible uniquement aux organisateurs et aux admins)
		kermesseRoutes.POST("", middleware.AuthMiddleware(config.RoleAdmin, config.RoleOrganizer), kermesseController.CreateKermesse)

		// Route pour récupérer une kermesse par ID
		kermesseRoutes.GET("/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent, config.RoleOrganizer, config.RoleStandLeader, config.RoleStudent), kermesseController.GetKermesseByID)

		// Route pour supprimer une kermesse par ID (accessible uniquement aux organisateurs et aux admins)
		kermesseRoutes.DELETE("/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleOrganizer), kermesseController.DeleteKermesse)
	}
}
