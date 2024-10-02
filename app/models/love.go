package models

type Love struct {
	ID      int    `json:"id"`
	Account string `json:"account"`
	PostID  int    `json:"post_id"`
	Like    int    `json:"like"`
	Collect int    `json:"collect"`
}
