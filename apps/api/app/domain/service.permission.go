package domain

import "github.com/google/uuid"

type PermissionServiceInterface interface {
	UserCan(userID uuid.UUID, perm Permission) bool
	UserPermissions(userID uuid.UUID) []string
}

type PermissionService struct {
	UserRoleRepository UserRoleRepositoryInterface
}

func (s PermissionService) UserCan(userID uuid.UUID, perm Permission) bool {
	roles, err := s.UserRoleRepository.FindByUserID(userID)
	if err != nil || len(roles) == 0 {
		return false
	}
	for _, r := range roles {
		if RoleHasPermission(r.Role, perm) {
			return true
		}
	}
	return false
}

func (s PermissionService) UserPermissions(userID uuid.UUID) []string {
	roles, err := s.UserRoleRepository.FindByUserID(userID)
	if err != nil {
		return []string{}
	}
	seen := map[Permission]bool{}
	result := []string{}
	for _, r := range roles {
		role, ok := Roles[r.Role]
		if !ok {
			continue
		}
		for _, p := range role.Permissions {
			if !seen[p] {
				seen[p] = true
				result = append(result, string(p))
			}
		}
	}
	return result
}
