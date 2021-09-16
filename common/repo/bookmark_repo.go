package repo

import (
	"micro_backend_film/common/entity"

	"gorm.io/gorm"
)

type BookmarkRepo struct {
	DB *gorm.DB
}

func (b *BookmarkRepo) AddBM(bookmark entity.Bookmark) (entity.Bookmark, error) {
	result := b.DB.Save(&bookmark)
	b.DB.Preload("User").Preload("Film").Find(&bookmark)
	return bookmark, result.Error
}

func (b *BookmarkRepo) RemoveByUserAndFilm(userID string, filmID string) error {
	result := b.DB.Where("user_id = ? AND film_id = ?", userID, filmID).Delete(&entity.Bookmark{})
	return result.Error
}

func (b *BookmarkRepo) FindBMByUser(userID string) ([]entity.Bookmark, error) {
	bookmarks := []entity.Bookmark{}
	result := b.DB.Where("user_id = ?", userID).Preload("Film").Find(&bookmarks)
	return bookmarks, result.Error
}
