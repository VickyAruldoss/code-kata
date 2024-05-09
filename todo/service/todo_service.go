package service

import (
	"code-kata/client"
	c "code-kata/config"
	"code-kata/todo/model"
	"fmt"

	color "github.com/fatih/color"

	logger "github.com/sirupsen/logrus"
)

type TodoService interface {
	Get(limit int, isEven bool) (todos []model.Todo, err error)
	Print([]model.Todo) error
}

type todoService struct {
	httpRequestHandler client.HttpRequestHandler
	config             c.Config
}

func NewTodoService(httpRequestHandler client.HttpRequestHandler, config c.Config) TodoService {
	return &todoService{httpRequestHandler, config}
}

func (service todoService) Get(limit int, even bool) (todos []model.Todo, err error) {
	for i := 1; len(todos) < limit; i++ {
		var todo model.Todo
		url := fmt.Sprintf("%s/%d", service.config.TodoApiDetails.Url, i)
		err = service.httpRequestHandler.Get(url, &todo)

		if err != nil {
			logger.Errorf("Error while making http request, error: %v", err)
			return
		}

		if service.canAppend(todo, even) {
			todos = append(todos, todo)
		}

	}
	return
}

func (service todoService) canAppend(todo model.Todo, even bool) bool {
	return !even || (even && todo.Id%2 == 0)
}

func (service todoService) Print(todos []model.Todo) error {
	for _, todo := range todos {
		status := "completed"
		if !todo.IsCompleted {
			status = "not completed"
		}
		str := fmt.Sprintf("Title: %s - %s", todo.Title, status)
		service.printMessage(str, todo.IsCompleted)
	}
	return nil
}

func (service todoService) printMessage(message string, isCompleted bool) {
	if isCompleted {
		color.Green(message)
	} else {
		color.Red(message)
	}
}
