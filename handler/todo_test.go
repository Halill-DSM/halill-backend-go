package handler

import (
	"bytes"
	"encoding/json"
	"halill/dto"
	"halill/ent"
	"halill/mocks"
	"halill/security"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllTodos(t *testing.T) {
	viper.SetConfigFile("./../config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	e := echo.New()
	g := e.Group("/todo")
	ts := new(mocks.TodoService)
	user := &ent.User{
		ID:       "hwc9169@gmail.com",
		Password: "password",
		Name:     "조호원",
	}
	deadline := time.Now().Add(24 * 3 * time.Hour)
	expectedResponse := []*dto.TodoResponse{
		{
			ID:          1,
			Title:       "Go 언어 공부하기",
			Content:     "장재휴의 Go 웹 프로그래밍 철저 입문",
			Deadline:    &deadline,
			IsCompleted: false,
		}, {
			ID:          2,
			Title:       "CEO가 해야하는 일",
			Content:     "역대 CEO가 공부하는 것들 찾아보기",
			IsCompleted: false,
		},
	}
	ts.On("GetAllTodos", mock.AnythingOfType("string")).Return(expectedResponse, nil)

	t.Run("모든 Todo 요청 성공", func(t *testing.T) {
		jwtProvider := security.NewJWTProvider(viper.GetString("jwt.secret"))
		accessToken, err := jwtProvider.GenerateAccessToken(user)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/todo", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+accessToken)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		th := NewTodoHandler(g, ts, jwtProvider.JwtSecret())
		err = middleware.JWTWithConfig(middleware.JWTConfig{
			Claims:     &security.JwtCustomClaims{},
			SigningKey: []byte(jwtProvider.JwtSecret()),
		})(th.GetAllTodos)(c)
		assert.NoError(t, err)
	})
}

func TestGetTodo(t *testing.T) {
	e := echo.New()
	g := e.Group("/todo")
	ts := new(mocks.TodoService)
	deadline := time.Now().Add(24 * 3 * time.Hour)
	user := &ent.User{
		ID:       "hwc9169@gmail.com",
		Password: "password",
		Name:     "조호원",
	}
	expectedResponse := &dto.TodoResponse{
		ID:          1,
		Title:       "Go 언어 공부하기",
		Content:     "장재휴의 Go 웹 프로그래밍 철저 입문",
		Deadline:    &deadline,
		IsCompleted: false,
	}
	ts.On("GetTodo", mock.AnythingOfType("int64"), mock.AnythingOfType("string")).Return(expectedResponse, nil)

	t.Run("Todo 요청 성공", func(t *testing.T) {
		jwtProvider := security.NewJWTProvider(viper.GetString("jwt.secret"))
		accessToken, err := jwtProvider.GenerateAccessToken(user)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/todo/1", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+accessToken)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:todo_id")
		c.SetParamNames("todo_id")
		c.SetParamValues("1")

		th := NewTodoHandler(g, ts, jwtProvider.JwtSecret())
		err = middleware.JWTWithConfig(middleware.JWTConfig{
			Claims:     &security.JwtCustomClaims{},
			SigningKey: []byte(jwtProvider.JwtSecret()),
		})(th.GetTodo)(c)
		assert.NoError(t, err)
	})
}

func TestCreateTodo(t *testing.T) {
	e := echo.New()
	g := e.Group("/todo")
	ts := new(mocks.TodoService)
	deadline := time.Now().Add(24 * 3 * time.Hour)
	user := &ent.User{
		ID:       "hwc9169@gmail.com",
		Password: "password",
		Name:     "조호원",
	}
	expectedResponse := &dto.TodoResponse{
		ID:          1,
		Title:       "Go 언어 공부하기",
		Content:     "장재휴의 Go 웹 프로그래밍 철저 입문",
		Deadline:    &deadline,
		IsCompleted: false,
	}
	ts.On("CreateTodo", mock.AnythingOfType("*dto.CreateTodoRequest"), mock.AnythingOfType("string")).Return(expectedResponse, nil)

	t.Run("Todo 생성 요청 성공", func(t *testing.T) {
		jwtProvider := security.NewJWTProvider(viper.GetString("jwt.secret"))
		accessToken, err := jwtProvider.GenerateAccessToken(user)
		assert.NoError(t, err)

		createTodoRequest := &dto.CreateTodoRequest{
			Title:    "Go 언어 공부하기",
			Content:  "장재휴의 Go 웹 프로그래밍 철저 입문",
			Deadline: &deadline,
		}
		request := &bytes.Buffer{}
		json.NewEncoder(request).Encode(createTodoRequest)

		req := httptest.NewRequest(http.MethodPost, "/todo", request)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+accessToken)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		th := NewTodoHandler(g, ts, jwtProvider.JwtSecret())
		err = middleware.JWTWithConfig(middleware.JWTConfig{
			Claims:     &security.JwtCustomClaims{},
			SigningKey: []byte(jwtProvider.JwtSecret()),
		})(th.CreateTodo)(c)
		assert.NoError(t, err)
	})
}

