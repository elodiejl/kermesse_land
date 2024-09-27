package routes

import (
	"back/controllers"
	"back/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// Initialiser les repositories et controllers
	ticketRepo := repositories.NewTicketRepository(db)
	ticketController := controllers.NewTicketController(ticketRepo)

	// Routes pour les tickets
	r.POST("/tickets", ticketController.CreateTicket)                           // Créer un ticket
	r.GET("/tickets/:id", ticketController.GetTicketByID)                       // Récupérer un ticket par ID
	r.GET("/tombola/:tombola_id/tickets", ticketController.GetTicketsByTombola) // Récupérer tous les tickets d'une tombola
	r.DELETE("/tickets/:id", ticketController.DeleteTicket)                     // Supprimer un ticket
}
