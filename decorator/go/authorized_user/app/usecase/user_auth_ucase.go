package usecase

import (
	"authorized_user/domain"
	"github.com/samber/lo"
)

type UserAuthUCase struct {
	user domain.User
}

func NewUserAuthUCase(user domain.User) domain.AuthorizedUser {
	return &UserAuthUCase{
		user: user,
	}
}

func (u UserAuthUCase) GetAuthorizedUser() []domain.UserModel {

	userList := u.user.GetUser()

	filter := lo.Filter(userList, func(item domain.UserModel, index int) bool {
		return checkRole(item.Role)
	})

	return filter
}

func checkRole(role string) bool {
	return role == "admin"
}
