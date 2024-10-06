package ui

import (
	"github.com/ReqqQ/eventpulse-user-go/src/app/user"

	"github.com/ReqqQ/eventpulse-go/src/app"
)

type GetUserQuery struct {
	userId string
}

func (g GetUserQuery) GetQueryUserId() string {
	return g.userId
}
func (g GetUserQuery) GetType() string {
	return "GetUserQuery"
}

type userFactoryImpl struct {
}

func (u *userFactoryImpl) GetUserQuery(userId string) user.GetUserQuery {
	return &GetUserQuery{userId: userId}
}

func BuildUserFactory() app.UserFactory {
	return &userFactoryImpl{}
}
