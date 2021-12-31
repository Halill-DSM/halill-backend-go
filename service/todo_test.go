package service

import (
	"halill/dto"
	"halill/ent"
	"halill/mocks"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllTodos(t *testing.T) {
	tr := new(mocks.TodoRepository)

	t.Run("전체 Todo 조회 성공", func(t *testing.T) {
		deadline := time.Now().Add(24 * 3 * time.Hour)
		expectedResponse := []*ent.Todo{
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
		tr.On("GetAllByEmail", mock.AnythingOfType("string")).Return(expectedResponse, nil)
		ts := NewTodoService(tr)

		email := "hwc9169@gmail.com"
		resp, err := ts.GetAllTodos(email)
		assert.NoError(t, err)

		expected := make([]*dto.TodoResponse, 0)
		for _, todo := range expectedResponse {
			expected = append(expected, dto.TodoToDTO(todo))
		}
		assert.Equal(t, expected, resp)
	})
}

func TestGetTodo(t *testing.T) {
	tr := new(mocks.TodoRepository)
	user := &ent.User{
		ID:       "hwc9169@gmail.com",
		Password: "password",
		Name:     "조호원",
	}
	deadline := time.Now().Add(24 * 3 * time.Hour)
	expectedResponse := &ent.Todo{
		ID:          1,
		Title:       "Go 언어 공부하기",
		Content:     "장재휴의 Go 웹 프로그래밍 철저 입문",
		Deadline:    &deadline,
		IsCompleted: false,
		Edges: ent.TodoEdges{
			User: user,
		},
	}
	t.Run("Todo 조회 성공", func(t *testing.T) {
		tr.On("Get", mock.AnythingOfType("int64")).Return(expectedResponse, nil)
		ts := NewTodoService(tr)

		todoID := int64(1)
		email := "hwc9169@gmail.com"
		resp, err := ts.GetTodo(todoID, email)
		assert.NoError(t, err)

		expected := dto.TodoToDTO(expectedResponse)
		assert.Equal(t, expected, resp)
	})
	t.Run("권한 오류 발생", func(t *testing.T) {
		tr.On("Get", mock.AnythingOfType("int64")).Return(expectedResponse, nil)
		ts := NewTodoService(tr)

		todoID := int64(1)
		email := "hwc9169@naver.com"
		_, err := ts.GetTodo(todoID, email)
		assert.Equal(t, err, echo.NewHTTPError(http.StatusForbidden, "해당 요청에 대한 권한이 없습니다."))
	})
}

func TestCreateTodo(t *testing.T) {
	tr := new(mocks.TodoRepository)

	t.Run("Todo 생성 성공", func(t *testing.T) {
		deadline := time.Now().Add(24 * 3 * time.Hour)
		expectedResponse := &ent.Todo{
			ID:          1,
			Title:       "Go 언어 공부하기",
			Content:     "장재휴의 Go 웹 프로그래밍 철저 입문",
			Deadline:    &deadline,
			IsCompleted: false,
		}
		tr.On("Create", mock.AnythingOfType("*ent.Todo")).Return(expectedResponse, nil)
		ts := NewTodoService(tr)

		email := "hwc9169@gmail.com"
		resp, err := ts.CreateTodo(&dto.CreateTodoRequest{
			Title:    "Go 언어 공부하기",
			Content:  "장재휴의 Go 웹 프로그래밍 철저 입문",
			Deadline: &deadline,
		}, email)
		assert.NoError(t, err)

		expected := dto.TodoToDTO(expectedResponse)
		assert.Equal(t, expected, resp)
	})
}

func TestCompleteTodo(t *testing.T) {
	tr := new(mocks.TodoRepository)
	user := &ent.User{
		ID:       "hwc9169@gmail.com",
		Password: "password",
		Name:     "조호원",
	}
	deadline := time.Now().Add(24 * 3 * time.Hour)
	expectedResponse := &ent.Todo{
		ID:          1,
		Title:       "Go 언어 공부하기",
		Content:     "장재휴의 Go 웹 프로그래밍 철저 입문",
		Deadline:    &deadline,
		IsCompleted: false,
		Edges: ent.TodoEdges{
			User: user,
		},
	}
	t.Run("Todo 생성 성공", func(t *testing.T) {
		tr.On("Get", mock.AnythingOfType("int64")).Return(expectedResponse, nil)
		tr.On("Complete", mock.AnythingOfType("int64")).Return(expectedResponse, nil)
		ts := NewTodoService(tr)

		todoID := int64(1)
		email := "hwc9169@gmail.com"
		resp, err := ts.CompleteTodo(todoID, email)
		assert.NoError(t, err)

		expected := dto.TodoToDTO(expectedResponse)
		assert.Equal(t, expected, resp)
	})
	t.Run("권한 오류 발생", func(t *testing.T) {
		tr.On("Get", mock.AnythingOfType("int64")).Return(expectedResponse, nil)
		tr.On("Complete", mock.AnythingOfType("int64")).Return(expectedResponse, nil)
		ts := NewTodoService(tr)

		todoID := int64(1)
		email := "hwc9169@naver.com"
		_, err := ts.CompleteTodo(todoID, email)
		assert.Equal(t, err, echo.NewHTTPError(http.StatusForbidden, "해당 요청에 대한 권한이 없습니다."))
	})
}

func TestDeleteTodo(t *testing.T) {
	tr := new(mocks.TodoRepository)
	user := &ent.User{
		ID:       "hwc9169@gmail.com",
		Password: "password",
		Name:     "조호원",
	}
	deadline := time.Now().Add(24 * 3 * time.Hour)
	expectedResponse := &ent.Todo{
		ID:          1,
		Title:       "Go 언어 공부하기",
		Content:     "장재휴의 Go 웹 프로그래밍 철저 입문",
		Deadline:    &deadline,
		IsCompleted: false,
		Edges: ent.TodoEdges{
			User: user,
		},
	}
	t.Run("Todo 삭제 성공", func(t *testing.T) {
		tr.On("Get", mock.AnythingOfType("int64")).Return(expectedResponse, nil)
		tr.On("Delete", mock.AnythingOfType("int64")).Return(expectedResponse, nil)
		ts := NewTodoService(tr)

		todoID := int64(1)
		email := "hwc9169@gmail.com"
		resp, err := ts.DeleteTodo(todoID, email)
		assert.NoError(t, err)

		expected := dto.TodoToDTO(expectedResponse)
		assert.Equal(t, expected, resp)
	})
	t.Run("권한 오류 발생", func(t *testing.T) {
		tr.On("Get", mock.AnythingOfType("int64")).Return(expectedResponse, nil)
		tr.On("Delete", mock.AnythingOfType("int64")).Return(expectedResponse, nil)
		ts := NewTodoService(tr)

		todoID := int64(1)
		email := "hwc9169@naver.com"
		_, err := ts.DeleteTodo(todoID, email)
		assert.Equal(t, err, echo.NewHTTPError(http.StatusForbidden, "해당 요청에 대한 권한이 없습니다."))

	})
}
