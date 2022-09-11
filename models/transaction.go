package models

type Transaction struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment"`
	Startdate string `json:"startdate" gorm:"type:date"`
	Duedate   string `json:"dueDate" gorm:"type:date"`
	// UserId			[]User	`json`
	Attache string `json:"atttach" gorm:"type: varchar(255)"`
	Status  string `json:"status" gorm:"type:text" form:"status"`
}