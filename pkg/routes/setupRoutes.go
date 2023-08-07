package routes

import (
	"github.com/gorilla/mux"
	. "github.com/zhosyaaa/todo-api/pkg/controllers"
)

var SetupRoutes = func(router *mux.Router) {
	router.HandleFunc("/projects", GetAllProjects).Methods("GET")
	router.HandleFunc("/projects", CreateProject).Methods("POST")
	router.HandleFunc("/projects/{projectId:[0-9]+}", GetProject).Methods("GET")
	router.HandleFunc("/projects/{projectId:[0-9]+}", UpdateProject).Methods("PUT")
	router.HandleFunc("/projects/{projectId:[0-9]+}/tasks", GetAllTasks).Methods("GET")
	router.HandleFunc("/projects/{projectId:[0-9]+}", DeleteProject).Methods("DELETE")
	router.HandleFunc("/projects/{projectId:[0-9]+}/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/projects/{projectId:[0-9]+}/tasks/{id:[0-9]+}", GetTask).Methods("GET")
	router.HandleFunc("/projects/{projectId:[0-9]+}/tasks/{id:[0-9]+}", UpdateTask).Methods("PUT")
	router.HandleFunc("/projects/{projectId:[0-9]+}/tasks/{id:[0-9]+}", DeleteTask).Methods("DELETE")
	router.HandleFunc("/projects/{projectId:[0-9]+}/tasks/{id:[0-9]+}/complete", CompleteTask).Methods("PUT")
	router.HandleFunc("/projects/{projectId:[0-9]+}/tasks/{id:[0-9]+}/complete", UndoTask).Methods("DELETE")
}
