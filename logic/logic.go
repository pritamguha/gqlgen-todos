package logic

import (
	"context"

	"github.com/pritamguha/gqlgen-todos/graph/model"
	"github.com/pritamguha/gqlgen-todos/repository"
)

func CreateTodo(ctx context.Context, name string) (*model.Todo, error) {
	todo, err := repository.CreateTodo(ctx, name)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := repository.Todos(ctx)

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

func GetTodoItem(ctx context.Context, id string) (*model.Todo, error) {
	todo, err := repository.GetTodoItem(ctx, id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
