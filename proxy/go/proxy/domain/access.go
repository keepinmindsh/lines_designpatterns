package domain

type AuthorizedAccess interface {
	Access()
}

type UnAuthorizedAccess interface {
	NotValidAccess()
}

type Acceptor interface {
	Accept(access AuthorizedAccess)
}
