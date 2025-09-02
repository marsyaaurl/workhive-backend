package response

import "time"

type WorklogResponse struct {
	WorklogID  string    `db:"worklog_id" json:"worklog_id"`
	EmployeeID string    `db:"employee_id" json:"employee_id"`
	ProjectID  string    `db:"project_id" json:"project_id"`
	TaskID     string    `db:"task_id" json:"task_id"`
	Hours      float64   `db:"hours" json:"hours"`
	WorkDate   time.Time `db:"work_date" json:"work_date"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
