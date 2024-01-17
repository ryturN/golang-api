package service

import (
	"github.com/golang-api/dto"
	"github.com/golang-api/entity"
	"github.com/golang-api/errorhandler"
	"github.com/golang-api/repository"
)

type PostService interface {
	Create(req *dto.PostRequest) error
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		repository: r,
	}
}

func (s *postService) Create(req *dto.PostRequest) error {
	post := entity.Post{
		UsersId: req.UsersId,
		Post:    req.Post,
	}
	if req.Picture != nil {
		post.PictureUrl = &req.Picture.Filename
	}
	if err := s.repository.Create(&post); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}
	return nil
}
