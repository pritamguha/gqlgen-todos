package logic

import (
	"context"
	"fmt"

	"github.com/pritamguha/gqlgen-todos/graph/model"
	"github.com/pritamguha/gqlgen-todos/repository"
)

func CreateTodo(ctx context.Context, name string) (model.Todo, error) {
	todo, err := repository.CreateTodo(ctx, name)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := repository.Todos(ctx)

	fmt.Println("todos", todos, &todos)

	if err != nil {
		return todos, nil
	}

	return todos, err
}

func EditTodo(ctx context.Context, id string, name *string, isActive *bool) (model.Todo, error) {
	todo, err := repository.EditTodo(ctx, id, name, isActive)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func DeleteTodo(ctx context.Context, id string) (bool, error) {
	isDeleted, err := repository.DeleteTodo(ctx, id)
	if err != nil {
		return isDeleted, err
	}

	return isDeleted, nil
}
