package cmd

import (
	"github.com/gorilla/mux"
	"github.com/zhosyaaa/todo-api/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.SetupRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
