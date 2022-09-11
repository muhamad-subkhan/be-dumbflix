package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type CategoryRepositories interface{
	CreateCategory(category models.Category)(models.Category, error)
	GetCategory(ID int) (models.Category, error)
	Findcategories() ([]models.Category, error)
	UpdateCategory(category models.Category)(models.Category, error)
	DeleteCategory(category models.Category)(models.Category, error)
}

func RepositoryCategory(db *gorm.DB) *repository {
	return &repository{db}
  
  }


  
func (r *repository) CreateCategory(category models.Category) (models.Category, error){
	err := r.db.Create(&category).Error

	return category, err
}


func (r *repository) Findcategories() ([]models.Category, error){
	var categories []models.Category
	err := r.db.Find(&categories).Error

	return categories, err
}


func (r *repository) GetCategory(ID int) (models.Category, error){
	var Category models.Category
	
	err := r.db.First(&Category, ID).Error
	
	return Category, err
}

func (r *repository) UpdateCategory(category models.Category) (models.Category, error){
	err := r.db.Save(&category).Error

	return category, err
}
func (r *repository) DeleteCategory(category models.Category) (models.Category, error){

	
	err := r.db.Delete(&category).Error

	return category, err
}

