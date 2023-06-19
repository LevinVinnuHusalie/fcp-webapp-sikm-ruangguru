package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) Store(task *model.Task) error {
	err := t.db.Create(task).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Update(id int, task *model.Task) error {
	result := t.db.Where("id = ?", id).Updates(task)
	if result.RowsAffected == 0 {
		return errors.New("no task updated")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (t *taskRepository) Delete(id int) error {
	var task model.Task
	result := t.db.Where("id = ?", id).Delete(&task)
	if result.RowsAffected == 0 {
		return errors.New("no task delete")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	err := t.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	var tasks []model.Task
	if err := t.db.Find(&tasks).Error; err != nil {
		return []model.Task{}, err
	}
	return tasks, nil
	// TODO: replace this
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	var tasc []model.TaskCategory
	result := t.db.Raw("SELECT id, title, concat('Category ', category_id) as category FROM tasks WHERE id = ? LIMIT 1", id).Scan(&tasc)
	if result.Error != nil {
		return []model.TaskCategory{}, result.Error
	}
	return tasc, nil
	// TODO: replace this
}
