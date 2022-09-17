package models

type Film struct {
	ID           int    `json:"id" gorm:"primary_key:auto_increment"`
	Title        string `json:"title" gorm:"type: varchar(255)"`
	TumbnailFilm string `json:"tumbnailfilm" gorm:"type: varchar(255)"`
	Year         string `json:"year"`
	Description  string `json:"deskripsi" gorm:"type:text" form:"deskripsi"`
	UserID       int    `json:"-" gorm:"type:int"`
}
