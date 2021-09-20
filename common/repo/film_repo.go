package repo

import (
	"log"
	"micro_backend_film/common/entity"
	"micro_backend_film/common/exception"
	"time"

	"github.com/jmoiron/sqlx"
)

type FilmRepo struct {
	DB *sqlx.DB
}

func (f *FilmRepo) FindAll() ([]entity.Film, error) {
	var films []entity.Film
	err := f.DB.Select(&films, "SELECT * FROM films")
	if err != nil {
		log.Println(err)
		return films, exception.FilmNotFound
	}

	return films, nil
}

func (f *FilmRepo) FindBookmarkedFilms(userID string) ([]entity.Film, error) {
	var films []entity.Film
	sql := `
		SELECT * 
		FROM films
		WHERE film_id IN (SELECT film_id FROM bookmarks WHERE user_id = ?)
	`
	if err := f.DB.Select(&films, sql, userID); err != nil {
		log.Println(err)
		return films, exception.FilmNotFound
	}

	return films, nil
}

func (f *FilmRepo) FindNoBookmarkedFilms(userID string) ([]entity.Film, error) {
	var films []entity.Film
	sql := `
		SELECT * 
		FROM films
		WHERE film_id NOT IN (SELECT film_id FROM bookmarks WHERE user_id = ?)
	`
	if err := f.DB.Select(&films, sql, userID); err != nil {
		log.Println(err)
		return films, exception.FilmNotFound
	}

	return films, nil
}

func (f *FilmRepo) FindByName(name string) (entity.Film, error) {
	var film entity.Film
	err := f.DB.Get(&film, "SELECT * FROM fils WHERE name LIKE %?%", name)
	if err != nil {
		log.Println(err)
		return film, exception.FilmNotFound
	}

	return film, nil

}

func (f *FilmRepo) SaveFilm(film entity.Film) (entity.Film, error) {
	film.CreatedAt = time.Now()
	film.UpdatedAt = time.Now()

	sql := `
		INSERT INTO films(film_id, name, year, rating, image, created_at, updated_at) 
		VALUES(:film_id, :name, :year, :rating, :image, :created_at, :updated_at)
	`

	_, err := f.DB.NamedExec(sql, film)
	if err != nil {
		log.Println(err)
		return film, exception.FilmNotCreated
	}

	return film, nil
}

func (f *FilmRepo) UpdateFilm(film entity.Film) (entity.Film, error) {
	film.UpdatedAt = time.Now()
	sql := `
		UPDATE films
		SET name = :name, 
			email = :email,
			year = :year,
			rating = :rating,
			image = :image,
			updated_at = :updated_at
		WHERE film_id = :film_id
	`

	_, err := f.DB.NamedExec(sql, film)
	if err != nil {
		log.Println(err)
		return film, exception.FilmNotUpdated
	}

	return film, nil
}
