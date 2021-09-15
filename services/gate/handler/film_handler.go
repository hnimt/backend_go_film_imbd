package handler

import (
	"context"
	"log"
	"micro_backend_film/pkg/model"
	"micro_backend_film/services/biz/pb/pb_film"
	"micro_backend_film/services/gate"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FilmHandler struct {
}

func (fh *FilmHandler) AllFilms(c *fiber.Ctx) error {
	res, err := gate.BizClient.AllFilms(context.Background(), &pb_film.ReqAllFilms{})
	if err != nil {
		log.Fatal("Error when calling AllFilms ", err)
	}

	return c.JSON(model.Response{
		Status: http.StatusOK,
		Msg:    "Get all films successfully",
		Data:   res.Films,
	})

}
