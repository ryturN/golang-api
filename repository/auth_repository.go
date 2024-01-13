package repository

import (
	"errors"

	"github.com/golang-api/entity"
	"gorm.io/gorm"
)

type AuthRepository interface {
	EmailExists(email string) bool
	Register(req *entity.Users) error
	GetUserByUsername(username string) (*entity.Users, error)
	UsernameExists(username string) bool
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(user *entity.Users) error {
	err := r.db.Create(&user).Error

	return err
}

func (r *authRepository) EmailExists(email string) bool {
	var user entity.Users
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return !errors.Is(err, gorm.ErrRecordNotFound)
	}
	return true
}

func (r *authRepository) UsernameExists(username string) bool {
	var users entity.Users
	if err := r.db.Where("username = ?", username).First(&users).Error; err != nil {
		return !errors.Is(err, gorm.ErrRecordNotFound)

	}
	return true
}

// func (r *authRepository) Login(req *entity.Login) error {
// 	var user entity.Users
// 	if err := r.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
// 		return &errorhandler.BadRequestError{Message: "email not found"}
// 	}

// 	// Compare the provided password with the hashed password from the database
// 	if err := helper.CheckPasswordHash(user.Password, req.Password); err != nil {
// 		return &errorhandler.BadRequestError{Message: "password incorrect"}
// 	}

// 	// Successful login
// 	return nil
// }

func (r *authRepository) GetUserByUsername(username string) (*entity.Users, error) {
	var users entity.Users

	err := r.db.First(&users, "username =?", username).Error

	return &users, err
}

// func (s *authService) Login(req *dto.LoginRequest) error {
// 	var user entity.Login
// 	if err := s.repository.("username =?", req.Username).First(&user).Error; err != nil {
// 		return &errorhandler.BadRequestError{Message: "email not found"}
// 	}
// 	if err := helper.CheckPasswordHash(user.Password, req.Password); err != nil {
// 		return &errorhandler.BadRequestError{Message: "password incorrect"}
// 	}
// 	return nil
// }
