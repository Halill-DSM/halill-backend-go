package security

import (
	"halill/ent"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestGenerateAccessToken(t *testing.T) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	assert.NoError(t, err)
	user := &ent.User{
		ID:       "hwc9169@gmail.com",
		Password: string(hashedPassword),
		Name:     "조호원",
	}

	t.Run("access token 생성 성공", func(t *testing.T) {
		jp := NewJWTProvider("test_secret")
		resp, err := jp.GenerateAccessToken(user)
		assert.NoError(t, err)

		claims := &JwtCustomClaims{}
		jwt.ParseWithClaims(resp, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jp.JwtSecret()), nil
		})
		assert.Equal(t, user.ID, claims.Email)
	})
}

func TestGenerateRefreshTOkdn(t *testing.T) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	assert.NoError(t, err)
	user := &ent.User{
		ID:       "hwc9169@gmail.com",
		Password: string(hashedPassword),
		Name:     "조호원",
	}

	t.Run("refresh token 생성 성공", func(t *testing.T) {
		jp := NewJWTProvider("test_secret")
		resp, err := jp.GenerateRefreshToken(user)
		assert.NoError(t, err)

		claims := &JwtCustomClaims{}
		jwt.ParseWithClaims(resp, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jp.JwtSecret()), nil
		})
		assert.Equal(t, user.ID, claims.Email)
	})
}
