package roles

import "github.com/gzcharleszhang/course-planner/internal/app/components/permissions"

// regular user
type Conrad struct {
	DisplayName string
	Permissions permissions.PermissionLevels
}

func NewConrad() Role {
	return Conrad{
		DisplayName: "user",
		Permissions: permissions.PermissionLevels{
			permissions.Unauthenticated,
			permissions.AuthRequired,
		},
	}
}

func (c Conrad) CanAccess(permission permissions.PermissionLevel) bool {
	return c.Permissions.Contains(permission)
}
