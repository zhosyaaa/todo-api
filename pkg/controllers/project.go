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

// /projects/{title} +

func GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID := vars["projectId"]
	ID, err := strconv.ParseInt(projectID, 0, 0)
	if err != nil {
		fmt.Println("error while parsing projectID in GetProject func")
		http.Error(w, "Invalid projectID", http.StatusBadRequest)
		return
	}
	ProjectDetails, _ := models.GetProjectById(ID)
	res, _ := json.Marshal(ProjectDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// /projects +

func CreateProject(w http.ResponseWriter, r *http.Request) {
	createProject := &models.Project{}
	_ = utils.ParseBody(r, createProject)
	b := createProject.CreateProject()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Use StatusCreated for successful resource creation
	w.Write(res)
}

// /projects/{title} +

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ProjectId := vars["projectId"]
	id, err := strconv.ParseInt(ProjectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing projectID in DeleteProject func")
		http.Error(w, "Invalid projectID", http.StatusBadRequest)
		return
	}
	ProjectDetails := models.DeleteProject(id)
	res, _ := json.Marshal(ProjectDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// /projects/{title} +

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	var updateProject = &models.Project{}
	_ = utils.ParseBody(r, updateProject)
	vars := mux.Vars(r)
	ProjectId := vars["projectId"]
	id, err := strconv.ParseInt(ProjectId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing projectID in UpdateProject func")
		http.Error(w, "Invalid projectID", http.StatusBadRequest)
		return
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

// /projects +

func GetAllProjects(w http.ResponseWriter, _ *http.Request) {
	allProjects := models.GetAllProjects()
	res, err := json.Marshal(allProjects)
	if err != nil {
		fmt.Println("error while marshaling projects to JSON")
		http.Error(w, "Error while processing data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
