package http

import (
	"net/http"

	"apps/api/app/domain"
	"libs/http_utils"
	"libs/jwt_utils"

	"github.com/google/uuid"
)

type AttachmentControllerInterface interface {
	IndexByIssue(w http.ResponseWriter, r *http.Request)
	IndexByProject(w http.ResponseWriter, r *http.Request)
	CreateForIssue(w http.ResponseWriter, r *http.Request)
	CreateForProject(w http.ResponseWriter, r *http.Request)
	Download(w http.ResponseWriter, r *http.Request)
	Destroy(w http.ResponseWriter, r *http.Request)
}

type AttachmentController struct {
	AttachmentService    domain.AttachmentServiceInterface
	IssueAuditLogService domain.IssueAuditLogServiceInterface
}

func (c AttachmentController) IndexByIssue(w http.ResponseWriter, r *http.Request) {
	issueID, err := http_utils.GetParamUUID(r, "issue_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	attachments, err := c.AttachmentService.AllByEntity("issue", issueID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, AttachmentSerializer{}.ModelToResponseMany(*attachments))
}

func (c AttachmentController) IndexByProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := http_utils.GetParamUUID(r, "project_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	attachments, err := c.AttachmentService.AllByEntity("project", projectID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http_utils.RespondWithJSON(w, http.StatusOK, AttachmentSerializer{}.ModelToResponseMany(*attachments))
}

func (c AttachmentController) CreateForIssue(w http.ResponseWriter, r *http.Request) {
	issueID, err := http_utils.GetParamUUID(r, "issue_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c.createAttachment(w, r, "issue", issueID)
}

func (c AttachmentController) CreateForProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := http_utils.GetParamUUID(r, "project_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c.createAttachment(w, r, "project", projectID)
}

func (c AttachmentController) createAttachment(w http.ResponseWriter, r *http.Request, entityType string, entityID uuid.UUID) {
	user_id := jwt_utils.GetTokenInfoFromRequest(r).UserID

	r.ParseMultipartForm(32 << 20)
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	mimeType := header.Header.Get("Content-Type")
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	attachment := domain.AttachmentModel{
		EntityType: entityType,
		EntityID:   entityID,
		Filename:   header.Filename,
		MimeType:   mimeType,
		Size:       header.Size,
	}

	if err := c.AttachmentService.Create(&attachment, file); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if entityType == "issue" {
		c.IssueAuditLogService.LogAttachmentAdded(user_id, attachment)
	}

	http_utils.RespondWithJSON(w, http.StatusCreated, AttachmentSerializer{}.ModelToResponse(attachment))
}

func (c AttachmentController) Download(w http.ResponseWriter, r *http.Request) {
	attachmentID, err := http_utils.GetParamUUID(r, "attachment_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	attachment, err := c.AttachmentService.FindByID(attachmentID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=\""+attachment.Filename+"\"")
	w.Header().Set("Content-Type", attachment.MimeType)
	http.ServeFile(w, r, attachment.StorePath)
}

func (c AttachmentController) Destroy(w http.ResponseWriter, r *http.Request) {
	attachmentID, err := http_utils.GetParamUUID(r, "attachment_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user_id := jwt_utils.GetTokenInfoFromRequest(r).UserID

	attachment, err := c.AttachmentService.FindByID(attachmentID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := c.AttachmentService.Destroy(attachmentID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if attachment.EntityType == "issue" {
		c.IssueAuditLogService.LogAttachmentRemoved(user_id, *attachment)
	}

	w.WriteHeader(http.StatusNoContent)
}
