package roles

import "github.com/gzcharleszhang/course-planner/internal/app/components/permissions"

// regular user
type Snew struct {
	DisplayName string
	Permissions permissions.PermissionLevels
}

func NewSnew() Role {
	return Snew{
		DisplayName: "admin",
		Permissions: permissions.PermissionLevels{
			permissions.Unauthenticated,
			permissions.AuthRequired,
			permissions.AdminRequired,
			permissions.SuperAdminRequired,
		},
	}
}

func (s Snew) CanAccess(permission permissions.PermissionLevel) bool {
	return s.Permissions.Contains(permission)
}
