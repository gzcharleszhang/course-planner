package roles

import "github.com/gzcharleszhang/course-planner/internal/app/components/permissions"

const ConradDisplayName = "user"
const ConradId = "user"

// regular user
type Conrad struct {
	Id          RoleId                       `json:"_id",bson:"_id"`
	DisplayName string                       `json:"display_name",bson:"display_name"`
	Permissions permissions.PermissionLevels `json:"permissions",bson:"permissions"`
}

func NewConrad() Role {
	return Conrad{
		Id:          ConradId,
		DisplayName: ConradDisplayName,
		Permissions: permissions.PermissionLevels{
			permissions.Unauthenticated,
			permissions.AuthRequired,
		},
	}
}

func (c Conrad) CanAccess(permission permissions.PermissionLevel) bool {
	return c.Permissions.Contains(permission)
}

func (c Conrad) GetRoleId() RoleId {
	return c.Id
}
