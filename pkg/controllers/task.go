package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zhosyaaa/todo-api/pkg/models"
	"github.com/zhosyaaa/todo-api/pkg/utils"
	"net/http"
	"strconv"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["projectId"]
	ID, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing GetAllTasks func")
		http.Error(w, "Invalid projectId", http.StatusBadRequest)
	}
	tasks, err := models.GetAllTasks(uint(ID))
	if err != nil {
		fmt.Println("error while getting tasks for the project")
		http.Error(w, "Error while getting tasks", http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(tasks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["projectId"]
	id, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing CreateTask func")
		http.Error(w, "Invalid projectId", http.StatusBadRequest)
	}
	task := models.Task{ProjectID: uint(id)}
	_ = utils.ParseBody(r, task)
	b := task.CreateTask()
	res, err := json.Marshal(b)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["projectId"]
	taskID := vars["id"]
	id, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing projectID GetTask func")
		http.Error(w, "Invalid projectId", http.StatusBadRequest)
	}
	taskId, err := strconv.ParseInt(taskID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing taskID GetTask func")
		http.Error(w, "Invalid taskID", http.StatusBadRequest)
	}

	task, _ := models.GetTaskById(uint(id), uint(taskId))
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updateTask = &models.Task{}
	_ = utils.ParseBody(r, updateTask)
	vars := mux.Vars(r)
	projectId := vars["projectId"]
	taskID := vars["id"]
	id, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing projectID UpdateTask func")
		http.Error(w, "Invalid projectId", http.StatusBadRequest)
	}
	taskId, err := strconv.ParseInt(taskID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing taskID UpdateTask func")
		http.Error(w, "Invalid taskID", http.StatusBadRequest)
	}
	task, db := models.GetTaskById(uint(id), uint(taskId))
	if updateTask.Title != "" {
		task.Title = updateTask.Title
	}
	if updateTask.Deadline != nil {
		task.Deadline = updateTask.Deadline
	}
	task.Done = updateTask.Done
	if updateTask.ProjectID >= 0 {
		task.ProjectID = updateTask.ProjectID
	}
	db.Save(&task)
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["projectId"]
	taskID := vars["id"]
	id, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing projectID DeleteTask func")
		http.Error(w, "Invalid projectId", http.StatusBadRequest)
	}
	taskId, err := strconv.ParseInt(taskID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing taskID DeleteTask func")
		http.Error(w, "Invalid taskID", http.StatusBadRequest)
	}
	task, err := models.DeleteTask(uint(id), uint(taskId))
	if err != nil {
		fmt.Println("error while deleting the task")
		http.Error(w, "Error while deleting the task", http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CompleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["projectId"]
	taskID := vars["id"]
	id, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing projectID CompleteTask func")
		http.Error(w, "Invalid projectId", http.StatusBadRequest)
	}
	taskId, err := strconv.ParseInt(taskID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing taskID CompleteTask func")
		http.Error(w, "Invalid taskID", http.StatusBadRequest)
	}
	task, db := models.GetTaskById(uint(id), uint(taskId))
	task.Undo()
	db.Save(task)
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UndoTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["projectId"]
	taskID := vars["id"]
	id, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing projectID UndoTask func")
		http.Error(w, "Invalid projectId", http.StatusBadRequest)
	}
	taskId, err := strconv.ParseInt(taskID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing taskID UndoTask func")
		http.Error(w, "Invalid taskID", http.StatusBadRequest)
	}
	task, db := models.GetTaskById(uint(id), uint(taskId))
	task.Complete()
	db.Save(task)
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
