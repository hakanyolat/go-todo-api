package repository

import (
	"github.com/hakanyolat/go-todo-api/model"
	"github.com/hakanyolat/go-todo-api/model/request"
	"github.com/jinzhu/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Find(id uint64) (*model.Task, error) {
	var task model.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) Create(cr *request.CreateTaskRequest) (*model.Task, error) {
	task := &model.Task{
		Title:     cr.Title,
		Completed: cr.Completed,
	}

	if err := r.db.Save(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) Update(ur *request.UpdateTaskRequest) (*model.Task, error) {
	var err error
	task := ur.Task
	task.Completed = ur.Completed

	if err = r.db.Model(task).Update("completed", task.Completed).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) Delete(task *model.Task) (bool, error) {
	if err := r.db.Delete(task).Error; err != nil {
		return false, err
	}

	return true, nil
}
