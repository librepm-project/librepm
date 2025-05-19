package domain

type Domain struct {
	OAuth2Service OAuth2ServiceInterface
}

func NewDomain() Domain {
	clientRepository := ClientRepository{}
	oauth2Service := OAuth2Service{
		ClientRepository: clientRepository,
	}
	return Domain{
		OAuth2Service: oauth2Service,
	}
}
