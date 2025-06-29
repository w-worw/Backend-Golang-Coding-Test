package router

import (
	handlers "7-solutions/handlers"
	"7-solutions/middleware"
	"7-solutions/repositories"
	services "7-solutions/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddUserRouter(r *gin.Engine, db *mongo.Database) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	userGroup := r.Group("/users")

	userGroup.Use(middleware.AuthenticationMiddleware())

	userGroup.POST("/", userHandler.CreateUser)
	userGroup.GET("/:id", userHandler.GetUserByID)
	userGroup.GET("/", userHandler.GetAllUsers)
	userGroup.PUT("/:id", userHandler.UpdateUser)
	userGroup.DELETE("/:id", userHandler.DeleteUser)
}
