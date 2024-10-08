package routes

import (
	"back/config"
	"back/controllers"
	middleware "back/middlewares"
	"back/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterTicketRoutes(r *gin.Engine, ticketRepo repositories.TicketRepository) {
	// Initialiser les repositories et controllers
	ticketController := controllers.NewTicketController(ticketRepo)

	// Routes pour les tickets
	r.POST("/tickets", middleware.AuthMiddleware(config.RoleAdmin, config.RoleOrganizer), ticketController.CreateTicket)                                            // Créer un ticket
	r.GET("/tickets/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent, config.RoleOrganizer, config.RoleStudent), ticketController.GetTicketByID) // Récupérer un ticket par ID
	r.GET("/tombola/:tombola_id/tickets", middleware.AuthMiddleware(config.RoleAdmin, config.RoleOrganizer), ticketController.GetTicketsByTombola)                  // Récupérer tous les tickets d'une tombola
	r.DELETE("/tickets/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleOrganizer), ticketController.DeleteTicket)                                      // Supprimer un ticket
}
