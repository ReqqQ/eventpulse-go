package factory

import (
	"github.com/ReqqQ/eventpulse-user-go/src/app/user/handler/interfaces"
)

type UserFactory interface {
	GetLoginDialogQuery(socialType string) interfaces.GetLoginDialogQuery
	GetUserQuery(userId string) interfaces.GetUserQuery
	GetCreateUserFromSocialCommand(code string) interfaces.CreateUserFromSocialCommand
}
