package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type AuthRepository interface{
	Register(user models.User) (models.User, error)
	Login(Email string) (models.User, error)
	Getuser(ID int) (models.User, error)
}

type repositoryAuth struct {
	db *gorm.DB
}

func RepositoriesAuth(db *gorm.DB)  *repository{
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Login(Email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", Email).Error

	return user, err
}

func (r *repository) Getuser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}
