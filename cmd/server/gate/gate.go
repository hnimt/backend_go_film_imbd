package main

import (
	"micro_backend_film/services/gate/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Handler
	filmHandler := &handler.FilmHandler{}
	authHandler := &handler.AuthHandler{}

	// Router
	films := app.Group("/films")
	films.Get("/", filmHandler.AllFilms)

	auth := app.Group("/auth")
	auth.Post("/signup", authHandler.Signup)
	auth.Post("/login", authHandler.Login)

	app.Listen(":8080")
}
