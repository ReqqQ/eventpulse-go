//go:generate wire
package di

import (
	appUser "github.com/ReqqQ/eventpulse-user-go/src/app/user"
	domainUser "github.com/ReqqQ/eventpulse-user-go/src/domain/user"
	infrastructureUser "github.com/ReqqQ/eventpulse-user-go/src/infrastructure/users"
	"github.com/ReqqQ/eventpulse-user-go/src/shared"
	"github.com/google/wire"

	"github.com/ReqqQ/eventpulse-go/src/ui"
)

var buildUserApp = wire.NewSet(
	shared.CreateUserBusInstance,
	appUser.CreateQueryHandler,
	infrastructureUser.CreateRepository,
	domainUser.BuildFactory,
)

func InitDIContainer() ui.Server {
	wire.Build(
		wire.NewSet(
			ui.BuildServer,
			ui.BuildApp,
			ui.BuildUserApp,
			buildUserApp,
		),
	)

	return nil
}
