package routes

import (
	"back/config"
	"back/controllers"
	middleware "back/middlewares"
	"back/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB, userRepo repositories.UserRepository) {
	userGroup := r.Group("/user")
	{
		userController := controllers.NewUserController(userRepo)
		//userGroup.Use(middleware.FeatureMiddleware(db, "USER"))
		userGroup.POST("/login", userController.Login)
		userGroup.POST("/register", userController.Register)
		userGroup.GET("/me",
			middleware.AuthMiddleware(config.RoleOrganizer, config.RoleAdmin, config.RoleStudent, config.RoleStandLeader, config.RoleParent),
			userController.GetMe,
		)
		userGroup.PUT("/me",
			middleware.AuthMiddleware(config.RoleOrganizer, config.RoleAdmin, config.RoleStudent, config.RoleStandLeader, config.RoleParent),
			userController.UpdateMe,
		)
		/*userGroup.DELETE("/me",
			middleware.AuthMiddleware(config.RoleOrganizer, config.RoleAdmin, config.RoleStudent, config.RoleStandLeader, config.RoleParent),
			userController.DeleteMe,
		)*/
	}
}
