package entity

import "time"

type Task struct {
	TaskID    string    `db:"task_id" json:"task_id"`
	Title     string    `db:"title" json:"title"`
	ProjectID string    `db:"project_id" json:"project_id"` // FK ke Project
	SprintID  string    `db:"sprint_id" json:"sprint_id"`   // FK ke Sprint
	Deadline  time.Time `db:"deadline" json:"deadline"`
	AssignTo  string    `db:"assign_to" json:"assign_to"` // FK ke Employee
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
