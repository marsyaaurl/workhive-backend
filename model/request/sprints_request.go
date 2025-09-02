package request

import "time"

type CreateSprintRequest struct {
	Title     string    `db:"title" json:"title"`
	ProjectID string    `db:"project_id" json:"project_id"`
	Deadline  time.Time `db:"deadline" json:"deadline"`
	AssignTo  string    `db:"assign_to" json:"assign_to"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type UpdateSprintRequest struct {
	SprintID  string    `db:"sprint_id" json:"sprint_id"`
	Title     string    `db:"title" json:"title"`
	Deadline  time.Time `db:"deadline" json:"deadline"`
	AssignTo  string    `db:"assign_to" json:"assign_to"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type DeleteSprintRequest struct {
	SprintID string `db:"sprint_id" json:"sprint_id"`
}
