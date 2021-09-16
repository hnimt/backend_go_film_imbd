package handler

import (
	"context"
	"micro_backend_film/common/model"
	"micro_backend_film/config/grpc"
	"micro_backend_film/services/auth/pb"

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

	resp, er := grpc.AuthClient.Signup(context.Background(), reqUser)
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
			Data:   resp,
		},
	)
}

func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	reqUser := new(pb.LoginReq)
	if err := c.BodyParser(reqUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			model.Response{
				Status: fiber.StatusInternalServerError,
				Msg:    err.Error(),
				Data:   nil,
			},
		)
	}

	resp, er := grpc.AuthClient.Login(context.Background(), reqUser)
	if er != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			model.Response{
				Status: fiber.StatusUnauthorized,
				Msg:    er.Error(),
				Data:   nil,
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		model.Response{
			Status: fiber.StatusOK,
			Msg:    "Login Successfully",
			Data:   resp,
		},
	)
}
