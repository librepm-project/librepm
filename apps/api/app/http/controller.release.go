package http

import (
	"encoding/json"
	"net/http"
	"time"

	"apps/api/app/domain"
)

type ReleaseControllerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}

type ReleaseController struct {
	ReleaseService domain.ReleaseServiceInterface
}

type CreateReleaseRequest struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	ReleasedAt *time.Time `json:"released_at"`
}

type UpdateReleaseRequest struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	ReleasedAt *time.Time `json:"released_at"`
}

func (c ReleaseController) Index(w http.ResponseWriter, r *http.Request) {
	releases, err := c.ReleaseService.All()
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, ReleaseSerializer{}.ModelToResponseMany(*releases))
}

func (c ReleaseController) Show(w http.ResponseWriter, r *http.Request) {
	release_id, err := GetParamUUID(r, "release_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	release, err := c.ReleaseService.Show(release_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, ReleaseSerializer{}.ModelToResponse(*release))
}

func (c ReleaseController) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateReleaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondBadRequest(w)
		return
	}

	if req.Name == "" {
		RespondBadRequest(w)
		return
	}

	release := domain.ReleaseModel{
		Name:       req.Name,
		Status:     req.Status,
		ReleasedAt: req.ReleasedAt,
	}

	if err := c.ReleaseService.Create(&release); err != nil {
		RespondBadRequest(w)
		return
	}

	RespondJSON(w, http.StatusCreated, ReleaseSerializer{}.ModelToResponse(release))
}

func (c ReleaseController) Update(w http.ResponseWriter, r *http.Request) {
	release_id, err := GetParamUUID(r, "release_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}

	var req UpdateReleaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondBadRequest(w)
		return
	}

	release := domain.ReleaseModel{
		Name:       req.Name,
		Status:     req.Status,
		ReleasedAt: req.ReleasedAt,
	}

	if err := c.ReleaseService.Update(release_id, &release); err != nil {
		RespondBadRequest(w)
		return
	}

	updated, err := c.ReleaseService.Show(release_id)
	if err != nil {
		RespondBadRequest(w)
		return
	}

	RespondJSON(w, http.StatusOK, ReleaseSerializer{}.ModelToResponse(*updated))
}

func (c ReleaseController) Destroy(w http.ResponseWriter, r *http.Request) {
	release_id, err := GetParamUUID(r, "release_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}

	if err := c.ReleaseService.Destroy(release_id); err != nil {
		RespondBadRequest(w)
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
