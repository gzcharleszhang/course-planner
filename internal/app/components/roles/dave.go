package roles

import "github.com/gzcharleszhang/course-planner/internal/app/components/permissions"

// regular user
type Dave struct {
	DisplayName string
	Permissions permissions.PermissionLevels
}

func NewDave() Role {
	return Dave{
		DisplayName: "admin",
		Permissions: permissions.PermissionLevels{
			permissions.Unauthenticated,
			permissions.AuthRequired,
			permissions.AdminRequired,
		},
	}
}

func (d Dave) CanAccess(permission permissions.PermissionLevel) bool {
	return d.Permissions.Contains(permission)
}
