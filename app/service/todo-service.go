package service

import (
	"gin_serve/app/dto"
	"gin_serve/app/model"
	"gin_serve/app/repo"
	"log"

	"github.com/mashingan/smapping"
)

type TodoService interface {
	CreateTodo(todo dto.TodoCreateDTO) model.Todo
	FindById(id string) model.Todo
	UpdateTodo(todo dto.TodoUpdateDTO) bool
}

type todoService struct {
	todoRepos repo.TodoRepo
}

func NewTodoRepo(todoRepo repo.TodoRepo) TodoService {
	return &todoService{
		todoRepos: todoRepo,
	}
}

func (service *todoService) CreateTodo(todo dto.TodoCreateDTO) model.Todo {
	todoToCreate := model.Todo{}
	err := smapping.FillStruct(&todoToCreate, smapping.MapFields(&todo))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	return service.todoRepos.InsertTodo(todoToCreate)
}

func (service *todoService) FindById(id string) model.Todo {
	return service.todoRepos.FindById(id)
}

func (service *todoService) UpdateTodo(todo dto.TodoUpdateDTO) bool {
	todoToUpdate := model.Todo{}
	err := smapping.FillStruct(&todoToUpdate, smapping.MapFields(&todo))

	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	service.todoRepos.UpdateTodo(todoToUpdate)
	return false
}
