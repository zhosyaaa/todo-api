package models

import (
	"github.com/jinzhu/gorm"
	"github.com/zhosyaaa/todo-api/pkg/configs"
	"time"
)

var db *gorm.DB

type Project struct {
	gorm.Model
	Title string `gorm:"unique" json:"title"`
	Tasks []Task `gorm:"ForeignKey:ProjectID" json:"tasks"`
}

type Task struct {
	gorm.Model
	Title     string     `json:"title"`
	Deadline  *time.Time `gorm:"default:null" json:"deadline"`
	Done      bool       `json:"done"`
	ProjectID uint       `json:"project_id"`
}

func (t *Task) Complete() {
	t.Done = true
}

func (t *Task) Undo() {
	t.Done = false
}

func init() {
	configs.ConnectDB()
	db = configs.GetDB()
	db = DBMigrate(db)
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Project{}, &Task{})
	db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	return db
}

func (p *Project) CreateProject() *Project {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetAllProjects() []Project {
	var Projects []Project
	db.Find(&Projects)
	return Projects
}
func GetProjectById(id int64) (*Project, *gorm.DB) {
	var getProject Project
	db := db.Where("ID=?", id).Find(&getProject)
	return &getProject, db
}
func DeleteProject(id int64) Project {
	var project Project
	db.Where("ID=?", id).Delete(project)
	return project
}

func (t *Task) CreateTask() *Task {
	db.NewRecord(t)
	db.Create(&t)
	return t
}

func GetAllTasks() []Task {
	var Tasks []Task
	db.Find(&Tasks)
	return Tasks
}
func GetTaskById(id int64) (*Task, *gorm.DB) {
	var getTask Task
	db := db.Where("ID=?", id).Find(&getTask)
	return &getTask, db
}
func DeleteTask(id int64) Task {
	var task Task
	db.Where("ID=?", id).Delete(task)
	return task
}
