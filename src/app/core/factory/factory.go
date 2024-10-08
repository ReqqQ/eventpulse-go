package factory

import (
	"github.com/ReqqQ/eventpulse-user-go/src/app/user/query"
)

type UserFactory interface {
	GetLoginDialogQuery(socialType string) query.GetLoginDialogQuery
	GetUserQuery(userId string) query.GetUserQuery
}
