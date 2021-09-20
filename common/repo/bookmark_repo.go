package repo

import (
	"log"
	"micro_backend_film/common/entity"
	"micro_backend_film/common/exception"

	"github.com/jmoiron/sqlx"
)

type BookmarkRepo struct {
	DB *sqlx.DB
}

func (b *BookmarkRepo) AddBM(bookmark entity.Bookmark) (entity.Bookmark, error) {
	sql := `
		INSERT INTO bookmarks(b_id, user_id, film_id)
		VALUES(:b_id, :user_id, :film_id)
	`

	_, err := b.DB.NamedExec(sql, bookmark)
	if err != nil {
		log.Println(err)
		return bookmark, exception.BMNotCreated
	}

	return bookmark, nil
}

func (b *BookmarkRepo) RemoveByUserAndFilm(userID string, filmID string) error {
	_, err := b.DB.Exec("DELETE FROM bookmarks WHERE user_id = ? AND film_id = ?", userID, filmID)
	if err != nil {
		log.Println(err)
		return exception.BMNotDeleted
	}

	return nil
}

