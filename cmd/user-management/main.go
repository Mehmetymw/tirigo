package main

import (
	"log"
	"tirigo/internal/user-management/config"
	"tirigo/internal/user-management/domain"
	"tirigo/internal/user-management/handler"
	"tirigo/internal/user-management/repository"
	"tirigo/internal/user-management/service"
	mainconfig "tirigo/pkg/config"
	"tirigo/pkg/jwt"
	"tirigo/pkg/redis"

	"github.com/gofiber/fiber"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("couldn't load config: %v", err)
	}

	redis.Init(cfg.RedisAddr)
	jwt.Init(cfg.JWTSecret)

	db := mainconfig.InitDatabase()
	db.AutoMigrate(&domain.User{})

	userRepo := repository.NewGormUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()
	api := app.Group("/api")
	userRoutes := api.Group("/user")

	userHandler.InitializeRoutes(userRoutes)

	log.Println("Starting server on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
