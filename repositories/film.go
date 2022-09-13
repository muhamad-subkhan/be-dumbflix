package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type FilmRepository interface{
	CreateFilm(film models.Film) (models.Film, error)
	GetFilm(ID int) (models.Film, error)
	FindFilm() ([]models.Film, error)
	UpdateFilm(film models.Film)(models.Film, error)
	DeleteFilm(film models.Film)(models.Film, error)
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

func (r *repository) FindFilm() ([]models.Film, error) {
	var film []models.Film
	err := r.db.Find(&film).Error

	return film, err
}

func (r *repository) UpdateFilm(film models.Film) (models.Film, error) {
	err := r.db.Save(&film).Error

	return film, err
}

func (r *repository) DeleteFilm(film models.Film) (models.Film, error) {
	err := r.db.Delete(&film).Error
	
	return film, err
}