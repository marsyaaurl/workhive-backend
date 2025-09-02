package request

import "time"

type CreateDailyRequest struct {
	Title     string    `db:"title" json:"title"`
	ProjectID string    `db:"project_id" json:"project_id"` // FK ke Project
	SprintID  string    `db:"sprint_id" json:"sprint_id"`   // FK ke Sprint
	Deadline  time.Time `db:"deadline" json:"deadline"`
	AssignTo  string    `db:"assign_to" json:"assign_to"` // FK ke Employee
}

type UpdateDailyRequest struct {
	DailyID  string    `db:"daily_id" json:"daily_id"`
	Title    string    `db:"title" json:"title"`
	Deadline time.Time `db:"deadline" json:"deadline"`
	AssignTo string    `db:"assign_to" json:"assign_to"` // FK ke Employee
}

type DeleteDailyRequest struct {
	DailyID string `db:"daily_id" json:"daily_id"`
}
