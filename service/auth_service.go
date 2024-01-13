package service

import (
	"errors"
	"fmt"

	"github.com/golang-api/dto"
	"github.com/golang-api/entity"
	"github.com/golang-api/errorhandler"
	"github.com/golang-api/helper"
	"github.com/golang-api/repository"
	gonanoid "github.com/matoous/go-nanoid"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	if emailExist := s.repository.EmailExists(req.Email); emailExist {
		return &errorhandler.BadRequestError{Message: "email already exists"}
	}
	if usernameExist := s.repository.UsernameExists(req.Username); usernameExist {
		return &errorhandler.BadRequestError{Message: "username already exists"}
	}
	if req.Password != req.ConfirmPassword {
		return &errorhandler.BadRequestError{Message: "passwords do not match"}
	}

	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}
	userIdPrefix := "user_"
	userIdNano, err := gonanoid.Nanoid()
	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}
	userId := fmt.Sprintf("%v%v", userIdPrefix, userIdNano)

	user := entity.Users{
		UsersId:  userId,
		FullName: req.FullName,
		Username: req.Username,
		Email:    req.Email,
		Password: passwordHash,
	}
	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}
	return nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.GetUserByUsername(req.Username)
	fmt.Println(user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorhandler.NotFoundError{Message: "Wrong username or Password"}
		}
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}
	if err := helper.CheckPasswordHash(user.Password, req.Password); err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Wrong  password"}
	}
	token, err := helper.NewGetJWT(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		UserId:   user.UsersId,
		Username: user.Username,
		FullName: user.FullName,
		Token:    token,
	}
	return &data, nil
}
