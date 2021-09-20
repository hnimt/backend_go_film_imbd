package entity

import "time"

type Film struct {
	FilmID     string    `db:"film_id" json:"id"`
	Name       string    `db:"name,omitempty" json:"name"`
	Year       string    `db:"year,omitempty" json:"year"`
	Rating     float64   `db:"rating,omitempty" json:"rating"`
	Image      string    `db:"image,omitempty" json:"image"`
	CreatedAt  time.Time `db:"created_at,omitempty" json:"-"`
	UpdatedAt  time.Time `db:"updated_at,omitempty" json:"-"`
	Bookmarked string    `json:"bookmarked"`
}
