package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"os"
)

var TokenAuth *jwtauth.JWTAuth

const UserIdClaimKey = "user_id"

func init() {
	TokenAuth = jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
}

func GenerateTokenForUser(id users.UserId) (*jwt.Token, string, error) {
	return TokenAuth.Encode(jwt.MapClaims{UserIdClaimKey: id})
}
