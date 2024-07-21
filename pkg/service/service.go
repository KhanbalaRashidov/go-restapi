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
}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NeService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		TodoList:      NewTodoListService(repository.TodoList),
	}
}
