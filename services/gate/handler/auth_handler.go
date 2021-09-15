package handler

import (
	"context"
	"micro_backend_film/pkg/model"
	"micro_backend_film/services/auth/pb"
	"micro_backend_film/services/gate"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct{}

func (ah *AuthHandler) Signup(c *fiber.Ctx) error {

	reqUser := new(pb.SignupReq)

	if err := c.BodyParser(reqUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			model.Response{
				Status: fiber.StatusInternalServerError,
				Msg:    err.Error(),
				Data:   nil,
			},
		)
	}

	res, er := gate.AuthClient.Signup(context.Background(), reqUser)
	if er != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			model.Response{
				Status: fiber.StatusInternalServerError,
				Msg:    er.Error(),
				Data:   nil,
			},
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		model.Response{
			Status: fiber.StatusCreated,
			Msg:    "Create Successfully",
			Data:   res,
		},
	)

}
