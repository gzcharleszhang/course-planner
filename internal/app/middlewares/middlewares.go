package middlewares

import (
	"context"
	"github.com/go-chi/jwtauth"
	"github.com/gzcharleszhang/course-planner/internal/app/components/auth"
	"github.com/gzcharleszhang/course-planner/internal/app/components/contextKeys"
	"github.com/gzcharleszhang/course-planner/internal/app/components/permissions"
	"github.com/gzcharleszhang/course-planner/internal/app/components/roles"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"github.com/gzcharleszhang/course-planner/internal/app/models/userModel"
	"net/http"
	"time"
)

func PermissionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// default to unauthenticated
		ctx := r.Context()
		token, claims, err := jwtauth.FromContext(ctx)
		// extract permission access level from token
		if err == nil && token != nil && token.Valid {
			userId, ok := claims[auth.UserIdClaimKey].(users.UserId)
			if !ok {
				http.Error(w, http.StatusText(401), 401)
				return
			}
			// check if token is expired
			expirationTime, ok := claims[auth.ExpirationClaimKey].(time.Time)
			if !ok || expirationTime.Before(time.Now()) {
				http.Error(w, http.StatusText(401), 401)
				return
			}
			role, err := userModel.GetUserRole(ctx, userId)
			if err != nil {
				http.Error(w, http.StatusText(401), 401)
				return
			}
			// set role field in the context
			if role != nil {
				ctx = context.WithValue(ctx, contextKeys.UserRoleKey, *role)
			}
			// set user id field in the context
			ctx = context.WithValue(ctx, contextKeys.UserIdKey, userId)
			// refresh token if needed
			if auth.ShouldRefreshToken(expirationTime) {
				_, token, err := auth.GenerateTokenForUser(userId)
				if err != nil {
					http.Error(w, http.StatusText(401), 401)
					return
				}
				auth.SetJwtCookie(token, w)
			}
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// for protecting admin routes
func VerifyAdminMiddleware(next http.Handler) http.Handler {
	return AuthMiddlewareFactory(permissions.AdminRequired)(next)
}

// for protecting authenticated routes
func VerifyAuthenticatedMiddleware(next http.Handler) http.Handler {
	return AuthMiddlewareFactory(permissions.AuthRequired)(next)
}

// returns a middleware that checks if request has the permission level required
func AuthMiddlewareFactory(permRequired permissions.PermissionLevel) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			role, ok := ctx.Value(contextKeys.UserRoleKey).(roles.Role)
			// reject if no role is defined and permRequired is not unauthenticated
			if !ok && permRequired != permissions.Unauthenticated {
				http.Error(w, http.StatusText(401), 401)
				return
			}
			// reject if given role does not have permission level access required
			if role.CanAccess(permRequired) {
				http.Error(w, http.StatusText(401), 401)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
