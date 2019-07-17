package middlewares

import (
	"context"
	"github.com/go-chi/jwtauth"
	"github.com/gzcharleszhang/course-planner/internal/app/components/auth"
	"github.com/gzcharleszhang/course-planner/internal/app/components/contextKeys"
	"github.com/gzcharleszhang/course-planner/internal/app/components/permissions"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"net/http"
	"time"
)

func PermissionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// default to unauthenticated
		ctx := context.WithValue(r.Context(), contextKeys.PermissionAccessKey, permissions.Unauthenticated)
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
			perm, err := users.GetUserPermissionAccess(ctx, userId)
			if err != nil {
				http.Error(w, http.StatusText(401), 401)
				return
			}
			// set permission access field in the context
			if perm != nil {
				ctx = context.WithValue(ctx, contextKeys.PermissionAccessKey, *perm)
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		perm := ctx.Value(contextKeys.PermissionAccessKey)
		if perm != permissions.Admin {
			http.Error(w, http.StatusText(401), 401)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// for protecting authenticated routes
func VerifyAuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		perm := ctx.Value(contextKeys.PermissionAccessKey)
		if perm == permissions.Unauthenticated {
			http.Error(w, http.StatusText(401), 401)
			return
		}
		next.ServeHTTP(w, r)
	})
}