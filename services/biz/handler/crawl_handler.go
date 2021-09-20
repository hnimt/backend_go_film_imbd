package handler

import (
	"log"
	"micro_backend_film/common/entity"
	"micro_backend_film/common/repo"
	"micro_backend_film/config/cache"
	"micro_backend_film/services/biz/pb/pb_crawl"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type CrawlHandler struct {
	RedisCache *redis.Client
	FilmRepo   *repo.FilmRepo
}

func (ch *CrawlHandler) CrawlFilm(ctx context.Context, req *pb_crawl.CrawledFilms) (*pb_crawl.ResCrawl, error) {
	res := &pb_crawl.ResCrawl{}
	rqFilms := req.Films
	rpFilms, err := ch.FilmRepo.FindAll()
	if err != nil || len(rpFilms) == 0 {
		ch.insertFirst(rqFilms)
		return res, nil
	}

	var cFilms []entity.Film
	cache.GetCache(ch.RedisCache, "films", &cFilms)
	if err != nil || len(cFilms) == 0 {
		if len(rpFilms) > 0 {
			log.Println("Set cached by repo")
			cache.SetCache(ch.RedisCache, "films", rpFilms)
			cache.GetCache(ch.RedisCache, "films", &cFilms)
		}
	}

	ch.updateFilm(cFilms, rqFilms)

	return res, nil
}

func (ch *CrawlHandler) insertFirst(req []*pb_crawl.CrawledFilm) {
	films := []entity.Film{}
	for _, v := range req {
		var film entity.Film
		film.FilmID = uuid.NewString()
		film.Name = v.Name
		film.Year = v.Year
		film.Rating = v.Rating
		savedFilm, err := ch.FilmRepo.SaveFilm(film)
		if err != nil {
			log.Println(err)
		}
		films = append(films, savedFilm)
	}

	cache.SetCache(ch.RedisCache, "films", films)
	log.Println("First Collect")
}

func (ch *CrawlHandler) updateFilm(cfilms []entity.Film, rqFilms []*pb_crawl.CrawledFilm) {
	log.Println("updateFilm cFilms", cfilms)
	// log.Println("updateFilm rqFilms", rqFilms)

	for i := 0; i < len(rqFilms); i++ {
		isInsert := true
		for j := 0; j < len(cfilms); j++ {

			if cfilms[j].Name == rqFilms[i].Name {

				if cfilms[j].Year != rqFilms[i].Year || cfilms[j].Rating != rqFilms[i].Rating {
					cfilms[j].Year = rqFilms[i].Year
					cfilms[j].Rating = rqFilms[i].Rating

					_, err := ch.FilmRepo.UpdateFilm(cfilms[j])
					if err != nil {
						log.Println(err)
					}

					log.Println("Update")
				}
				isInsert = false
				log.Println("Don't do anythings")
				break
			}
		}
		if isInsert {
			var film entity.Film
			film.FilmID = uuid.NewString()
			film.Name = rqFilms[i].Name
			film.Year = rqFilms[i].Year
			film.Rating = rqFilms[i].Rating
			savedFilm, err := ch.FilmRepo.SaveFilm(film)
			if err != nil {
				log.Println(err)
			}
			cfilms = append(cfilms, savedFilm)
			log.Println("Insert")
		}
	}

	cache.SetCache(ch.RedisCache, "films", cfilms)
}
