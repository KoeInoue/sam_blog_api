package repository

import "blog/model"

type PostRepositoryInterface interface {
	GetOne(string) (model.Post, error)
	GetAll() ([]model.Post, error)
	Post(*model.Post) (bool, error)
	Edit(*model.Post) (bool, error)
	Delete(id string) (bool, error)
}
