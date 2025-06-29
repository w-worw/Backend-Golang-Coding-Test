package main

import (
	"7-solutions/database"
	"7-solutions/repositories"
	"7-solutions/router"
	"7-solutions/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func task(db *mongo.Database) {
	userRepository := repositories.NewUserRepository(db)
	count, err := userRepository.CountUsers()
	if err != nil {
		log.Printf("Error counting users: %v", err)
		return
	}
	log.Printf("Number of users: %d", count)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	r := gin.Default()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	db := database.NewMongoDB(dbHost, dbPort, dbName)

	if err := utils.EnsureEmailUniqueIndex(db); err != nil {
		log.Fatalf("Error ensuring email unique index: %v", err)
	}

	router.AddAuthRouter(r, db)
	router.AddUserRouter(r, db)

	s := gocron.NewScheduler()
	s.Every(10).Seconds().Do(task, db)
	go func() {
		<-s.Start()
	}()

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
