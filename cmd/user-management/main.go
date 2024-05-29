package main

import (
	"log"
	"tirigo/internal/user-management/domain"
	"tirigo/internal/user-management/handler"
	"tirigo/internal/user-management/repository"
	"tirigo/internal/user-management/service"
	"tirigo/pkg/config"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := config.Database()
	db.AutoMigrate(&domain.User{})

	app := fiber.New()

	userRepo := repository.NewGormUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	userHandler.InitializeRoutes(app)

	log.Println("Starting server on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
