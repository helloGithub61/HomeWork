package models

import "time"

type Email struct {
	ID          int       `json:"id"`
	Create_time time.Time `josn:"create_time"`
	Address     string    `josn:"address"`
	Code        string    `josn:"code"`
	Type        int       `josn:"type"`
}
