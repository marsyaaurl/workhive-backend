package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/marsyaaurl/workhive-backend/entity"
	"github.com/marsyaaurl/workhive-backend/model/request"
	"github.com/marsyaaurl/workhive-backend/model/response"
	"github.com/marsyaaurl/workhive-backend/repository"
)

type SprintService interface {
	FindAll(projectID string) ([]response.SprintResponse, error)
	Create(req request.CreateSprintRequest) (response.SprintResponse, error)
	Update(req request.UpdateSprintRequest) (response.SprintResponse, error)
	FindByID(sprintID string) (response.SprintResponse, error)
	Delete(sprintID string) error
}

type sprintServiceImpl struct {
	sprintRepo repository.SprintRepository
}

func NewSprintService(repo repository.SprintRepository) SprintService {
	return &sprintServiceImpl{sprintRepo: repo}
}

func (service *sprintServiceImpl) FindAll(projectID string) ([]response.SprintResponse, error) {
	sprints, err := service.sprintRepo.FindAll(projectID)
	if err != nil {
		return nil, err
	}

	var responses []response.SprintResponse
	for _, s := range sprints {
		responses = append(responses, mapToServiceResponse(s))
	}

	return responses, nil
}

func (service *sprintServiceImpl) Create(req request.CreateSprintRequest) (response.SprintResponse, error) {
	sprint := entity.Sprint{
		SprintID:  uuid.New().String(),
		Title:     req.Title,
		ProjectID: req.ProjectID,
		Deadline:  req.Deadline,
		AssignTo:  req.AssignTo,
	}

	saved, err := service.sprintRepo.Create(sprint)
	if err != nil {
		return response.SprintResponse{}, err
	}
	return mapToServiceResponse(saved), nil
}

func (service *sprintServiceImpl) Update(req request.UpdateSprintRequest) (response.SprintResponse, error) {
	sprint := entity.Sprint{
		SprintID: req.SprintID,
		Title:    req.Title,
		Deadline: req.Deadline,
		AssignTo: req.AssignTo,
	}

	saved, err := service.sprintRepo.Update(sprint) // <-- Panggil fungsi yang benar
	if err != nil {
		return response.SprintResponse{}, err
	}
	return mapToServiceResponse(saved), nil
}

func (service *sprintServiceImpl) FindByID(sprintID string) (response.SprintResponse, error) {
	sprint, err := service.sprintRepo.FindByID(sprintID)
	if err != nil {
		return response.SprintResponse{}, err
	}
	return mapToServiceResponse(sprint), nil
}

func (service *sprintServiceImpl) Delete(sprintID string) error {
	return service.sprintRepo.Delete(sprintID)
}

func mapToServiceResponse(sprint entity.Sprint) response.SprintResponse {
	var deadline time.Time
	if sprint.Deadline != "" {
		deadline, _ = time.Parse("2006-01-02", sprint.Deadline)
	}
	return response.SprintResponse{
		SprintID:  sprint.SprintID,
		Title:     sprint.Title,
		ProjectID: sprint.ProjectID,
		Deadline:  deadline,
		AssignTo:  sprint.AssignTo,
		CreatedAt: sprint.CreatedAt,
	}
}
