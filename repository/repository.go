package repository

import (
    _"gorm.io/driver/mysql"
    "gorm.io/gorm"
	"TodoApp/models"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
    return &TodoRepository{db: db}
}

func (repo TodoRepository) Add (model *models.Todo) error {
	if err := repo.db.Create(model).Error; err != nil {
		return err
	}
	return nil
}

func (repo TodoRepository) List () []models.Todo {
	todos := []models.Todo{}
	if err := repo.db.Select("ID","Text").Find(&todos).Error; err != nil {
		panic(err)
	}
	return todos
}

func (repo TodoRepository) Update (newTodo *models.Todo) error {
	if err := repo.db.Model(&models.Todo{}).Where("id = ?", newTodo.ID).Updates(newTodo).Error; err != nil {
		return err
	}
	return nil
}

func (repo TodoRepository) Delete (id int) error {
	if err := repo.db.Delete(&models.Todo{},id).Error; err != nil {
		return err
	}
	
	return nil
}