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
