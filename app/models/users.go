package models

type User struct {
	ID        int     `json:"id"`
	Account   string  `json:"account"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	UserType  int     `json:"user_type" gorm:"default:0"`
	Image     string  `json:"image" gorm:"default:./asset/pic/df.png"`
	FontSize  float32 `json:"font_size" gorm:"default:12"`
	FontColor string  `json:"font_color" gorm:"default:black"`
	Status    int     `json:"status" gorm:"default:0"`
	TKey      string  `json:"tkey"`
}
