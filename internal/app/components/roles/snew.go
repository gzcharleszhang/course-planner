package roles

import "github.com/gzcharleszhang/course-planner/internal/app/components/permissions"

const SnewDisplayName = "super_admin"
const SnewId = "suer_admin"

// regular user
type Snew struct {
	Id          RoleId                       `json:"_id",bson:"_id"`
	DisplayName string                       `json:"display_name",bson:"display_name"`
	Permissions permissions.PermissionLevels `json:"permissions",bson:"permissions"`
}

func NewSnew() Role {
	return Snew{
		Id:          SnewId,
		DisplayName: SnewDisplayName,
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

func (s Snew) GetRoleId() RoleId {
	return s.Id
}
