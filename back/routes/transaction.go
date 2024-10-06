package routes

import (
	"back/config"
	"back/controllers"
	middleware "back/middlewares"
	"back/repositories"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.Engine, transactionRepo repositories.TransactionRepository) {
	transactionController := controllers.NewTransactionController(transactionRepo)

	r.POST("/transactions", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent), transactionController.CreateTransaction)
	r.GET("/transactions/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent), transactionController.GetTransaction)
	r.PUT("/transactions/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent), transactionController.UpdateTransaction)
	r.DELETE("/transactions/:id", middleware.AuthMiddleware(config.RoleAdmin), transactionController.DeleteTransaction)
	r.GET("/transactions", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent), transactionController.GetTransactionsByParentID)
	r.POST("/api/create-payment-intent", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent), controllers.CreatePaymentIntent)
	//r.POST("/api/complete-purchase", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent), controllers.CompletePurchase)
}
