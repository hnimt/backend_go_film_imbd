package repo

import (
	"log"
	"micro_backend_film/common/entity"
	"micro_backend_film/common/exception"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	DB *sqlx.DB
}

func (u *UserRepo) SaveUser(user entity.User) (entity.User, error) {

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	sql := `
		INSERT INTO users(user_id, email, password, role, created_at, updated_at) VALUES(:user_id, :email, :password, :role, :created_at, :updated_at)
	`

	_, err := u.DB.NamedExec(sql, user)
	if err != nil {
		log.Println(err)
		return user, exception.UserNotCreated
	}
	return user, nil
}

func (u *UserRepo) SelectUserById(userId string) (entity.User, error) {
	var user entity.User

	err := u.DB.Get(&user, "SELECT * FROM users WHERE user_id = ?", userId)
	if err != nil {
		log.Println(err)
		return user, exception.UserNotFound
	}

	return user, nil
}

func (u *UserRepo) UpdateUser(user entity.User) (entity.User, error) {
	user.UpdatedAt = time.Now()

	sql := `
		UPDATE users
		SET
			email = (CASE WHEN LENGTH(:email) = 0 THEN email ELSE :email END),
			updated_at = :updated_at
		WHERE user_id = :user_id
	`

	_, err := u.DB.NamedExec(sql, user)
	if err != nil {
		log.Println(err)
		return user, exception.UserNotUpdated
	}

	return user, nil
}

func (u *UserRepo) CheckLogin(email string) (entity.User, error) {
	var user entity.User

	log.Println(email)
	err := u.DB.Get(&user, "SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		log.Println(err)
		return user, exception.UserNotFound
	}

	return user, nil
}
