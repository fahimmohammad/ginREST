package todo

import (
	"errors"

	"github.com/google/uuid"
)

type serviceInterface interface {
	PostTodo(todo Todo) (Todo, error)
	GetAllTodo() ([]Todo, error)
	GetSingleTodo(todoID string) (Todo, error)
	RemoveTodo(todoID string) error
	EditTodo(todoID string, todo Todo) (Todo, error)
}

// Service struct
type Service struct {
	repo *repository
}

// PostTodo - posts
func (todoService *Service) PostTodo(todo Todo) (Todo, error) {
	todo.ID = generateUUID()

	postResult, err := todoService.repo.createTodo(todo)
	if err != nil {
		return Todo{}, errors.New("cannot create")
	}
	return postResult, nil
}

// GetAllTodo - gets all
func (todoService *Service) GetAllTodo() ([]Todo, error) {
	getResult, err := todoService.repo.readAllTodo()
	if err != nil {
		return []Todo{}, errors.New("cannot get")
	}
	return getResult, nil
}

// GetSingleTodo - gets single
func (todoService *Service) GetSingleTodo(todoID string) (Todo, error) {
	getResult, err := todoService.repo.readSingleTodo(todoID)
	if err != nil {
		return Todo{}, errors.New("cannot get")
	}
	return getResult, nil
}

// RemoveTodo - removes
func (todoService *Service) RemoveTodo(todoID string) error {
	err := todoService.repo.deleteTodo(todoID)
	if err != nil {
		return err
	}
	return nil
}

// EditTodo - edits
func (todoService *Service) EditTodo(todoID string, todo Todo) (Todo, error) {
	updateResult, err := todoService.repo.updateTodo(todoID, todo)
	if err != nil {
		return Todo{}, errors.New("cannot get")
	}
	return updateResult, nil
}

func generateUUID() string {
	return uuid.New().String()
}

// NewTodoService - returns service
func NewTodoService(repo *repository) *Service {
	return &Service{
		repo: repo,
	}
}
