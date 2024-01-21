package service

import (
	"github.com/golang-api/dto"
	"github.com/golang-api/entity"
	"github.com/golang-api/errorhandler"
	"github.com/golang-api/helper"
	"github.com/golang-api/repository"
)

type ProfileService interface {
	UpdateProfile(req *dto.UpdateProfileRequest) error
}

type profileService struct {
	repository repository.ProfileRepository
}

func NewProfileService(r repository.ProfileRepository) *profileService {
	return &profileService{
		repository: r}
}

func (s *profileService) UpdateProfile(req *dto.UpdateProfileRequest) error {
	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}
	users := entity.Users{
		UsersId:  req.UsersId,
		Username: req.Username,
		FullName: req.FullName,
		Email:    req.Email,
		Password: passwordHash,
	}
	if req.Username == "" || req.FullName == "" || req.Email == "" || req.Password == "" {
		return &errorhandler.BadRequestError{Message: "must filled "}
	}
	if err := s.repository.UpdateProfile(&users); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}
	return nil
}
