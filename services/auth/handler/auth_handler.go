package handler

import (
	"context"
	"micro_backend_film/common/entity"
	"micro_backend_film/common/repo"
	"micro_backend_film/common/security"
	"micro_backend_film/config/cache"
	"micro_backend_film/services/auth/pb"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthHandler struct {
	RdCache  *redis.Client
	UserRepo *repo.UserRepo
}

func (ah *AuthHandler) Signup(ctx context.Context, req *pb.SignupReq) (*pb.AuthRes, error) {

	hashPwd := security.HashAndSalt([]byte(req.Password))
	role := "MEMBER"
	userId := uuid.NewString()

	user := entity.User{
		UserID:   userId,
		Email:    req.Email,
		Password: hashPwd,
		Role:     role,
		Token:    "",
	}

	savedUser, err := ah.UserRepo.SaveUser(user)
	if err != nil {
		return nil, &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: err.Error(),
		}
	}

	cache.SetCache(ah.RdCache, "user", savedUser)

	token, _ := security.GenToken(savedUser.UserID, savedUser.Role)

	return &pb.AuthRes{
		Email: req.Email,
		Role:  role,
		Token: token,
	}, nil

}

func (ah *AuthHandler) Login(ctx context.Context, req *pb.LoginReq) (*pb.AuthRes, error) {

	user, err := ah.UserRepo.CheckLogin(req.Email)
	if err != nil {
		return nil, &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: err.Error(),
		}
	}

	isTruePass := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTruePass {
		return nil, &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "Wrong password",
		}
	}

	cache.SetCache(ah.RdCache, "user", user)

	token, _ := security.GenToken(user.UserID, user.Role)

	return &pb.AuthRes{
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	}, nil
}
