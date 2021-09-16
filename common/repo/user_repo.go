package repo

import (
	"micro_backend_film/common/entity"
	"time"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) SaveUser(user entity.User) (entity.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	result := u.DB.Save(&user)
	return user, result.Error
}

func (u *UserRepo) SelectUserById(userId string) (entity.User, error) {
	var user entity.User
	result := u.DB.Find(&user, userId)
	return user, result.Error
}

func (u *UserRepo) UpdateUser(user entity.User) (entity.User, error) {
	user.UpdatedAt = time.Now()
	result := u.DB.Save(&user)
	return user, result.Error
}

func (u *UserRepo) CheckLogin(email string) (entity.User, error) {
	var user entity.User
	result := u.DB.Where("email = ?", email).Take(&user)
	return user, result.Error
}

func (u *UserRepo) IsDuplicateEmail(email string) (bool, error) {
	var user entity.User
	result := u.DB.Where("email = ?", email).Take(&user)
	return !(result.Error == nil), result.Error
}
