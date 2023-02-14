package service

import (
	"errors"
	"gin_serve/app/dto"
	"gin_serve/app/model"
	"gin_serve/app/repo"
	"gin_serve/message"
	"log"

	"github.com/mashingan/smapping"
)

type TodoService interface {
	CreateTodo(todo dto.TodoCreateDTO, useID uint64) (dto.TodoDTO, error)
	FindById(id uint64) model.Todo
	UpdateTodoStatus(id uint64, status model.Todo_Status_Type, userId uint64) (bool, error)
	UpdateTodoContent(id uint64, content string, userId uint64) (bool, error)
	DeleteTodo(id uint64, userId uint64) (bool, error)
	FindAll(userId uint64, limit, page, size int) ([]dto.TodoDTO, int64, error)
}

type todoService struct {
	todoRepos repo.TodoRepo
}

func NewTodoService(todoRepo repo.TodoRepo) TodoService {
	return &todoService{
		todoRepos: todoRepo,
	}
}

func (service *todoService) CreateTodo(todo dto.TodoCreateDTO, userID uint64) (dto.TodoDTO, error) {
	todoToCreate := model.Todo{
		Content: todo.Content,
		UserID:  userID,
	}

	t, err := service.todoRepos.InsertTodo(todoToCreate)

	var todoDTO = dto.TodoDTO{}

	if err == nil {
		err := smapping.FillStruct(&todoDTO, smapping.MapFields(&t))

		if err != nil {
			log.Fatalf("Failed map %v", err)
		}

		return todoDTO, nil
	}

	return todoDTO, err

}

func (service *todoService) FindById(id uint64) model.Todo {
	return service.todoRepos.FindById(id)
}

func (service *todoService) FindAll(userId uint64, limit, page, size int) ([]dto.TodoDTO, int64, error) {

	mTodos, total, err := service.todoRepos.FindAll(userId, limit, page, size)

	if err != nil {
		return []dto.TodoDTO{}, 0, err
	}

	var todos []dto.TodoDTO
	for _, t := range mTodos {

		todo := dto.TodoDTO{}
		err := smapping.FillStruct(&todo, smapping.MapFields(&t))

		if err != nil {
			log.Fatalf("Failed map %v", err)
		}

		todos = append(todos, todo)
	}

	return todos, total, nil
}

func (service *todoService) UpdateTodoStatus(id uint64, status model.Todo_Status_Type, userID uint64) (bool, error) {
	todo := service.FindById(id)

	if todo.UserID == userID {
		res := service.todoRepos.UpdateTodoStatus(id, status)
		if res.Error == nil {
			return true, nil
		}
		return false, res.Error
	}
	return false, errors.New(message.Unauthorized)
}

func (service *todoService) UpdateTodoContent(id uint64, content string, userID uint64) (bool, error) {

	todo := service.FindById(id)

	if todo.UserID == userID {
		res := service.todoRepos.UpdateTodoContent(id, content)
		if res.Error == nil {
			return true, nil
		}
		return false, res.Error
	}
	return false, errors.New(message.Unauthorized)
}

func (service *todoService) DeleteTodo(id, userID uint64) (bool, error) {
	todo := service.FindById(id)

	if todo.UserID == userID {
		res := service.todoRepos.DeleteTodo(id)
		if res.Error == nil {
			return true, nil
		}
		return false, res.Error
	}

	return false, errors.New(message.Unauthorized)
}
