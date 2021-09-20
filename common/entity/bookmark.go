package entity

type Bookmark struct {
	BID    string `db:"b_id" json:"id"`
	UserID string `db:"user_id" json:"userID"`
	FilmID string `db:"film_id" json:"filmID"`
}
