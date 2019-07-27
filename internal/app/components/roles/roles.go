package roles

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/permissions"
)

type RoleDisplayName string
type RoleId string

type Role interface {
	// check if role has certain permission level access
	CanAccess(permission permissions.PermissionLevel) bool
	GetRoleId() RoleId
}

func GetRoleFromId(id RoleId) Role {
	switch id {
	case ConradId:
		return Role(NewConrad())
	case DaveId:
		return Role(NewDave())
	case SnewId:
		return Role(NewSnew())
	}
	return Role(NewConrad())
}
