package ui

import (
	"github.com/ReqqQ/eventpulse-go/src/app"
)

type FBTokenUser struct {
	AccessToken string `json:"access_token"`
}
type FBUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type serverImpl struct {
	apps app.App
}

func (s *serverImpl) GetApps() app.App {
	return s.apps
}

func BuildServer(apps app.App) app.Server {
	return &serverImpl{apps: apps}
}
