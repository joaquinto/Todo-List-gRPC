package handler

import (
	"context"
	"net/http"

	"github.com/joaquinto/Todo-List-gRPC/client/response"
	"github.com/joaquinto/Todo-List-gRPC/model"
)

func (c *Client) GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	request := &model.GetTodos{}
	ctx := context.Background()
	res, err := c.ServiceClient.GetAllTodo(ctx, request)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}
	todosResponse := &model.TodosResponse{
		Todos: res.GetTodos(),
	}
	response.JSON(w, http.StatusOK,
		"Todos fetched successfully", todosResponse)
}
