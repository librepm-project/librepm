package domain

type SettingServiceInterface interface {
	GetAll() (*[]SettingModel, error)
	GetByKey(key string) (*SettingModel, error)
	Update(key string, value string) error
}

type SettingService struct {
	SettingRepository SettingRepositoryInterface
}

func (s SettingService) GetAll() (*[]SettingModel, error) {
	return s.SettingRepository.All()
}

func (s SettingService) GetByKey(key string) (*SettingModel, error) {
	return s.SettingRepository.FindByKey(key)
}

func (s SettingService) Update(key string, value string) error {
	return s.SettingRepository.Update(key, value)
}
