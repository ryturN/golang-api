package repository

import (
	"github.com/golang-api/entity"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	UpdateProfile(profile *entity.Users) error
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *profileRepository {
	return &profileRepository{
		db: db,
	}
}

func (r *profileRepository) UpdateProfile(profile *entity.Users) error {
	usersIsExist := &entity.Users{}
	err := r.db.Where("users_id = ?", profile.UsersId).First(usersIsExist).Error
	if err != nil {
		return err
	}
	err = r.db.Model(usersIsExist).Updates(profile).Error
	if err != nil {

		return err
	}
	return nil
}
