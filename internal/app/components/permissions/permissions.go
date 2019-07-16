package permissions

type Permission string

const (
	Admin           = "admin"
	Unauthenticated = "unauthenticated"
	Authenticated   = "authenticated"
)
