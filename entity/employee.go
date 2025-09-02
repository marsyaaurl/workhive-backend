package entity

type Employee struct {
	EmployeeID string `db:"employee_id" json:"employee_id"`
	Email      string `db:"email" json:"email"`
	Password   string `db:"password" json:"password"`
	FullName   string `db:"full_name" json:"full_name"`
	Role       string `db:"role" json:"role"`
}
