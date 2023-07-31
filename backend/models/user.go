package models

type User struct {
	ID       uint   `json:"id" gorm:"unique"`
	TeamName string `json:"team_name" gorm:"unique"`
	Password []byte `json:"-"`
	Jwt      string `json:"-"`
	Score    int64  `json:"score"`
}
