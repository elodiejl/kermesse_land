package routes

import (
	"back/config"
	"back/controllers"
	middleware "back/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterMessageRoutes(r *gin.Engine) {
	// Créer un contrôleur de chat
	messageController := controllers.NewMessageController()

	// Route pour la connexion WebSocket
	r.GET("/chat", middleware.AuthMiddleware(config.RoleAdmin, config.RoleOrganizer, config.RoleStandLeader), messageController.WebSocketHandler)
}
