package sqlxmodels

type Comment struct {
	ID     int    `db:"id"`
	PostID int    `db:"post_id"` // Foreign key to Post
	Text   string `db:"text"`
}
