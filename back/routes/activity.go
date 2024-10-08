package routes

import (
	"back/config"
	"back/controllers"
	middleware "back/middlewares"
	"back/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterActivityRoutes(r *gin.Engine, db *gorm.DB) {
	// Initialiser les repositories et controllers
	activityRepo := repositories.NewActivityRepositoryImpl(db)
	activityController := controllers.NewActivityController(activityRepo)

	// Routes pour les activités
	r.POST("/activities", middleware.AuthMiddleware(config.RoleAdmin, config.RoleStandLeader), activityController.CreateActivity)
	r.GET("/activities/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleStudent, config.RoleStandLeader), activityController.GetActivityByID)
	r.DELETE("/activities/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleOrganizer, config.RoleStandLeader), activityController.DeleteActivity)
	r.GET("/stands/:stand_id/activities", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent, config.RoleOrganizer, config.RoleStandLeader, config.RoleStudent), activityController.GetActivitiesByStandID)
}
