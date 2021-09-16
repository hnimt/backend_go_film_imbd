package entity

import "time"

type User struct {
	UserID    string    `gorm:"primaryKey" json:"-"`
	Email     string    `gorm:"unique" json:"email,omitempty"`
	Password  string    `gorm:"not null" json:"-"`
	Role      string    `gorm:"role,omitempty" json:"role,omitempty"`
	CreatedAt time.Time `gorm:"omitempty" json:"-"`
	UpdatedAt time.Time `gorm:"omitempty" json:"-"`
	Token     string    `gorm:"-" json:"token,omitempty"`
}
