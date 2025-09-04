package request

type CreateSprintRequest struct {
	Title     string `db:"title" json:"title"`
	ProjectID string `db:"project_id" json:"project_id"`
	Deadline  string `db:"deadline" json:"deadline"`
	AssignTo  string `db:"assign_to" json:"assign_to"`
}

type UpdateSprintRequest struct {
	SprintID string `db:"sprint_id" json:"sprint_id"`
	Title    string `db:"title" json:"title"`
	Deadline string `db:"deadline" json:"deadline"`
	AssignTo string `db:"assign_to" json:"assign_to"`
}

type DeleteSprintRequest struct {
	SprintID string `db:"sprint_id" json:"sprint_id"`
}
