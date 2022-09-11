package models

// User model struct
type User struct {
  ID          	int			  `json:"id" gorm:"primary_key:auto_increment"`
  Fullname 		  string		`json:"fullname" gorm:"type:varchar(255)"`
  Email		      string 		`json:"email" gorm:"type:varchar(255)"`
  Password 	  	string		`json:"password" gorm:"type:varchar(50)"`
  Gender 	      string	  `json:"gender" gorm:"type:varchar(255)"`
  Phone         string    `json:"phone" gorm:"type:varchar(255)"`
  Address       string    `json:"address" gorm:"type:varchar(255)"`
  Subscribe     string    `json:"subscribe" gorm:"type:varchar(50)"`
}