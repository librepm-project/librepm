package http

import (
	"net/http"

	"apps/api/app/domain"
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
	issueID, err := GetParamUUID(r, "issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	comments, err := c.CommentService.AllByEntity(domain.EntityTypeIssue, issueID)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, CommentSerializer{}.ModelToResponseMany(*comments))
}

func (c CommentController) Create(w http.ResponseWriter, r *http.Request) {
	issueID, err := GetParamUUID(r, "issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	userID := GetUserIDFromRequest(r)

	var req CommentRequest
	if err := DecodeJSON(r, &req); err != nil {
		RespondBadRequest(w)
		return
	}

	comment := domain.CommentModel{
		EntityType: domain.EntityTypeIssue,
		EntityID:   issueID,
		UserID:     userID,
		Content:    req.Content,
	}

	if err := c.CommentService.Create(&comment); err != nil {
		RespondBadRequest(w)
		return
	}

	created, err := c.CommentService.FindByID(comment.ID)
	if err != nil {
		RespondInternalError(w)
		return
	}
	RespondJSON(w, http.StatusCreated, CommentSerializer{}.ModelToResponse(*created))
}

func (c CommentController) Update(w http.ResponseWriter, r *http.Request) {
	commentID, err := GetParamUUID(r, "comment_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}

	var req CommentRequest
	if err := DecodeJSON(r, &req); err != nil {
		RespondBadRequest(w)
		return
	}

	patch := domain.CommentModel{Content: req.Content}
	if err := c.CommentService.Update(commentID, &patch); err != nil {
		RespondBadRequest(w)
		return
	}

	updated, err := c.CommentService.FindByID(commentID)
	if err != nil {
		RespondInternalError(w)
		return
	}
	RespondJSON(w, http.StatusOK, CommentSerializer{}.ModelToResponse(*updated))
}

func (c CommentController) Destroy(w http.ResponseWriter, r *http.Request) {
	commentID, err := GetParamUUID(r, "comment_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	if err := c.CommentService.Destroy(commentID); err != nil {
		RespondBadRequest(w)
		return
	}
	RespondNoContent(w)
}
