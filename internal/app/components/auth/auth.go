package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"net/http"
	"os"
	"time"
)

var TokenAuth *jwtauth.JWTAuth

const UserIdClaimKey = "user_id"
const ExpirationClaimKey = "exp"

// refresh token if it expires in 3 days
const RefreshTokenTime = 3 * 24 * time.Hour

func init() {
	TokenAuth = jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
}

func GenerateTokenForUser(id users.UserId) (*jwt.Token, string, error) {
	// token expires in 14 days
	claims := jwt.MapClaims{
		UserIdClaimKey:     id,
		ExpirationClaimKey: time.Now().AddDate(0, 0, 14),
	}
	return TokenAuth.Encode(claims)
}

func SetJwtCookie(tokenString string, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:  "jwt",
		Value: tokenString,
	}
	http.SetCookie(w, &cookie)
}

func ClearJwtCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "",
	}
	http.SetCookie(w, &cookie)
}

// should refresh if token will expire in refresh token time
func ShouldRefreshToken(exp time.Time) bool {
	refreshTime := time.Now().Add(RefreshTokenTime)
	return exp.Before(refreshTime)
}
