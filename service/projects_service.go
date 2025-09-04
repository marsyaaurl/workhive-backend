package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/marsyaaurl/workhive-backend/entity"
	"github.com/marsyaaurl/workhive-backend/model/request"
	"github.com/marsyaaurl/workhive-backend/model/response"
	"github.com/marsyaaurl/workhive-backend/repository"
)

type ProjectService interface {
	FindAll(projectOwner string) ([]response.ProjectResponse, error)
	Create(req request.CreateProjectRequest) (response.ProjectResponse, error)
	Update(req request.UpdateProjectRequest) (response.ProjectResponse, error)
	FindByID(projectID string, projectOwner string) (response.ProjectResponse, error)
	Delete(projectID string) error
}

type projectServiceImpl struct {
	projectRepo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectServiceImpl{projectRepo: repo}
}

func (service *projectServiceImpl) FindAll(projectOwner string) ([]response.ProjectResponse, error) {
	projects, err := service.projectRepo.FindAll(projectOwner)
	if err != nil {
		return nil, err
	}

	var responses []response.ProjectResponse
	for _, p := range projects {
		responses = append(responses, mapToProjectResponse(p))
	}
	return responses, nil
}

func (service *projectServiceImpl) Create(req request.CreateProjectRequest) (response.ProjectResponse, error) {
	project := entity.Project{
		ProjectID:    uuid.New().String(),
		Title:        req.Title,
		Description:  req.Description,
		ProjectOwner: req.ProjectOwner,
		Status:       req.Status,
		Priority:     req.Priority,
		Deadline:     req.Deadline.Format("2006-01-02T00:00:00Z"),
		AssignTo:     req.AssignTo,
	}
	saved, err := service.projectRepo.Create(project)
	if err != nil {
		return response.ProjectResponse{}, err
	}
	return mapToProjectResponse(saved), nil
}

func (service *projectServiceImpl) Update(req request.UpdateProjectRequest) (response.ProjectResponse, error) {
	project := entity.Project{
		ProjectID:    req.ProjectID,
		Title:        req.Title,
		Description:  req.Description,
		ProjectOwner: req.ProjectOwner,
		Status:       req.Status,
		Priority:     req.Priority,
		Deadline:     req.Deadline.Format("2006-01-02T00:00:00Z"),
		AssignTo:     req.AssignTo,
	}
	updated, err := service.projectRepo.Update(project)
	if err != nil {
		return response.ProjectResponse{}, err
	}
	return mapToProjectResponse(updated), nil
}

func (service *projectServiceImpl) FindByID(projectID string, projectOwner string) (response.ProjectResponse, error) {
	project, err := service.projectRepo.FindByID(projectID, projectOwner)
	if err != nil {
		return response.ProjectResponse{}, err
	}
	return mapToProjectResponse(project), nil
}

func (service *projectServiceImpl) Delete(projectID string) error {
	return service.projectRepo.Delete(projectID)
}

func mapToProjectResponse(project entity.Project) response.ProjectResponse {
	var deadline time.Time
	if project.Deadline != "" {
		deadline, _ = time.Parse("2006-01-02", project.Deadline)
	}
	return response.ProjectResponse{
		ProjectID:    project.ProjectID,
		Title:        project.Title,
		Description:  project.Description,
		ProjectOwner: project.ProjectOwner,
		Status:       project.Status,
		Priority:     project.Priority,
		Deadline:     deadline,
		AssignTo:     project.AssignTo,
		CreatedAt:    project.CreatedAt,
	}
}
