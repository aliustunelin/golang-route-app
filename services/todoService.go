package services

import (
	"golang-route-app/dto"
	"golang-route-app/models"
	"golang-route-app/repos"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultTodoService struct {
	Repo repos.TodoRepos
}

type TodoService interface {
	TodoInsert(todo models.Todo) (*dto.TodoDTO, error)
	TodoGetAll() ([]models.Todo, error)
	TodoDelete(id primitive.ObjectID) (bool, error)
}

func (t DefaultTodoService) TodoInsert(todo models.Todo) (*dto.TodoDTO, error) {
	var res dto.TodoDTO
	if len(todo.Title) <= 2 {
		res.Status = false
		return &res, nil
	}
	result, err := t.Repo.Insert(todo)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}

	res = dto.TodoDTO{Status: result}
	return &res, nil
}

func (t DefaultTodoService) TodoGetAll() ([]models.Todo, error) {
	result, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t DefaultTodoService) TodoDelete(id primitive.ObjectID) (bool, error) {
	result, err := t.Repo.Delete(id)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func NewTodoService(Repo repos.TodoRepos) DefaultTodoService {
	return DefaultTodoService{Repo: Repo}
}
