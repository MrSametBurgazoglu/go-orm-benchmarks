package sqlxmodels

type Post struct {
	ID       int       `db:"id"`
	UserID   int       `db:"user_id"` // Foreign key to User
	Title    string    `db:"title"`
	Content  string    `db:"content"`
	Comments []Comment `db:"-"`
}
