package response

type EmployeeResponse struct {
	EmployeeID string `db:"employee_id" json:"employee_id"`
	Email      string `db:"email" json:"email"`
	FullName   string `db:"full_name" json:"full_name"`
	Role       string `db:"role" json:"role"`
	Token      string `json:"token,omitempty"`
}
