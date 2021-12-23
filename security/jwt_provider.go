package security

import (
	"halill/ent"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTProvider interface {
	GenerateAccessToken(*ent.User) (string, error)
	GenerateRefreshToken(*ent.User) (string, error)
	JwtSecret() string
}

type jwtProvider struct {
	jwtSecret string
}

func NewJWTProvider(jwtSecret string) JWTProvider {
	return &jwtProvider{jwtSecret}
}

func (j *jwtProvider) GenerateAccessToken(user *ent.User) (string, error) {
	accessTokenClaims := &JwtCustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	at, err := accessToken.SignedString([]byte(j.jwtSecret))
	if err != nil {
		return "", err
	}

	return at, nil
}

func (j *jwtProvider) GenerateRefreshToken(user *ent.User) (string, error) {
	accessTokenClaims := &JwtCustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 14).Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	at, err := accessToken.SignedString([]byte(j.jwtSecret))
	if err != nil {
		return "", err
	}

	return at, nil
}

func (j *jwtProvider) JwtSecret() string {
	return j.jwtSecret
}
