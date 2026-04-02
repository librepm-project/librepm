package seed

import "apps/api/app/domain"

func (s SeedService) createUserRole(items []UserRoleData) error {
	var err error
	for _, item := range items {
		user, findErr := s.UserRepository.FindByEmail(item.UserEmail)
		if findErr != nil {
			return findErr
		}
		m := &domain.UserRoleModel{
			UserID: user.ID,
			Role:   item.Role,
		}
		err = s.UserRoleRepository.Create(m)
	}
	return err
}
