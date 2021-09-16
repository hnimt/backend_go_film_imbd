package entity

type Bookmark struct {
	BID    string `json:"bmID" gorm:"primaryKey"`
	UserID string `json:"-" gorm:"omitempty"`
	FilmID string `json:"-" gorm:"omitempty"`

	User User `json:"user,omitempty" gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE;onDelete:CASCADE"`
	Film Film `json:"film,omitempty" gorm:"foreignkey:FilmID;constraint:onUpdate:CASCADE;onDelete:CASCADE"`
}
