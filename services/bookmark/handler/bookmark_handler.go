package handler

import (
	"context"
	"micro_backend_film/common/entity"
	"micro_backend_film/common/repo"
	"micro_backend_film/config/cache"
	"micro_backend_film/services/bookmark/pb"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type BookmarkHandler struct {
	RedisCache *redis.Client
	BMRepo     *repo.BookmarkRepo
}

func (bh *BookmarkHandler) AddBookmark(ctx context.Context, req *pb.ReqAddBookmark) (*pb.ResAddBookmark, error) {
	bookmark := entity.Bookmark{
		BID:    uuid.NewString(),
		UserID: req.UserID,
		FilmID: req.FilmID,
	}

	_, err := bh.BMRepo.AddBM(bookmark)
	if err != nil {
		return nil, err
	}

	var cFilms []entity.Film
	cache.GetCache(bh.RedisCache, req.UserID, &cFilms)
	for i := 0; i < len(cFilms); i++ {
		if cFilms[i].FilmID == req.FilmID {
			cFilms[i].Bookmarked = "true"
			break
		}
	}
	cache.SetCache(bh.RedisCache, req.UserID, cFilms)

	result := &pb.ResAddBookmark{
		Res: "Created successfully",
	}

	return result, nil
}

func (bh *BookmarkHandler) DelBookmark(ctx context.Context, req *pb.ReqDelBookmark) (*pb.ResDelBookmark, error) {
	err := bh.BMRepo.RemoveByUserAndFilm(req.UserID, req.FilmID)
	if err != nil {
		return nil, err
	}

	var cFilms []entity.Film
	cache.GetCache(bh.RedisCache, req.UserID, &cFilms)
	for i := 0; i < len(cFilms); i++ {
		if cFilms[i].FilmID == req.FilmID {
			cFilms[i].Bookmarked = "false"
			break
		}
	}
	cache.SetCache(bh.RedisCache, req.UserID, cFilms)

	return &pb.ResDelBookmark{}, nil
}
