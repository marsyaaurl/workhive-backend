package entity

import "time"

type Worklog struct {
	WorklogID  string    `db:"worklog_id"`
	EmployeeID string    `db:"employee_id"` // FK to Employee
	ClockIn    time.Time `db:"clock_in"`
	ClockOut   time.Time `db:"clock_out"`
	WorkDate   time.Time `db:"work_date"`
	TotalHours float64   `db:"total_hours"`
}
