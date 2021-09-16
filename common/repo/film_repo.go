package repo

import (
	"micro_backend_film/common/entity"
	"time"

	"gorm.io/gorm"
)

type FilmRepo struct {
	DB *gorm.DB
}

func (f *FilmRepo) FindAll() ([]entity.Film, error) {
	var films []entity.Film
	result := f.DB.Find(&films)
	return films, result.Error
}

func (f *FilmRepo) FindByName(name string) (entity.Film, error) {
	var film entity.Film
	key := "%" + name + "%"
	result := f.DB.Where("name LIKE ?", key).First(&film)
	return film, result.Error
}

func (f *FilmRepo) SaveFilm(film entity.Film) entity.Film {
	film.CreatedAt = time.Now()
	film.UpdatedAt = time.Now()
	f.DB.Save(&film)
	return film
}

func (f *FilmRepo) UpdateFilm(film entity.Film) entity.Film {
	film.UpdatedAt = time.Now()
	f.DB.Save(&film)
	return film
}
