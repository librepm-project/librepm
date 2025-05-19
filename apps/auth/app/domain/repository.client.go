package domain

import (
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/store"
)

type ClientRepositoryInterface interface {
	AddToManager(manager *manage.Manager)
}

type ClientRepository struct {
}

func (c ClientRepository) AddToManager(manager *manage.Manager) {
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost:8081/auth",
	})
	manager.MapClientStorage(clientStore)
}
