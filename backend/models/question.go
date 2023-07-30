package models

type Question struct {
	ID       uint   `json:"id" gorm:"unique"`
	Question string `json:"question"`
	A        string `json:"a"`
	B        string `json:"b"`
	C        string `json:"c"`
	D        string `json:"d"`
	Answer   string `json:"answer"`
	Reward   int64  `json:"reward"`
}
