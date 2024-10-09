package factory

import (
	"github.com/ReqqQ/eventpulse-user-go/src/app/user/handler/interfaces"

	"github.com/ReqqQ/eventpulse-go/src/app/core/factory"
)

const (
	getLoginDialogQueryType        = "GetLoginDialogQuery"
	getUserQueryType               = "GetUserQuery"
	getCreateUserFromSocialCommand = "GetCreateUserFromSocialCommand"
)

type getUserQueryImpl struct {
	userId string
}

func (g getUserQueryImpl) GetQueryType() string {
	return getUserQueryType
}

func (g getUserQueryImpl) GetQueryUserId() string {
	return g.userId
}

type getLoginDialogQueryImpl struct {
	socialType string
}

func (g getLoginDialogQueryImpl) GetQueryType() string {
	return getLoginDialogQueryType
}

func (g getLoginDialogQueryImpl) GetSocialType() string {
	return g.socialType
}

type userFactory struct{}

type createUserFromSocialCommandImpl struct {
	code string
}

func (c *createUserFromSocialCommandImpl) GetCommandType() string {
	return getCreateUserFromSocialCommand
}

func (c *createUserFromSocialCommandImpl) GetCode() string {
	return c.code
}

func (u *userFactory) GetCreateUserFromSocialCommand(code string) interfaces.CreateUserFromSocialCommand {
	return &createUserFromSocialCommandImpl{code: code}
}

func (u *userFactory) GetLoginDialogQuery(socialType string) interfaces.GetLoginDialogQuery {
	return &getLoginDialogQueryImpl{socialType: socialType}
}

func (u *userFactory) GetUserQuery(userId string) interfaces.GetUserQuery {
	return &getUserQueryImpl{userId: userId}
}

func BuildFactory() factory.UserFactory {
	return &userFactory{}
}
