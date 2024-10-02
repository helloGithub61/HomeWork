package models

import "time"

type Comment struct {
	ID         int       `json:"id"`
	CreateTime time.Time `json:"create_time"`
	Account    string    `json:"account"`
	ComContent string    `json:"com_content"`
	Target     int       `json:"target"`
	PostID     int       `json:"post_id"`
}
