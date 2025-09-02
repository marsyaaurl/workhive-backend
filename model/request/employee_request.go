package request

type SignupRequest struct {
	Email    string `db:"email" json:"email" validate:"required,email"`
	FullName string `db:"full_name" json:"full_name" validate:"required"`
	Password string `db:"password" json:"password" validate:"required"`
	Role     string `db:"role" json:"role"`
}

type LoginRequest struct {
	Email    string `db:"email" json:"email" validate:"required,email"`
	Password string `db:"password" json:"password" validate:"required"`
}
