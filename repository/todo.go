package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pritamguha/gqlgen-todos/graph/model"
)

func CreateTodo(ctx context.Context, task string) (*model.Todo, error) {
	var id string

	query := `INSERT INTO public.todo
	(name)
	VALUES($1) RETURNING id`

	err := Db.QueryRow(query, task).Scan(&id)

	if err == nil {
		todo, err2 := GetTodoItem(ctx, id)

		if err2 != nil {
			return nil, err2
		}
		return todo, nil
	}
	return nil, err
}

func Todos(ctx context.Context) ([]*model.Todo, error) {
	result := []*model.Todo{}

	query := `SELECT * FROM public.todo`

	row, err := Db.Query(query)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		eachRow := model.Todo{}
		err2 := row.Scan(&eachRow.ID,
			&eachRow.Name,
			&eachRow.DateCreated,
			&eachRow.IsActive,
		)

		if err2 != nil {
			return nil, err2
		}
		result = append(result, &eachRow)
	}
	return result, nil
}

func EditTodo(ctx context.Context, id string, name *string, isActive *bool) (model.Todo, error) {
	todo := model.Todo{}

	query := ` update public.todo SET  `
	var inputargs []interface{}
	if name != nil {
		query += `name = ? `
		inputargs = append(inputargs, &name)
	}

	if isActive != nil {
		if name != nil {
			query += ` , `
		}
		query += `isActive = ?`
		inputargs = append(inputargs, &isActive)
	}
	query += ` where id = ? returning id`
	inputargs = append(inputargs, id)

	query = sqlx.Rebind(sqlx.DOLLAR, query) // sqlx.rebind takes two input and replace ? to the $ symbol
	sqlErr := Db.QueryRowContext(ctx, query, inputargs...).Scan(&todo.ID)
	if sqlErr != nil {
		return todo, sqlErr
	}

	query2 := `SELECT * from public.todo where id = ` + todo.ID

	err2 := Db.QueryRow(query2).Scan(&todo.ID, &todo.Name, &todo.DateCreated, &todo.IsActive)

	if err2 != nil {
		return todo, err2
	}
	return todo, nil
}

func DeleteTodo(ctx context.Context, id string) (bool, error) {
	query := `DELETE from public.todo where id = ` + id

	_, err := Db.Exec(query)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetTodoItem(ctx context.Context, id string) (*model.Todo, error) {
	todo := model.Todo{}
	query := `SELECT * from public.todo where id = ` + id

	err := Db.QueryRow(query).Scan(&todo.ID, &todo.Name, &todo.DateCreated, &todo.IsActive)

	if err != nil {
		return nil, err
	}
	return &todo, nil
}
