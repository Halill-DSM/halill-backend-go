package handler

import (
	"bytes"
	"encoding/json"
	"halill/dto"
	"halill/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	e := echo.New()
	g := e.Group("")
	us := new(mocks.UserService)
	expectedResponse := &dto.TokenResponse{
		AccessToken:  "test-access-token",
		RefreshToken: "test-refresh-token",
	}
	us.On("LoginUser", mock.AnythingOfType("*dto.LoginRequest")).Return(expectedResponse, nil)

	t.Run("로그인 요청 성공", func(t *testing.T) {
		uh := NewUserHandler(g, us)
		loginRequest := &dto.LoginRequest{
			Email:    "hwc9169@gmail.com",
			Password: "password",
		}
		request := &bytes.Buffer{}
		json.NewEncoder(request).Encode(loginRequest)
		req := httptest.NewRequest(http.MethodPost, "/login", request)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := uh.Login(c)
		assert.NoError(t, err)
	})
}

func TestRegist(t *testing.T) {
	e := echo.New()
	g := e.Group("")
	us := new(mocks.UserService)
	expectedResponse := &dto.UserResponse{
		Name:  "조호원",
		Email: "hwc9169@gmail.com",
	}
	us.On("RegistUser", mock.AnythingOfType("*dto.RegistRequest")).Return(expectedResponse, nil)

	t.Run("회원가입 요청 성공", func(t *testing.T) {
		uh := NewUserHandler(g, us)
		registRequest := &dto.RegistRequest{
			Email:    "hwc9169@gmail.com",
			Password: "password",
			Name:     "조호원",
		}
		request := &bytes.Buffer{}
		json.NewEncoder(request).Encode(registRequest)
		req := httptest.NewRequest(http.MethodPost, "/signup", request)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := uh.Register(c)
		assert.NoError(t, err)
	})
}

func TestRefresh(t *testing.T) {
	e := echo.New()
	g := e.Group("")
	us := new(mocks.UserService)
	expectedResponse := &dto.TokenResponse{
		AccessToken:  "asdf.asdf.asdf",
		RefreshToken: "asdf.asdf.asdf",
	}
	us.On("RefreshToken", mock.AnythingOfType("*dto.RefreshTokenRequest")).Return(expectedResponse, nil)

	t.Run("토큰 요청 성공", func(t *testing.T) {
		uh := NewUserHandler(g, us)
		registRequest := &dto.RefreshTokenRequest{
			RefreshToken: "asdf.asdf.asdf",
		}
		request := &bytes.Buffer{}
		json.NewEncoder(request).Encode(registRequest)
		req := httptest.NewRequest(http.MethodPut, "/refresh", request)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := uh.Refresh(c)
		assert.NoError(t, err)
	})
}
