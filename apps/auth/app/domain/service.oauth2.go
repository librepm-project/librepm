package domain

import (
	"log"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
)

type OAuth2ServiceInterface interface {
	Get() *server.Server
}

type OAuth2Service struct {
	ClientRepository ClientRepositoryInterface
}

func userAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	return "000000", nil
}

func internalErrorHandler(err error) (re *errors.Response) {
	log.Println("Internal Error:", err.Error())
	return
}

func responseErrorHandler(re *errors.Response) {
	log.Println("Response Error:", re.Error.Error())
}

func (s OAuth2Service) Get() *server.Server {
	manager := manage.NewDefaultManager()

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	s.ClientRepository.AddToManager(manager)

	srv.UserAuthorizationHandler = userAuthorizationHandler
	srv.SetInternalErrorHandler(internalErrorHandler)
	srv.SetResponseErrorHandler(responseErrorHandler)

	return srv
}
