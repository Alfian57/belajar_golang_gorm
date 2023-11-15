package repository

import "github.com/Alfian57/belajar_golang_gorm/model"

type UserRepository interface {
	Create(name string) model.User
	Delete(id int)
	FindAll() []model.User
}
