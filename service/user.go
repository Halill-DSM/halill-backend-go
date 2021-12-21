package service

type UserService interface {
}

type userServiceImpl struct {
}

func NewUserSerice() UserService {
	return &userServiceImpl{}
}
