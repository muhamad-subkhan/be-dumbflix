package transactionsdto

import "time"

type CreateTransactionRequest struct {
	ID        string `json:"id"`
	StartDate time.Time    	`json:"startdate" form:"startdate"`
	DueDate   time.Time    	`json:"duedate" form:"duedate"`
	Attache   string		`json:"attache"`
	Status    string		`json:"status"`
}

type UpdateTransactionRequest struct {
	ID        string		`json:"id"`
	StartDate time.Time		`json:"statusdate"`
	DueDate   time.Time		`json:"duedate"`
	Attache   string 		`json:"attache"`
	Status    string		`json:"status"`
}
