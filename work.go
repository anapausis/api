package domain

type Work struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	UserID int    `json:"user_id"`
}
