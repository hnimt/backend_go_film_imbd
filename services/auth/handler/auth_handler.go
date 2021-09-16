package handler

import (
	"context"
	"micro_backend_film/config/cache"
	"micro_backend_film/pkg/entity"
	"micro_backend_film/pkg/security"
	"micro_backend_film/services/auth/pb"
	"micro_backend_film/services/auth/repo"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthHandler struct {
	RdCache  *redis.Client
	UserRepo *repo.UserRepo
}

func (ah *AuthHandler) Signup(ctx context.Context, req *pb.SignupReq) (*pb.AuthRes, error) {

	isDuplicateEmail, _ := ah.UserRepo.IsDuplicateEmail(req.Email)
	if !isDuplicateEmail {
		return nil, &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "Duplicated email",
		}
	}

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

	ah.UserRepo.SaveUser(user)

	cache.SetCache(ah.RdCache, "user", user)

	return &pb.AuthRes{
		Email: req.Email,
		Role:  role,
	}, nil

}

func (ah *AuthHandler) Login(ctx context.Context, req *pb.LoginReq) (*pb.AuthRes, error) {

	user, err := ah.UserRepo.CheckLogin(req.Email)
	if err != nil {
		return nil, &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "Wrong email",
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

	result := &pb.AuthRes{
		Email: user.Email,
		Role:  user.Role,
	}

	return result, nil
}
