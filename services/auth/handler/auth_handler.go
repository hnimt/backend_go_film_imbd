package handler

import (
	"context"
	"errors"
	"log"
	"micro_backend_film/pkg/entity"
	"micro_backend_film/pkg/security"
	"micro_backend_film/services/auth/pb"
	"micro_backend_film/services/auth/repo"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type AuthHandler struct {
	RdCache  *redis.Client
	UserRepo *repo.UserRepo
}

func (ah *AuthHandler) Signup(ctx context.Context, req *pb.SignupReq) (*pb.AuthRes, error) {

	isDuplicateEmail, _ := ah.UserRepo.IsDuplicateEmail(req.Email)
	if !isDuplicateEmail {
		return nil, errors.New("duplicated email")
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

	// Gen token
	token, err := security.GenToken(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.AuthRes{
		Email: req.Email,
		Role:  role,
		Token: token,
	}, nil

}

func (ah *AuthHandler) Login(ctx context.Context, req *pb.LoginReq) (*pb.AuthRes, error) {
	return nil, nil
}
