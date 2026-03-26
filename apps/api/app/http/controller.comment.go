package http

import (
	"encoding/json"
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
	"libs/jwt_utils"
)

type CommentControllerInterface interface {
	IndexByIssue(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}

type CommentController struct {
	CommentService domain.CommentServiceInterface
}

func (c CommentController) IndexByIssue(w http.ResponseWriter, r *http.Request) {
	issueID, err := http_utils.GetParamUUID(r, "issue_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	comments, err := c.CommentService.AllByEntity("issue", issueID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, CommentSerializer{}.ModelToResponseMany(*comments))
}

func (c CommentController) Create(w http.ResponseWriter, r *http.Request) {
	issueID, err := http_utils.GetParamUUID(r, "issue_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID := jwt_utils.GetTokenInfoFromRequest(r).UserID

	var req CommentRequest
	json.NewDecoder(r.Body).Decode(&req)

	comment := domain.CommentModel{
		EntityType: "issue",
		EntityID:   issueID,
		UserID:     userID,
		Content:    req.Content,
	}

	if err := c.CommentService.Create(&comment); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	created, err := c.CommentService.FindByID(comment.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusCreated, CommentSerializer{}.ModelToResponse(*created))
}

func (c CommentController) Update(w http.ResponseWriter, r *http.Request) {
	commentID, err := http_utils.GetParamUUID(r, "comment_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req CommentRequest
	json.NewDecoder(r.Body).Decode(&req)

	patch := domain.CommentModel{Content: req.Content}
	if err := c.CommentService.Update(commentID, &patch); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updated, err := c.CommentService.FindByID(commentID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, CommentSerializer{}.ModelToResponse(*updated))
}

func (c CommentController) Destroy(w http.ResponseWriter, r *http.Request) {
	commentID, err := http_utils.GetParamUUID(r, "comment_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := c.CommentService.Destroy(commentID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
