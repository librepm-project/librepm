package mocksvc

type MockOnboardService struct {
	ExecuteFunc func(siteTitle string, email string, password string, firstName string, lastName string) error
}

func (m *MockOnboardService) Execute(siteTitle string, email string, password string, firstName string, lastName string) error {
	if m.ExecuteFunc != nil {
		return m.ExecuteFunc(siteTitle, email, password, firstName, lastName)
	}
	return nil
}
