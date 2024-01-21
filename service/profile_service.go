package service

import (
	"github.com/golang-api/dto"
	"github.com/golang-api/repository"
)

type ProfileService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	UpdateProfil()
}

type profileService struct {
	repository repository.AuthRepository
}

func NewProfileService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r}
}

// func (s *postService) UpdateProfile(req *dto.ProfileRequest) {
// 	post := en{
// 		UsersId: req.UsersId,
// 		Post:    req.Post,
// 	}
// 	if req.Picture != nil {
// 		post.PictureUrl = &req.Picture.Filename
// 	}
// 	if err := s.repository.Update(&post); err != nil {
// 		return &errorhandler.InternalServerError{Message: err.Error()}
// 	}
// 	return nil
// }
