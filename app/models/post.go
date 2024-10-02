package models

import "time"

type Post struct {
	ID          int       `json:"id"`
	CreateTime  time.Time `json:"create_time"`
	Account     string    `json:"user_id"`
	PostContent string    `json:"post_content"`
	ShowType    int       `json:"show_type"`
}
