package handler

import (
	"context"
	"log"
	"micro_backend_film/common/model"
	"micro_backend_film/config/grpc"
	"micro_backend_film/services/biz/pb/pb_film"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type FilmHandler struct {
}

func (fh *FilmHandler) AllFilms(c *fiber.Ctx) error {
	res, err := grpc.BizGateClient.AllFilms(context.Background(), &pb_film.ReqAllFilms{})
	if err != nil {
		log.Println("Error when calling AllFilms ", err)
	}

	return c.JSON(model.Response{
		Status: http.StatusOK,
		Msg:    "Get all films successfully",
		Data:   res.Films,
	})
}

func (fh *FilmHandler) FindFilmsByUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(string)

	res, err := grpc.BizGateClient.FindFilmsByUser(context.Background(), &pb_film.ReqFindFilmsByUser{UserID: userId})
	if err != nil {
		log.Println("Error when calling AllFilms ", err)
	}

	return c.JSON(model.Response{
		Status: http.StatusOK,
		Msg:    "Get films by user successfully",
		Data:   res.Films,
	})
}
