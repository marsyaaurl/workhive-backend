package response

import "time"

type ProjectResponse struct {
	ProjectID    string    `db:"project_id" json:"project_id"`
	Title        string    `db:"title" json:"title"`
	Description  string    `db:"description" json:"description"`
	ProjectOwner string    `db:"project_owner" json:"project_owner"`
	Status       string    `db:"status" json:"status"`
	Priority     string    `db:"priority" json:"priority"`
	Deadline     time.Time `db:"deadline" json:"deadline"`
	AssignTo     string    `db:"assign_to" json:"assign_to"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
