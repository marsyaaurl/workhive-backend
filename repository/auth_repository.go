package repository

import (
	"github.com/marsyaaurl/workhive-backend/entity"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Save(employee entity.Employee) (entity.Employee, error)
	FindByEmail(email string) (entity.Employee, error)
}

type employeeRepositoryImpl struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepositoryImpl{db: db}
}

func (repository *employeeRepositoryImpl) Save(employee entity.Employee) (entity.Employee, error) {
	err := repository.db.Create(&employee).Error
	return employee, err
}

func (repository *employeeRepositoryImpl) FindByEmail(email string) (entity.Employee, error) {
	var employee entity.Employee
	err := repository.db.Where("email = ?", email).First(&employee).Error
	return employee, err
}
