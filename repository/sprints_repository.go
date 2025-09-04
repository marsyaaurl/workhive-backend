package repository

import (
	"github.com/marsyaaurl/workhive-backend/entity"
	"gorm.io/gorm"
)

type SprintRepository interface {
	FindAll(projectID string) ([]entity.Sprint, error)
	Create(sprint entity.Sprint) (entity.Sprint, error)
	Update(sprint entity.Sprint) (entity.Sprint, error)
	FindByID(sprintID string) (entity.Sprint, error)
	Delete(sprintID string) error
}

type SprintRepositoryImpl struct {
	db *gorm.DB
}

func NewSprintRepository(db *gorm.DB) SprintRepository {
	return &SprintRepositoryImpl{db: db}
}

func (repository *SprintRepositoryImpl) FindAll(projectID string) ([]entity.Sprint, error) {
	var sprints []entity.Sprint
	err := repository.db.Where("project_id = ?", projectID).Find(&sprints).Error
	return sprints, err
}

func (repository *SprintRepositoryImpl) Create(sprint entity.Sprint) (entity.Sprint, error) {
	err := repository.db.Create(&sprint).Error
	return sprint, err
}

func (repository *SprintRepositoryImpl) Update(sprint entity.Sprint) (entity.Sprint, error) {
	err := repository.db.Where("sprint_id = ?", sprint.SprintID).Updates(sprint).Error
	return sprint, err
}

func (repository *SprintRepositoryImpl) FindByID(sprintID string) (entity.Sprint, error) {
	var sprint entity.Sprint
	err := repository.db.Where("sprint_id = ?", sprintID).First(&sprint).Error
	return sprint, err
}

func (repository *SprintRepositoryImpl) Delete(sprintID string) error {
	err := repository.db.Where("sprint_id = ?", sprintID).Delete(&entity.Sprint{}).Error
	return err
}
