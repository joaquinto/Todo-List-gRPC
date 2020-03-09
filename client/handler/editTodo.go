package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	cm "github.com/joaquinto/Todo-List-gRPC/client/model"
	"github.com/joaquinto/Todo-List-gRPC/client/response"
	"github.com/joaquinto/Todo-List-gRPC/model"
)

func (c *Client) EditTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["id"]
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
		Id:          todoID,
		Title:       todo.Title,
		Description: todo.Description,
	}
	res, err := c.ServiceClient.EditTodo(ctx, request)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}
	todosResponse := &model.TodosResponse{
		Todos: res.GetTodos(),
	}
	response.JSON(w, http.StatusOK,
		"Todo Edited successfully", todosResponse)
}
