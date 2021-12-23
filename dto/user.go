package dto

import "halill/ent"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RegistRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func UserToDTO(src *ent.User) *UserResponse {
	return &UserResponse{
		Email: src.ID,
		Name:  src.Name,
	}
}
