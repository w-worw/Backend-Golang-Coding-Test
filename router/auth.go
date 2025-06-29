package router

import (
	handlers "7-solutions/handlers"
	"7-solutions/repositories"
	services "7-solutions/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAuthRouter(
	r *gin.Engine,
	db *mongo.Database,
) {
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)

	authGroup := r.Group("/auth")

	authGroup.POST("/register", authHandler.RegisterUser)
	authGroup.POST("/login", authHandler.AuthenticateUser)
}
