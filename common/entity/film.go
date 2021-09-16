package entity

import "time"

type Film struct {
	FilmID    string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"omitempty"`
	Year      string    `json:"year" gorm:"omitempty"`
	Rating    float64   `json:"rating" gorm:"omitempty"`
	CreatedAt time.Time `gorm:"omitempty" json:"-"`
	UpdatedAt time.Time `gorm:"omitempty" json:"-"`
}
