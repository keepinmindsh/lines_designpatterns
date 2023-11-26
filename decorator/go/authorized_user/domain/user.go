package domain

type (
	UserModel struct {
		UserId   string
		UserName string
		Tel      string
		Address  string
		Role     string ``
	}

	User interface {
		GetUser() []UserModel
	}

	AuthorizedUser interface {
		GetAuthorizedUser() []UserModel
	}
)
