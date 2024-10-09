package factory

import (
	"github.com/ReqqQ/eventpulse-user-go/src/app/user/handler/interfaces"

	"github.com/ReqqQ/eventpulse-go/src/app/core/factory"
)

const (
	getLoginDialogQueryType = "GetLoginDialogQuery"
	getUserQueryType        = "GetUserQuery"
)

type getUserQuery struct {
	userId string
}

func (g getUserQuery) GetQueryType() string {
	return getUserQueryType
}

func (g getUserQuery) GetQueryUserId() string {
	return g.userId
}

type getLoginDialogQuery struct {
	socialType string
}

func (g getLoginDialogQuery) GetQueryType() string {
	return getLoginDialogQueryType
}

func (g getLoginDialogQuery) GetSocialType() string {
	return g.socialType
}

type userFactory struct{}

func (u *userFactory) GetLoginDialogQuery(socialType string) interfaces.GetLoginDialogQuery {
	return getLoginDialogQuery{socialType: socialType}
}

func (u *userFactory) GetUserQuery(userId string) interfaces.GetUserQuery {
	return getUserQuery{userId: userId}
}

func BuildFactory() factory.UserFactory {
	return &userFactory{}
}
