package repository

import (
	"github.com/Alfian57/belajar_golang_gorm/helper"
	"github.com/Alfian57/belajar_golang_gorm/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (repository *UserRepositoryImpl) Create(name string) model.User {
	user := model.User{
		Name: name,
	}

	err := repository.db.Create(&user).Error
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(id int) {
	err := repository.db.Delete(&model.User{}, id).Error
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindAll() []model.User {
	var users []model.User

	err := repository.db.Find(&users).Error
	helper.PanicIfError(err)

	return users
}
