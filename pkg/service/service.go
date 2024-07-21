package service

import "github.com/KhanbalaRashidov/go-restapi/pkg/repository"

type Authorization interface {
}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NeService(repository *repository.Repository) *Service {
	return &Service{}
}
