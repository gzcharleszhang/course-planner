package permissions

type PermissionLevel string
type PermissionLevels []PermissionLevel

// each level responds to a set of web services/privilege that a role has access to
const (
	SuperAdminRequired PermissionLevel = "super_admin"
	AdminRequired      PermissionLevel = "admin"
	Unauthenticated    PermissionLevel = "unauthenticated"
	AuthRequired       PermissionLevel = "authenticated"
)

// whether or not a slice of permission levels contains certain permission
func (perms PermissionLevels) Contains(perm PermissionLevel) bool {
	permMap := map[PermissionLevel]bool{}
	for _, p := range perms {
		permMap[p] = true
	}
	_, ok := permMap[perm]
	return ok
}
