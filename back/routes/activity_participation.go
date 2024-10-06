package routes

import (
	"back/controllers"
	"back/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterActivityParticipationRoutes(r *gin.Engine, db *gorm.DB) {

	participationRepo := repositories.NewActivityParticipationRepositoryImpl(db)
	participationController := controllers.NewActivityParticipationController(participationRepo)

	// Routes pour les participations aux activités
	r.POST("/participations", participationController.CreateParticipation)
	r.GET("/participations/:id", participationController.GetParticipationByID)
	r.GET("/users/:user_id/participations", participationController.GetParticipationsByUserID)
	r.DELETE("/participations/:id", participationController.DeleteParticipation)
}
