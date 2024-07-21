package repository

import (
	gorestapi "github.com/KhanbalaRashidov/go-restapi"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user gorestapi.User) (int, error)
	GetUser(username, password string) (gorestapi.User, error)
}

type TodoList interface {
	Create(userId int, list gorestapi.TodoList) (int, error)
	GetAll(userId int) ([]gorestapi.TodoList, error)
	GetById(userId, listId int) (gorestapi.TodoList, error)
	Update(userId, id int, input gorestapi.UpdateListInput) error
	Delete(userId, id int) error
}

type TodoItem interface {
	Create(listId int, item gorestapi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]gorestapi.TodoItem, error)
	GetById(userId, itemId int) (gorestapi.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input gorestapi.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NeRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
