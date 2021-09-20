package exception

import "errors"

var (
	UserNotCreated = errors.New("Cannot create user")
	UserNotFound   = errors.New("Cannot found user")
	UserNotUpdated = errors.New("Cannot update user")

	FilmNotFound   = errors.New("Film not found")
	FilmNotCreated = errors.New("Cannot create film")
	FilmNotUpdated = errors.New("Cannot update film")

	BMNotCreated = errors.New("Cannot create bookmark")
	BMNotDeleted = errors.New("Cannot delete bookmark")
)
