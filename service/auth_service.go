package service

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/marsyaaurl/workhive-backend/entity"
	"github.com/marsyaaurl/workhive-backend/model/request"
	"github.com/marsyaaurl/workhive-backend/model/response"
	"github.com/marsyaaurl/workhive-backend/repository"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("supersecretkey")

type AuthService interface {
	Signup(req request.SignupRequest) (response.EmployeeResponse, error)
	Login(req request.LoginRequest) (response.EmployeeResponse, error)
}

type authServiceImpl struct {
	employeeRepo repository.EmployeeRepository
}

func NewAuthService(repo repository.EmployeeRepository) AuthService {
	return &authServiceImpl{employeeRepo: repo}
}

func (service *authServiceImpl) Signup(req request.SignupRequest) (response.EmployeeResponse, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	employee := entity.Employee{
		EmployeeID: uuid.New().String(),
		Email:      req.Email,
		FullName:   req.FullName,
		Password:   string(hashedPassword),
		Role:       req.Role,
	}

	saved, err := service.employeeRepo.Save(employee)
	if err != nil {
		return response.EmployeeResponse{}, err
	}

	return response.EmployeeResponse{
		EmployeeID: saved.EmployeeID,
		Email:      saved.Email,
		FullName:   saved.FullName,
		Role:       saved.Role,
	}, nil
}

func (service *authServiceImpl) Login(req request.LoginRequest) (response.EmployeeResponse, error) {
	employee, err := service.employeeRepo.FindByEmail(req.Email)
	if err != nil {
		return response.EmployeeResponse{}, err
	}

	// check password
	if bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(req.Password)) != nil {
		return response.EmployeeResponse{}, err
	}

	// generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": employee.EmployeeID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, _ := token.SignedString(jwtKey)

	return response.EmployeeResponse{
		EmployeeID: employee.EmployeeID,
		Email:      employee.Email,
		FullName:   employee.FullName,
		Role:       employee.Role,
		Token:      tokenString,
	}, nil
}
