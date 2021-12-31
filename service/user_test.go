package service

import (
	"halill/dto"
	"halill/ent"
	"halill/mocks"
	"halill/security"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginUser(t *testing.T) {
	ur := new(mocks.UserRepository)
	jp := new(mocks.JWTProvider)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	assert.NoError(t, err)
	user := &ent.User{
		ID:       "hwc9169@gmail.com",
		Password: string(hashedPassword),
		Name:     "조호원",
	}
	t.Run("User 로그인 성공", func(t *testing.T) {
		expectedResponse := &dto.TokenResponse{
			AccessToken:  "asdf.asdf.asdf",
			RefreshToken: "asdf.asdf.asdf",
		}
		ur.On("GetByEmail", mock.AnythingOfType("string")).Return(user, nil)
		jp.On("GenerateAccessToken", mock.AnythingOfType("*ent.User")).Return("asdf.asdf.asdf", nil)
		jp.On("GenerateRefreshToken", mock.AnythingOfType("*ent.User")).Return("asdf.asdf.asdf", nil)
		us := NewUserSerice(ur, jp)

		resp, err := us.LoginUser(&dto.LoginRequest{
			Email:    "hwc9169@gmail.com",
			Password: "password",
		})
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, resp)
	})
	t.Run("User 로그인 실패", func(t *testing.T) {
		ur.On("GetByEmail", mock.AnythingOfType("string")).Return(user, nil)
		jp.On("GenerateAccessToken", mock.AnythingOfType("*ent.User")).Return("asdf.asdf.asdf", nil)
		jp.On("GenerateRefreshToken", mock.AnythingOfType("*ent.User")).Return("asdf.asdf.asdf", nil)
		us := NewUserSerice(ur, jp)

		_, err := us.LoginUser(&dto.LoginRequest{
			Email:    "hwc9169@gmail.com",
			Password: "different_password",
		})
		assert.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
	})
}

func TestRegistUser(t *testing.T) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	assert.NoError(t, err)
	user := &ent.User{
		ID:       "hwc9169@gmail.com",
		Password: string(hashedPassword),
		Name:     "조호원",
	}

	t.Run("유저 회원가입 성공", func(t *testing.T) {
		expectedResponse := &dto.UserResponse{
			Email: "hwc9169@gmail.com",
			Name:  "조호원",
		}
		ur := new(mocks.UserRepository)
		jp := new(mocks.JWTProvider)
		ur.On("GetByEmail", mock.AnythingOfType("string")).Return(nil, &ent.NotFoundError{})
		ur.On("CreateUser", mock.AnythingOfType("*ent.User")).Return(user, nil)
		us := NewUserSerice(ur, jp)

		resp, err := us.RegistUser(&dto.RegistRequest{
			Email:    "hwc9169@gmail.com",
			Password: "password",
			Name:     "조호원",
		})

		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, resp)
	})

	t.Run("이미 사용중인 이메일일 때", func(t *testing.T) {
		ur := new(mocks.UserRepository)
		jp := new(mocks.JWTProvider)
		ur.On("GetByEmail", mock.AnythingOfType("string")).Return(user, nil)
		us := NewUserSerice(ur, jp)

		_, err := us.RegistUser(&dto.RegistRequest{
			Email:    "hwc9169@gmail.com",
			Password: "password",
			Name:     "조호원",
		})
		assert.EqualError(t, err, echo.NewHTTPError(http.StatusBadRequest, "이미 사용중인 이메일입니다.").Error())
	})
}

func TestRefreshToken(t *testing.T) {
	ur := new(mocks.UserRepository)
	jp := new(mocks.JWTProvider)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	assert.NoError(t, err)
	t.Run("User 토큰 리프레시 성공", func(t *testing.T) {
		user := &ent.User{
			ID:       "hwc9169@gmail.com",
			Password: string(hashedPassword),
			Name:     "조호원",
		}

		refreshToken, err := security.NewJWTProvider("test_secret").GenerateRefreshToken(user)
		assert.NoError(t, err)

		expectedResponse := &dto.TokenResponse{
			AccessToken:  "asdf.asdf.asdf",
			RefreshToken: refreshToken,
		}
		ur.On("GetByEmail", mock.AnythingOfType("string")).Return(user, nil)
		jp.On("JwtSecret").Return("test_secret")
		jp.On("GenerateAccessToken", mock.AnythingOfType("*ent.User")).Return("asdf.asdf.asdf", nil)
		us := NewUserSerice(ur, jp)

		resp, err := us.RefreshToken(&dto.RefreshTokenRequest{
			RefreshToken: refreshToken,
		})
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, resp)
	})
}
