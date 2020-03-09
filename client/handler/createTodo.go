package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	cm "github.com/joaquinto/Todo-List-gRPC/client/model"
	"github.com/joaquinto/Todo-List-gRPC/client/response"
	"github.com/joaquinto/Todo-List-gRPC/model"
)

type Client struct {
	ServiceClient model.TodoServiceClient
}

func (c *Client) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading data: %v", err.Error())
		return
	}
	todo := &cm.Todo{}
	json.Unmarshal(data, todo)
	errMessage, notValid := cm.ValidateInput(todo)
	if notValid {
		response.JSON(w, http.StatusBadRequest,
			"Validation Error", errMessage)
		return
	}
	todo.Prepare()
	ctx := context.Background()
	request := &model.Todo{
		Title:       todo.Title,
		Description: todo.Description,
	}
	res, err := c.ServiceClient.CreateTodo(ctx, request)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}
	todosResponse := &model.TodosResponse{
		Todos: res.GetTodos(),
	}
	response.JSON(w, http.StatusCreated,
		"Todo saved successfully", todosResponse)
}
