# Go Todo REST API Example
A RESTful API for simple todo application with Go

It is a just simple tutorial or example for making simple RESTful API with Go using **gorilla/mux** (A nice mux library) and **gorm** (An ORM for Go)

## Installation & Run
```bash
# Download this project
go get github.com/zhosyaaa/todo-api
```

Before running API server, you should set the database config with yours or set the your database config with my values on [app.go](https://github.com/zhosyaaa/todo-api/blob/master/pkg/configs/app.go)
```go
dsn := "host=localhost user=yourUser password=password dbname=name port=5432 sslmode=disable TimeZone=Asia/Shanghai"

```

```bash
# Build and Run
cd go-todo-rest-api-example
go build
./go-todo-rest-api-example

# API Endpoint : http://127.0.0.1:3000
```

## Structure
```
├── pkg
│   ├── config          // Configuration
│   │   ├── app.go    // for connecting db

│   └── controllers
│   │   ├── projects.go  // Our API core handlers for projects
│   │   ├── task.go       // Our API core handlers for tasks

│   └── models
│   │   ├── models.go   //APIs for models

│   └── routes
│   │   ├── setupRoutes.go  //our API for routing

│   └── utils
│   │   ├── utils.go  //our API for decoding JSON data from an HTTP request


└── main.go
```

## API

#### /projects
* `GET` : Get all projects
* `POST` : Create a new project

#### /projects/:title
* `GET` : Get a project
* `PUT` : Update a project
* `DELETE` : Delete a project

#### /projects/:title/archive
* `PUT` : Archive a project
* `DELETE` : Restore a project

#### /projects/:title/tasks
* `GET` : Get all tasks of a project
* `POST` : Create a new task in a project

#### /projects/:title/tasks/:id
* `GET` : Get a task of a project
* `PUT` : Update a task of a project
* `DELETE` : Delete a task of a project

#### /projects/:title/tasks/:id/complete
* `PUT` : Complete a task of a project
* `DELETE` : Undo a task of a project

