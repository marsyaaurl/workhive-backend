package request

import "time"

type CreateProjectRequest struct {
	Title        string    `db:"title" json:"title"`
	Description  string    `db:"description" json:"description"`
	ProjectOwner string    `db:"project_owner" json:"project_owner"` // FK ke employees
	Status       string    `db:"status" json:"status"`
	Priority     string    `db:"priority" json:"priority"`
	Deadline     time.Time `db:"deadline" json:"deadline"`   // DATE di DB
	AssignTo     string    `db:"assign_to" json:"assign_to"` // FK ke employees
}

type UpdateProjectRequest struct {
	ProjectID    string    `db:"project_id" json:"project_id"`
	Title        string    `db:"title" json:"title"`
	Description  string    `db:"description" json:"description"`
	ProjectOwner string    `db:"project_owner" json:"project_owner"` // FK ke employees
	Status       string    `db:"status" json:"status"`
	Priority     string    `db:"priority" json:"priority"`
	Deadline     time.Time `db:"deadline" json:"deadline"`   // DATE di DB
	AssignTo     string    `db:"assign_to" json:"assign_to"` // FK ke employees
}

type DeleteProjectRequest struct {
	ProjectID string `db:"project_id" json:"project_id"`
}
