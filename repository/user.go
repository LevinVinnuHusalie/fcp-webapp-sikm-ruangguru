package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, nil
		}
		return model.User{}, err
	}
	return user, nil
	// TODO: replace this
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	var tasc []model.UserTaskCategory
	result := r.db.Raw("SELECT u.id, u.fullname, u.email, t.title as task, t.deadline, t.priority, t.status, concat('Category ', t.category_id) as category FROM users u JOIN tasks t ON u.id = t.user_id").Scan(&tasc)
	// result := r.db.Preload("Category").Preload("task").Scan(&tasc)
	if result.Error != nil {
		return []model.UserTaskCategory{}, result.Error
	}
	return tasc, nil // TODO: replace this
}
