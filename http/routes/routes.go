package routes

import (
	"database/sql"
	"fmt"

	"github.com/ahsansaif47/cdc-app/http/controllers"
	"github.com/ahsansaif47/cdc-app/http/handlers"
	"github.com/ahsansaif47/cdc-app/repository/postgres"
	"github.com/ahsansaif47/cdc-app/repository/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title						CDC-APP Local API
// @version					1.0
// @description				This is a swagger for CDC-APP
// @host						localhost:8081
// @BasePath					/api/v1
// @schemes					http
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Meals-Service
func InitRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	userRoutes := v1.Group("/users")
	fmt.Println(userRoutes)
	// InitUserRoutes(userRoutes)
}

func InitUserRoutes(userRoutes fiber.Router, db *sql.DB, cache redis.ICacheRepository) {
	userRepo := postgres.NewUserRepository(db)
	userService := controllers.NewUserService(userRepo, cache)
	userHandlers := handlers.NewAuthHandler(userService)
	// fmt.Println(mealsHandlers)

	userRoutes.Post("/create", userHandlers.CreateUser)
	userRoutes.Post("/login", userHandlers.CreateUser)
	userRoutes.Get("/otp", userHandlers.GenerateOTP)
	userRoutes.Post("/otp", userHandlers.VerifyOTP)

	// TODO - CRUD API HERE
	// userRoutes.Put("/update/:id", userHandlers.Update)
	// userRoutes.Delete("/delete/:id", userHandlers.DeleteItem)
	// userRoutes.Get("/all-meals", userHandlers.GetUser)
	// userRoutes.Get("/meals/:id", userHandlers.GetItemByID)
	// userRoutes.Get("/filter", userHandlers.FilterItems)
	// userRoutes.Get("/search", userHandlers.SearchItems)
}
