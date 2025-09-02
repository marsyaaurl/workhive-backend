package entity

import "time"

type DailyStandup struct {
	DailyID       string    `db:"daily_id" json:"daily_id"`
	ProjectID     string    `db:"project_id" json:"project_id"`   // FK ke Project
	EmployeeID    string    `db:"employee_id" json:"employee_id"` // FK ke Employee
	YesterdayWork string    `db:"yesterday_work" json:"yesterday_work"`
	TodayPlan     string    `db:"today_plan" json:"today_plan"`
	Blockers      string    `db:"blockers" json:"blockers"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}
