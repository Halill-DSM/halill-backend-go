package repository

type UserRepository interface {
}

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}
