package filmdto

type FilmRequest struct {
	ID          int    `json:"id"`
	Title       string `json:"title" form:"title" validate:"require"`
	Year        string `json:"year" form:"year" gorm:"type:date"`
	Description string `json:"description" gorm:"type:text"`
}