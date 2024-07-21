package service

import (
	gorestapi "github.com/KhanbalaRashidov/go-restapi"
	"github.com/KhanbalaRashidov/go-restapi/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list gorestapi.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]gorestapi.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (gorestapi.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}

func (s *TodoListService) Update(userId, id int, input gorestapi.UpdateListInput) error {
	return s.repo.Update(userId, id, input)
}
