package request

import "time"

type ClockInRequest struct {
	EmployeeID string    `db:"employee_id" json:"employee_id"` // FK to Employee
	ClockIn    time.Time `db:"clock_in" json:"clock_in"`
	WorkDate   time.Time `db:"work_date" json:"work_date"`
}

type ClockOutRequest struct {
	EmployeeID string    `db:"employee_id" json:"employee_id"` // FK to Employee
	ClockOut   time.Time `db:"clock_out" json:"clock_out"`
	WorkDate   time.Time `db:"work_date" json:"work_date"`
}
