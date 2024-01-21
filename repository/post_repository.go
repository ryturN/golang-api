package repository

import (
	"github.com/golang-api/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Post) error
	Update(post *entity.Post) error
	Delete(post *entity.Post) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Create(post *entity.Post) error {
	err := r.db.Create(&post).Error
	return err
}

func (r *postRepository) Update(post *entity.Post) error {
	postIsExist := &entity.Post{}
	err := r.db.Where("users_id = ?", post.UsersId).First(postIsExist).Error
	if err != nil {
		return err
	}
	err = r.db.Model(postIsExist).Updates(post).Error
	if err != nil {

		return err
	}
	return nil
}
func (r *postRepository) Delete(post *entity.Post) error {
	getPost := &entity.Post{}
	err := r.db.Where("users_id = ?", post.UsersId).First(getPost).Error
	if err != nil {
		return err
	}
	err = r.db.Model(getPost).Delete(post).Error
	if err != nil {

		return err
	}
	return nil
}
