package main

import (
	"micro_backend_film/common/middleware"
	"micro_backend_film/services/gate/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Handler
	filmHandler := &handler.FilmHandler{}
	authHandler := &handler.AuthHandler{}
	bmHandler := &handler.BMHandler{}

	// Router
	films := app.Group("/films")
	films.Get("/", filmHandler.AllFilms)

	auth := app.Group("/auth")
	auth.Post("/signup", authHandler.Signup)
	auth.Post("/login", authHandler.Login)

	bm := app.Group("/bookmark", middleware.JwtMiddleware())
	bm.Get("/", bmHandler.HandleFindByUser)
	bm.Post("/add", bmHandler.HandleAddBM)
	bm.Post("/del", bmHandler.HandleDelBM)

	app.Listen(":8080")
}
