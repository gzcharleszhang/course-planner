package roles

import (
	"context"
	"fmt"
	"github.com/gzcharleszhang/course-planner/internal/app/components/contextKeys"
	"github.com/gzcharleszhang/course-planner/internal/app/components/permissions"
	"github.com/pkg/errors"
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

func GetRoleFromContext(ctx context.Context) (Role, error) {
	ctxRole := ctx.Value(contextKeys.UserRoleKey)
	role, ok := ctxRole.(Role)
	if !ok {
		return nil, errors.New(fmt.Sprintf("cannot convert %v to a role", ctxRole))
	}
	return role, nil
}
