package routes

import (
	"back/config"
	"back/controllers"
	middleware "back/middlewares"
	"back/repositories"
	"github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.Engine, studentRepo repositories.StudentRepository) {
	studentController := controllers.NewStudentController(studentRepo)
	studentRoutes := router.Group("/students")
	{
		studentRoutes.POST("/", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent), studentController.CreateStudent)
		studentRoutes.GET("/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent), studentController.GetStudentByID)
		studentRoutes.GET("/parent/:parent_id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent), studentController.GetStudentsByParentID)
		studentRoutes.PUT("/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent, config.RoleStudent), studentController.UpdateStudent)
		studentRoutes.DELETE("/:id", middleware.AuthMiddleware(config.RoleAdmin, config.RoleParent), studentController.DeleteStudent)
	}
}
