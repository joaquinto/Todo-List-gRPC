package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaquinto/Todo-List-gRPC/model"
)

type TodoServiceServer struct{}

var todos []*model.Todo

func (s *TodoServiceServer) CreateTodo(ctx context.Context, request *model.Todo) (*model.TodosResponse, error) {

	todo := &model.Todo{
		Id:          uuid.New().String(),
		Title:       request.GetTitle(),
		Description: request.GetDescription(),
		Completed:   false,
	}

	todos = append(todos, todo)

	response := &model.TodosResponse{
		Todos: todos,
	}

	return response, nil
}

func (s *TodoServiceServer) GetTodo(ctx context.Context, request *model.TodoID) (*model.TodoResponse, error) {
	todoId := request.GetId()
	var response *model.TodoResponse

	for _, todo := range todos {
		if todo.Id == todoId {
			response = &model.TodoResponse{
				Todo: todo,
			}
		}
	}

	return response, nil
}

func (s *TodoServiceServer) GetAllTodo(ctx context.Context, request *model.GetTodos) (*model.TodosResponse, error) {

	response := &model.TodosResponse{
		Todos: todos,
	}

	return response, nil
}

func (s *TodoServiceServer) EditTodo(ctx context.Context, request *model.Todo) (*model.TodosResponse, error) {
	todoId := request.GetId()

	for _, todo := range todos {
		if todo.Id == todoId {
			todo.Title = request.GetTitle()
			todo.Description = request.GetDescription()
		}
	}

	response := &model.TodosResponse{
		Todos: todos,
	}

	return response, nil
}

func (s *TodoServiceServer) MarkTodo(ctx context.Context, request *model.TodoID) (*model.TodosResponse, error) {
	todoId := request.GetId()

	for _, todo := range todos {
		if todo.Id == todoId {
			todo.Completed = !todo.Completed
		}
	}

	response := &model.TodosResponse{
		Todos: todos,
	}

	return response, nil
}

func (s *TodoServiceServer) DeleteTodo(ctx context.Context, request *model.TodoID) (*model.TodosResponse, error) {
	todoId := request.GetId()

	for i, todo := range todos {
		if todo.Id == todoId {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}

	response := &model.TodosResponse{
		Todos: todos,
	}

	return response, nil
}
