package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type FilmRepository interface{
	CreateFilm(film models.Film) (models.Film, error)
	GetFilm(ID int) (models.Film, error)
}

func RepositoryFilm(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateFilm(Film models.Film) (models.Film, error){
	err := r.db.Create(&Film).Error

	return Film, err
}

func (r *repository) GetFilm(ID int) (models.Film, error) {
	var Film models.Film
	
	err := r.db.First(&Film, ID).Error

	return Film, err
}