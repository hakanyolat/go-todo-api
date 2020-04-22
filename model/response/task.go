package response

import "github.com/hakanyolat/go-todo-api/model"

type TaskResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func NewTaskResponse(task model.Task) TaskResponse {
	return TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Completed: task.Completed,
	}
}

func NewTasksResponse(tasks []model.Task) []TaskResponse {
	response := make([]TaskResponse, 0)
	for _, task := range tasks {
		response = append(response, NewTaskResponse(task))
	}
	return response
}
