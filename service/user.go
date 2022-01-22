package service

import (
	"halill/dto"
	"halill/ent"
	"halill/repository"
	"halill/security"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	LoginUser(*dto.LoginRequest) (*dto.TokenResponse, error)
	RegistUser(*dto.RegistRequest) (*dto.UserResponse, error)
	RefreshToken(*dto.RefreshTokenRequest) (*dto.TokenResponse, error)
}

type userServiceImpl struct {
	ur repository.UserRepository
	jp security.JWTProvider
}

func NewUserSerice(ur repository.UserRepository, jp security.JWTProvider) UserService {
	return &userServiceImpl{
		ur: ur,
		jp: jp,
	}
}

func (s *userServiceImpl) LoginUser(r *dto.LoginRequest) (*dto.TokenResponse, error) {
	user, err := s.verifyUser(r)
	if err != nil {
		return nil, err
	}

	at, err := s.jp.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	rt, err := s.jp.GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.TokenResponse{
		AccessToken:  at,
		RefreshToken: rt,
	}, nil
}

func (s *userServiceImpl) verifyUser(r *dto.LoginRequest) (*ent.User, error) {
	user, err := s.ur.GetByEmail(r.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userServiceImpl) RegistUser(r *dto.RegistRequest) (*dto.UserResponse, error) {
	_, err := s.ur.GetByEmail(r.Email)

	// 이미 사용중인 이메일이면 실패
	if err == nil {
		err = echo.NewHTTPError(http.StatusBadRequest, "이미 사용중인 이메일입니다.")
		return nil, err
	}
	// Notfound 에러가 아니면 실패
	if _, ok := err.(*echo.HTTPError); !ok {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &ent.User{
		ID:       r.Email,
		Password: string(hashedPassword),
		Name:     r.Name,
	}
	newUser, err := s.ur.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return dto.UserToDTO(newUser), nil
}

func (s *userServiceImpl) RefreshToken(r *dto.RefreshTokenRequest) (*dto.TokenResponse, error) {
	claims := &security.JwtCustomClaims{}
	_, err := jwt.ParseWithClaims(r.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jp.JwtSecret()), nil
	})
	if err != nil {
		return nil, err
	}

	user, err := s.ur.GetByEmail(claims.Email)
	if err != nil {
		return nil, err
	}

	at, err := s.jp.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.TokenResponse{
		AccessToken:  at,
		RefreshToken: r.RefreshToken,
	}, nil
}
