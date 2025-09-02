package response

import "time"

type DailyResponse struct {
	DailyID   string    `db:"daily_id" json:"daily_id"`
	Title     string    `db:"title" json:"title"`
	ProjectID string    `db:"project_id" json:"project_id"`
	SprintID  string    `db:"sprint_id" json:"sprint_id"`
	Deadline  time.Time `db:"deadline" json:"deadline"`
	AssignTo  string    `db:"assign_to" json:"assign_to"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
