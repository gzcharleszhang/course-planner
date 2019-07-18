package roles

import "github.com/gzcharleszhang/course-planner/internal/app/components/permissions"

type Role interface {
	// check if role has certain permission level access
	CanAccess(permission permissions.PermissionLevel) bool
}
