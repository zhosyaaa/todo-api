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

var NewProject models.Project

func GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID := vars["projectId"]
	ID, err := strconv.ParseInt(projectID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing getProject func")
	}
	ProjectDetails, _ := models.GetProjectById(ID)
	res, _ := json.Marshal(ProjectDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	createProject := &models.Project{}
	utils.ParseBody(r, createProject)
	b := createProject.CreateProject()
	res, _ := json.Marshal(b)
	w.Write(res)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ProjectId := vars["projectId"]
	id, err := strconv.ParseInt(ProjectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing DeleteProject func")
	}
	ProjectDetails := models.DeleteProject(id)
	res, _ := json.Marshal(ProjectDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	var updateProject = &models.Project{}
	utils.ParseBody(r, updateProject)
	vars := mux.Vars(r)
	ProjectId := vars["projectId"]
	id, err := strconv.ParseInt(ProjectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing UpdateProject func")
	}
	ProjectDetails, db := models.GetProjectById(id)
	if updateProject.Title != "" {
		ProjectDetails.Title = updateProject.Title
	}
	if updateProject.Tasks != nil {
		ProjectDetails.Tasks = updateProject.Tasks
	}
	db.Save(&ProjectDetails)
	res, _ := json.Marshal(ProjectDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	allProjects := models.GetAllProjects()
	res, err := json.Marshal(allProjects)
	if err != nil {
		fmt.Println("error while parsing GetAllProjects func")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
