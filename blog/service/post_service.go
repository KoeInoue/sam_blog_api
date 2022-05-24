// service for business logic
package service

import (
	"blog/model"
	"blog/repository"
	"encoding/json"
)

// PostService
type PostService struct {
	repo repository.PostRepositoryInterface
}

// NewPostService returns PostService struct
func NewPostService() *PostService {
	ps := new(PostService)
	ps.repo = repository.NewPostRepository()

	return ps
}

// GetOne returns a post struct
func (ps *PostService) GetOne(id string) (model.Post, error) {
	return ps.repo.GetOne(id)
}

// GetAll returns array of post struct
func (ps *PostService) GetAll() ([]model.Post, error) {
	return ps.repo.GetAll()
}

// Post returns bool or error, Store post to DB
func (ps *PostService) Post(bodyString string) (bool, error) {
	p := model.NewPost()

	if err := json.Unmarshal([]byte(bodyString), p); err != nil {
		return false, err
	}

	return ps.repo.Post(p)
}

// Edit returns bool or error, Update post in DB
func (ps *PostService) Edit(bodyString string) (bool, error) {
	p := model.NewPost()

	if err := json.Unmarshal([]byte(bodyString), p); err != nil {
		return false, err
	}

	return ps.repo.Edit(p)
}

// Delete returns bool or error, Delete post in DB
func (ps *PostService) Delete(id string) (bool, error) {
	return ps.repo.Delete(id)
}
