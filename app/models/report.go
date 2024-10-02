package models

type Report struct {
	ID      int    `json:"id"`
	PostID  int    `json:"post_id"`
	Content string `json:"Content"`
	Account string `json:"user_id"`
	Reason  string `json:"reason"`
	Status  int    `json:"status"`
}
