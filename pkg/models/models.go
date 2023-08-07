package models

import (
	"fmt"
	"github.com/zhosyaaa/todo-api/pkg/configs"
	"gorm.io/gorm"
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
	db.AutoMigrate(&Project{}, &Task{})
}

func (p *Project) CreateProject() *Project {
	db.Create(p)
	if p.Tasks != nil {
		for i := range p.Tasks {
			p.Tasks[i].ProjectID = p.ID
			p.Tasks[i].CreateTask()
		}
	}
	return p
}
func GetAllProjects() []Project {
	var Projects []Project
	db.Preload("Tasks").Find(&Projects)
	return Projects
}

func GetProjectById(id int64) (*Project, *gorm.DB) {
	var getProject Project
	db := db.Where("ID=?", id).Find(&getProject)
	if db.Error != nil {
		return nil, db
	}
	getProject.Tasks, _ = GetAllTasks(uint(id))
	return &getProject, db
}

func DeleteProject(id int64) *Project {
	var project Project
	db.Where("ID=?", id).Delete(&project)
	if db.Error != nil {
		return nil
	}
	return &project
}

func (t *Task) CreateTask() *Task {
	db.Create(&t)
	return t
}

func GetAllTasks(projectId uint) ([]Task, error) {
	var tasks []Task
	db.Where("project_id = ?", projectId).Find(&tasks)
	return tasks, nil
}

func GetTaskById(projectId uint, taskId uint) (*Task, *gorm.DB) {
	var getTask Task
	result := db.Where("ID = ? AND project_id = ?", taskId, projectId).First(&getTask)
	if result.Error != nil {
		return nil, db
	}

	return &getTask, db
}

func DeleteTask(projectId uint, taskId uint) (*Task, error) {
	var task Task
	result := db.Where("ID = ? AND project_id = ?", taskId, projectId).Delete(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("task not found")
	}
	return &task, nil
}