func TestCompleteTodo(t *testing.T) {
	e := echo.New()
	g := e.Group("/todo")
	ts := new(mocks.TodoService)
	deadline := time.Now().Add(24 * 3 * time.Hour)
	user := &ent.User{
		ID:       "hwc9169@gmail.com",
		Password: "password",
		Name:     "조호원",
	}
	expectedResponse := &dto.TodoResponse{
		ID:          1,
		Title:       "Go 언어 공부하기",
		Content:     "장재휴의 Go 웹 프로그래밍 철저 입문",
		Deadline:    &deadline,
		IsCompleted: true,
	}
	ts.On("CompleteTodo", mock.AnythingOfType("int64"), mock.AnythingOfType("string")).Return(expectedResponse, nil)

	t.Run("Todo 완료 요청 성공", func(t *testing.T) {
		jwtProvider := security.NewJWTProvider(viper.GetString("jwt.secret"))
		accessToken, err := jwtProvider.GenerateAccessToken(user)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPatch, "/todo/1", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+accessToken)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:todo_id")
		c.SetParamNames("todo_id")
		c.SetParamValues("1")

		th := NewTodoHandler(g, ts, jwtProvider.JwtSecret())
		err = middleware.JWTWithConfig(middleware.JWTConfig{
			Claims:     &security.JwtCustomClaims{},
			SigningKey: []byte(jwtProvider.JwtSecret()),
		})(th.CompleteTodo)(c)
		assert.NoError(t, err)
	})
}

func TestDeleteTodo(t *testing.T) {
	e := echo.New()
	g := e.Group("/todo")
	ts := new(mocks.TodoService)
	deadline := time.Now().Add(24 * 3 * time.Hour)
	user := &ent.User{
		ID:       "hwc9169@gmail.com",
		Password: "password",
		Name:     "조호원",
	}
	expectedResponse := &dto.TodoResponse{
		ID:          1,
		Title:       "Go 언어 공부하기",
		Content:     "장재휴의 Go 웹 프로그래밍 철저 입문",
		Deadline:    &deadline,
		IsCompleted: true,
	}
	ts.On("DeleteTodo", mock.AnythingOfType("int64"), mock.AnythingOfType("string")).Return(expectedResponse, nil)

	t.Run("Todo 삭제 요청 성공", func(t *testing.T) {
		jwtProvider := security.NewJWTProvider(viper.GetString("jwt.secret"))
		accessToken, err := jwtProvider.GenerateAccessToken(user)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodDelete, "/todo/1", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+accessToken)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:todo_id")
		c.SetParamNames("todo_id")
		c.SetParamValues("1")

		th := NewTodoHandler(g, ts, jwtProvider.JwtSecret())
		err = middleware.JWTWithConfig(middleware.JWTConfig{
			Claims:     &security.JwtCustomClaims{},
			SigningKey: []byte(jwtProvider.JwtSecret()),
		})(th.DeleteTodo)(c)
		assert.NoError(t, err)
	})
}
