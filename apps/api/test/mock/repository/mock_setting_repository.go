package mockrepo

import "apps/api/app/domain"

type MockSettingRepository struct {
	AllFunc        func() (*[]domain.SettingModel, error)
	FindByKeyFunc  func(key string) (*domain.SettingModel, error)
	UpdateFunc     func(key string, value string) error
	SeedFunc       func() error
}

func (m *MockSettingRepository) All() (*[]domain.SettingModel, error) {
	if m.AllFunc != nil {
		return m.AllFunc()
	}
	return &[]domain.SettingModel{}, nil
}

func (m *MockSettingRepository) FindByKey(key string) (*domain.SettingModel, error) {
	if m.FindByKeyFunc != nil {
		return m.FindByKeyFunc(key)
	}
	return &domain.SettingModel{}, nil
}

func (m *MockSettingRepository) Update(key string, value string) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(key, value)
	}
	return nil
}

func (m *MockSettingRepository) Seed() error {
	if m.SeedFunc != nil {
		return m.SeedFunc()
	}
	return nil
}
