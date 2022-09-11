package models

type Episode struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type:varchar(255)" `
	Year          int    `json:"year" `
	// Category      []Category `json:"category" gorm:"many2many:episode_categories; constraint:OnUpdates:CASCADE,OnDelete:CASCADE;"`
	// CategoryId    []int      `json:"-" form:"category_id" gorm:"-"`
}