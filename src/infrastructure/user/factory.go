package ui

import (
	"github.com/ReqqQ/eventpulse-user-go/src/app/user"

	"github.com/ReqqQ/eventpulse-go/src/app"
)

type GetUserQuery struct {
	userId string
}

type GetLoginDialogQuery struct {
	socialType string
}

func (g GetLoginDialogQuery) GetQueryType() string {
	return "getLoginDialogQuery"
}

func (g GetLoginDialogQuery) GetSocialType() string {
	return g.socialType
}

func (g GetUserQuery) GetQueryUserId() string {
	return g.userId
}
func (g GetUserQuery) GetQueryType() string {
	return "GetUserQuery"
}

type userFactoryImpl struct {
}

func (u *userFactoryImpl) GetUserQuery(userId string) user.GetUserQuery {
	return &GetUserQuery{userId: userId}
}

func (u *userFactoryImpl) GetLoginDialogQuery(socialType string) user.GetLoginDialogQuery {
	return &GetLoginDialogQuery{socialType: socialType}
}

func BuildUserFactory() app.UserFactory {
	return &userFactoryImpl{}
}
