package service

import (
	gorestapi "github.com/KhanbalaRashidov/go-restapi"
	"github.com/KhanbalaRashidov/go-restapi/pkg/repository"
)

type Authorization interface {
	CreateUser(user gorestapi.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list gorestapi.TodoList) (int, error)
	GetAll(userId int) ([]gorestapi.TodoList, error)
	GetById(userId, listId int) (gorestapi.TodoList, error)
	Update(userId, id int, input gorestapi.UpdateListInput) error
	Delete(userId, id int) error
}

type TodoItem interface {
	Create(userId, listId int, item gorestapi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]gorestapi.TodoItem, error)
	GetById(userId, itemId int) (gorestapi.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input gorestapi.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NeService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		TodoList:      NewTodoListService(repository.TodoList),
		TodoItem:      NewTodoItemService(repository.TodoItem, repository.TodoList),
	}
}
