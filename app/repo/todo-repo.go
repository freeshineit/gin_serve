package repo

import (
	"gin_serve/app/model"

	"gorm.io/gorm"
)

type TodoRepo interface {
	InsertTodo(todo model.Todo) (model.Todo, error)
	UpdateTodoStatus(id uint64, status model.Todo_Status_Type) *gorm.DB
	UpdateTodoContent(id uint64, content string) *gorm.DB
	DeleteTodo(id uint64) *gorm.DB
	FindById(id uint64) model.Todo
	FindAll(userId uint64, limit, page int) ([]model.Todo, int64, error)
}

type todoConnection struct {
	connection *gorm.DB
}

// NewTodoRepo new todo repository
func NewTodoRepo(db *gorm.DB) TodoRepo {
	return &todoConnection{
		connection: db,
	}
}

func (db *todoConnection) InsertTodo(todo model.Todo) (model.Todo, error) {
	res := db.connection.Save(&todo)
	return todo, res.Error
}

func (db *todoConnection) UpdateTodoStatus(id uint64, status model.Todo_Status_Type) *gorm.DB {
	var todo model.Todo
	return db.connection.Model(&todo).Where("id = ?", id).Update("status", status)
}

func (db *todoConnection) UpdateTodoContent(id uint64, content string) *gorm.DB {
	var todo model.Todo
	return db.connection.Model(&todo).Where("id = ?", id).Update("content", content)
}

func (db *todoConnection) FindById(id uint64) model.Todo {
	var todo model.Todo
	db.connection.Where("id = ?", id).Take(&todo)
	return todo
}

func (db *todoConnection) FindAll(userId uint64, limit, page int) ([]model.Todo, int64, error) {
	todos := []model.Todo{}
	var total int64
	err := db.connection.Limit(limit).Offset((page-1)*limit).Where("user_id = ?", userId).Order("created_at desc").Find(&todos).Offset(-1).Limit(-1).Count(&total).Error
	return todos, total, err
}

func (db *todoConnection) DeleteTodo(id uint64) *gorm.DB {
	return db.connection.Where("id = ?", id).Delete(&model.Todo{})
}
