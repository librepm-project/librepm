package http

import (
	"net/http"

	"apps/api/app/domain"
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
	issueID, err := GetParamUUID(r, "issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	attachments, err := c.AttachmentService.AllByEntity(domain.EntityTypeIssue, issueID)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, AttachmentSerializer{}.ModelToResponseMany(*attachments))
}

func (c AttachmentController) IndexByProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := GetParamUUID(r, "project_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	attachments, err := c.AttachmentService.AllByEntity(domain.EntityTypeProject, projectID)
	if err != nil {
		RespondBadRequest(w)
		return
	}
	RespondJSON(w, http.StatusOK, AttachmentSerializer{}.ModelToResponseMany(*attachments))
}

func (c AttachmentController) CreateForIssue(w http.ResponseWriter, r *http.Request) {
	issueID, err := GetParamUUID(r, "issue_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	c.createAttachment(w, r, domain.EntityTypeIssue, issueID)
}

func (c AttachmentController) CreateForProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := GetParamUUID(r, "project_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	c.createAttachment(w, r, domain.EntityTypeProject, projectID)
}

func (c AttachmentController) createAttachment(w http.ResponseWriter, r *http.Request, entityType string, entityID uuid.UUID) {
	user_id := GetUserIDFromRequest(r)

	r.ParseMultipartForm(32 << 20)
	file, header, err := r.FormFile("file")
	if err != nil {
		RespondBadRequest(w)
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
		RespondInternalError(w)
		return
	}

	if entityType == domain.EntityTypeIssue {
		c.IssueAuditLogService.LogAttachmentAdded(user_id, attachment)
	}

	RespondJSON(w, http.StatusCreated, AttachmentSerializer{}.ModelToResponse(attachment))
}

func (c AttachmentController) Download(w http.ResponseWriter, r *http.Request) {
	attachmentID, err := GetParamUUID(r, "attachment_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}
	attachment, err := c.AttachmentService.FindByID(attachmentID)
	if err != nil {
		RespondNotFound(w)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=\""+attachment.Filename+"\"")
	w.Header().Set("Content-Type", attachment.MimeType)
	http.ServeFile(w, r, attachment.StorePath)
}

func (c AttachmentController) Destroy(w http.ResponseWriter, r *http.Request) {
	attachmentID, err := GetParamUUID(r, "attachment_id")
	if err != nil {
		RespondBadRequest(w)
		return
	}

	user_id := GetUserIDFromRequest(r)

	attachment, err := c.AttachmentService.FindByID(attachmentID)
	if err != nil {
		RespondNotFound(w)
		return
	}

	if err := c.AttachmentService.Destroy(attachmentID); err != nil {
		RespondBadRequest(w)
		return
	}

	if attachment.EntityType == domain.EntityTypeIssue {
		c.IssueAuditLogService.LogAttachmentRemoved(user_id, *attachment)
	}

	RespondNoContent(w)
}
