package handler

import (
	"context"
	"micro_backend_film/common/model"
	"micro_backend_film/common/model/req"
	"micro_backend_film/config/grpc"
	"micro_backend_film/services/bookmark/pb"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type BMHandler struct{}

func (bm *BMHandler) HandleAddBM(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(string)

	reqAddBM := new(req.ReqBM)

	if err := c.BodyParser(reqAddBM); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: fiber.StatusBadRequest,
			Msg:    err.Error(),
			Data:   nil,
		})
	}

	reqAdd := &pb.ReqAddBookmark{
		UserID: userId,
		FilmID: reqAddBM.FilmID,
	}

	res, err := grpc.BMClient.AddBookmark(context.Background(), reqAdd)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: fiber.StatusBadRequest,
			Msg:    err.Error(),
			Data:   nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status: fiber.StatusOK,
		Msg:    "Add book mark successfully",
		Data:   res.Film,
	})

}

func (bm *BMHandler) HandleDelBM(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(string)

	reqDelBM := new(req.ReqBM)

	if err := c.BodyParser(reqDelBM); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: fiber.StatusBadRequest,
			Msg:    err.Error(),
			Data:   nil,
		})
	}

	reqDel := &pb.ReqDelBookmark{
		UserID: userId,
		FilmID: reqDelBM.FilmID,
	}

	res, err := grpc.BMClient.DelBookmark(context.Background(), reqDel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: fiber.StatusBadRequest,
			Msg:    err.Error(),
			Data:   nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status: fiber.StatusOK,
		Msg:    "Delete bookmark successfully",
		Data:   res,
	})
}

func (bm *BMHandler) HandleFindByUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"].(string)

	reqFind := &pb.ReqFindByUser{
		UserID: userId,
	}

	res, err := grpc.BMClient.FindByUser(context.Background(), reqFind)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: fiber.StatusBadRequest,
			Msg:    err.Error(),
			Data:   nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status: fiber.StatusOK,
		Msg:    "Find successfully",
		Data:   res,
	})
}
