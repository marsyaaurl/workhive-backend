package repository

import (
	"github.com/marsyaaurl/workhive-backend/entity"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	FindAll(projectOwner string) ([]entity.Project, error)                  // get all projects
	Create(project entity.Project) (entity.Project, error)                  // add new project
	Update(project entity.Project) (entity.Project, error)                  // update project
	FindByID(projectID string, projectOwner string) (entity.Project, error) // get project by id
	Delete(projectID string) error                                          // delete project by id
}

type ProjectRepositoryImpl struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &ProjectRepositoryImpl{db: db}
}

func (repository *ProjectRepositoryImpl) FindAll(projectOwner string) ([]entity.Project, error) {
	var projects []entity.Project
	err := repository.db.Where("project_owner = ?", projectOwner).Find(&projects).Error
	return projects, err
}

func (repository *ProjectRepositoryImpl) Create(project entity.Project) (entity.Project, error) {
	err := repository.db.Create(&project).Error
	return project, err
}

func (repository *ProjectRepositoryImpl) Update(project entity.Project) (entity.Project, error) {
	err := repository.db.Model(&entity.Project{}).
		Where("project_id = ? AND project_owner = ?", project.ProjectID, project.ProjectOwner).
		Updates(project).Error
	return project, err
}

func (repository *ProjectRepositoryImpl) FindByID(projectID string, projectOwner string) (entity.Project, error) {
	var project entity.Project
	err := repository.db.Where("project_id = ? AND project_owner = ?", projectID, projectOwner).First(&project).Error
	return project, err
}

func (repository *ProjectRepositoryImpl) Delete(projectID string) error {
	err := repository.db.Where("project_id = ?", projectID).Delete(&entity.Project{}).Error
	return err
}
