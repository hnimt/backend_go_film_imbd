package handler

import (
	"context"
	"log"
	"micro_backend_film/common/entity"
	"micro_backend_film/common/repo"
	"micro_backend_film/config/cache"
	"micro_backend_film/services/biz/pb/pb_film"

	"github.com/go-redis/redis"
)

type FilmHandler struct {
	RedisCache *redis.Client
	FilmRepo   *repo.FilmRepo
}

func (fh *FilmHandler) AllFilms(ctx context.Context, req *pb_film.ReqAllFilms) (*pb_film.ResAllFilms, error) {
	var cFilms []entity.Film
	cache.GetCache(fh.RedisCache, "films", &cFilms)

	if cFilms == nil || len(cFilms) == 0 {
		rpFilms, err := fh.FilmRepo.FindAll()
		if err != nil {
			log.Println("AllFilms error allfilms")
		}
		cache.SetCache(fh.RedisCache, "films", rpFilms)
		cache.GetCache(fh.RedisCache, "films", &cFilms)
		log.Println("AllFilms add cache from repo")
	}

	res := &pb_film.ResAllFilms{}

	for _, rpFilm := range cFilms {
		res.Films = append(res.Films, &pb_film.Film{
			ID:     rpFilm.FilmID,
			Name:   rpFilm.Name,
			Year:   rpFilm.Year,
			Rating: rpFilm.Rating,
		})
	}

	return res, nil
}
