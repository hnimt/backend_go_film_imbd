package main

import (
	"micro_backend_film/common/middleware"
	"micro_backend_film/config"
	"micro_backend_film/services/gate/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Config
	config := config.Config()

	app := fiber.New()

	// Handler
	filmHandler := &handler.FilmHandler{}
	authHandler := &handler.AuthHandler{}
	bmHandler := &handler.BMHandler{}

	// Router
	auth := app.Group("/auth")
	auth.Post("/signup", authHandler.Signup)
	auth.Post("/login", authHandler.Login)

	films := app.Group("/films", middleware.JwtMiddleware())
	films.Get("/", filmHandler.FindFilmsByUser)

	bm := app.Group("/bookmark", middleware.JwtMiddleware())
	bm.Post("/add", bmHandler.HandleAddBM)
	bm.Post("/del", bmHandler.HandleDelBM)

	app.Listen(config.Servers["gate"].Port)
}
