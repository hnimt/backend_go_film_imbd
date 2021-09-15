package handler

import (
	"context"
	"micro_backend_film/config/cache"
	"micro_backend_film/pkg/entity"
	"micro_backend_film/services/biz/pb/pb_film"
	"micro_backend_film/services/biz/repo"

	"github.com/go-redis/redis"
)

type FilmHandler struct {
	RedisCache *redis.Client
	FilmRepo   *repo.FilmRepo
}

func (fh *FilmHandler) AllFilms(ctx context.Context, req *pb_film.ReqAllFilms) (*pb_film.ResAllFilms, error) {
	var cFilms []entity.Film
	cache.GetCache(fh.RedisCache, "films", &cFilms)

	res := &pb_film.ResAllFilms{}

	for _, rpFilm := range cFilms {
		res.Films = append(res.Films, &pb_film.Film{
			Name:   rpFilm.Name,
			Year:   rpFilm.Year,
			Rating: rpFilm.Rating,
		})
	}

	return res, nil
}
