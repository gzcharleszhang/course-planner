package roles

import "github.com/gzcharleszhang/course-planner/internal/app/components/permissions"

const DaveDisplayName = "admin"
const DaveId = "admin"

// regular user
type Dave struct {
	Id          RoleId                       `json:"_id",bson:"_id"`
	DisplayName string                       `json:"display_name",bson:"display_name"`
	Permissions permissions.PermissionLevels `json:"permissions",bson:"permissions"`
}

func NewDave() Role {
	return Dave{
		Id:          DaveId,
		DisplayName: DaveDisplayName,
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

func (d Dave) GetRoleId() RoleId {
	return d.Id
}
