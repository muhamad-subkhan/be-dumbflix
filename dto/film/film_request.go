package filmdto

type FilmResponse struct {
	ID           int    `json:"id" `
	Title        string `json:"title" `
	TumbnailFilm string `json:"tumbnailfilm" `
	Year         string `json:"year"`
	Description  string `json:"deskripsi"`
	UserID       int    `json:"user_id"`
}
type UpdateFilm struct {
	ID           int    `json:"id" `
	Title        string `json:"title" `
	TumbnailFilm string `json:"tumbnailfilm" `
	Year         string `json:"year"`
	Description  string `json:"deskripsi"`
}