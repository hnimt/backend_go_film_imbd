package entity

import "time"

type User struct {
	UserID    string    `db:"user_id" json:"id"`
	Email     string    `db:"email" json:"email,omitempty"`
	Password  string    `db:"password" json:"-"`
	Role      string    `db:"role,omitempty" json:"role,omitempty"`
	CreatedAt time.Time `db:"created_at,omitempty" json:"-"`
	UpdatedAt time.Time `db:"updated_at,omitempty" json:"-"`
	Token     string    `json:"token,omitempty"`
}
