package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/pritamguha/gqlgen-todos/graph/model"
	"github.com/pritamguha/gqlgen-todos/logic"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo, err := logic.CreateTodo(ctx, input.Name)

	return &todo, err
}

// EditTodo is the resolver for the editTodo field.
func (r *mutationResolver) EditTodo(ctx context.Context, input model.EditTodo) (*model.Todo, error) {
	todo, err := logic.EditTodo(ctx, input.ID, input.Name, input.IsActive)

	return &todo, err
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (bool, error) {
	isDeleted, err := logic.DeleteTodo(ctx, id)
	return isDeleted, err
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := logic.Todos(ctx)

	return todos, err
}

// GetTodoItem is the resolver for the getTodoItem field.
func (r *queryResolver) GetTodoItem(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: GetTodoItem - getTodoItem"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }