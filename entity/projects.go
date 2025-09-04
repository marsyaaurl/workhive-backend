package entity

import (
	"time"
)

type Project struct {
	ProjectID    string    `db:"project_id" json:"project_id"`
	Title        string    `db:"title" json:"title"`
	Description  string    `db:"description" json:"description"`
	ProjectOwner string    `db:"project_owner" json:"project_owner"` // FK ke employees
	Status       string    `db:"status" json:"status"`
	Priority     string    `db:"priority" json:"priority"`
	Deadline     string    `db:"deadline" json:"deadline"`   // DATE di DB
	AssignTo     string    `db:"assign_to" json:"assign_to"` // FK ke employees
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
