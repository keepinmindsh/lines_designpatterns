package usecase

import "authorized_user/domain"

type UserUCase struct {
}

func NewUser() *UserUCase {
	return &UserUCase{}
}

func (u UserUCase) GetUser() []domain.UserModel {
	//TODO implement me
	panic("implement me")
}
