package mocksvc

import (
	"apps/api/app/domain"
)

type MockSettingService struct {
	GetAllFunc    func() (*[]domain.SettingModel, error)
	GetByKeyFunc  func(key string) (*domain.SettingModel, error)
	UpdateFunc    func(key string, value string) error
}

func (m *MockSettingService) GetAll() (*[]domain.SettingModel, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc()
	}
	return &[]domain.SettingModel{}, nil
}

func (m *MockSettingService) GetByKey(key string) (*domain.SettingModel, error) {
	if m.GetByKeyFunc != nil {
		return m.GetByKeyFunc(key)
	}
	return &domain.SettingModel{}, nil
}

func (m *MockSettingService) Update(key string, value string) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(key, value)
	}
	return nil
}
