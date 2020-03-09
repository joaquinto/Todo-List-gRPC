package handler

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaquinto/Todo-List-gRPC/client/response"
	"github.com/joaquinto/Todo-List-gRPC/model"
)

func (c *Client) GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["id"]
	ctx := context.Background()
	request := &model.TodoID{
		Id: todoID,
	}
	res, err := c.ServiceClient.GetTodo(ctx, request)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}
	todoResponse := &model.TodoResponse{
		Todo: res.GetTodo(),
	}
	response.JSON(w, http.StatusOK,
		"Todo fetched successfully", todoResponse)
}
