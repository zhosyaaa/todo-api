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

// /projects/{title}/tasks
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["projectId"]
	ID, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing GetAllTasks func")
		http.Error(w, "Invalid projectId", http.StatusBadRequest)
	}
	project, db := models.GetProjectById(ID)
	if project == nil {
		fmt.Println("error while find project in GetAllTasks fun")
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}
	tasks := []models.Task{}
	if err := db.Model(&project).Related(&tasks).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	utils.ParseBody(r, task)
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

// /projects/{title}/tasks/{id:[0-9]+}
// надо изменить и найти айди проекта а то айди просто не робит
func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["id"]
	id, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing GetTask func")
		http.Error(w, "Invalid projectId", http.StatusBadRequest)
	}
	task, _ := models.GetTaskById(id)
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// /projects/{title}/tasks/{id:[0-9]+}
// надо изменить и найти айди проекта а то айди просто не робит
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updateTask = &models.Task{}
	utils.ParseBody(r, updateTask)
	vars := mux.Vars(r)
	taskId := vars["id"]
	id, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing UpdateTask func")
		http.Error(w, "Invalid TaskId", http.StatusBadRequest)
	}
	task, db := models.GetTaskById(id)
	if updateTask.Title == "" {
		task.Title = updateTask.Title
	}
	if updateTask.Deadline == nil {
		task.Deadline = updateTask.Deadline
	}
	if !updateTask.Done {
		task.Done = updateTask.Done
	}
	if updateTask.ProjectID < 0 {
		task.ProjectID = updateTask.ProjectID
	}
	db.Save(&task)
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTask(http.ResponseWriter, *http.Request) {

}
func CompleteTask(http.ResponseWriter, *http.Request) {

}
func UndoTask(http.ResponseWriter, *http.Request) {

}
