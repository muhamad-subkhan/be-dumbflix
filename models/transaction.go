package models

import "time"

type Transaction struct {
	ID        int    		`json:"id" gorm:"primary_key:auto_increment"`
	StartDate time.Time    	`json:"startdate" form:"stardate"`
	DueDate   time.Time   	`json:"duedate" form:"duedate"`
	Attache   string 		`json:"attache" form:"attache"`
	Status    string 		`json:"status" form:"status"`
}