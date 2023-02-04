package repo

import (
	"gin_serve/app/model"

	"gorm.io/gorm"
)

type TodoRepo interface {
	InsertTodo(user model.Todo) model.Todo
	UpdateTodo(user model.Todo) model.Todo
	FindById(id string) model.Todo
}

type todoConnect struct {
	connection *gorm.DB
}

// NewTodoRepo new todo repository
func NewTodoRepo(db *gorm.DB) UserRepo {
	return &userConnection{
		connection: db,
	}
}

func (db *todoConnect) InsertTodo(todo model.Todo) model.Todo {
	db.connection.Save(&todo)

	return todo
}

func (db *todoConnect) UpdateTodo(todo model.Todo) model.Todo {
	return model.Todo{}
}

func (db *todoConnect) FindById(id string) model.Todo {
	var todo model.Todo
	db.connection.Where("id = ?", id).Take(&todo)
	return todo
}
