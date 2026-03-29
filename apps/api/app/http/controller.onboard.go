package http

import (
	"apps/api/app/domain"
	"net/http"
)

type OnboardControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type OnboardController struct {
	OnboardService domain.OnboardServiceInterface
}

type OnboardRequest struct {
	SiteTitle string `json:"siteTitle"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (c OnboardController) Create(w http.ResponseWriter, r *http.Request) {
	var body OnboardRequest
	if err := DecodeJSON(r, &body); err != nil {
		RespondBadRequest(w)
		return
	}
	if err := c.OnboardService.Execute(body.SiteTitle, body.Email, body.Password, body.FirstName, body.LastName); err != nil {
		RespondBadRequest(w)
		return
	}
	RespondNoContent(w)
}
