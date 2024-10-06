//go:generate wire
package di

import (
	appUser "github.com/ReqqQ/eventpulse-user-go/src/app/user"
	domainUser "github.com/ReqqQ/eventpulse-user-go/src/domain/user"
	userInfrastructureFacebook "github.com/ReqqQ/eventpulse-user-go/src/infrastructure/facebook"
	userInfrastructureUser "github.com/ReqqQ/eventpulse-user-go/src/infrastructure/users"
	"github.com/ReqqQ/eventpulse-user-go/src/shared"
	"github.com/google/wire"

	"github.com/ReqqQ/eventpulse-go/src/app"
	appInfrastructure "github.com/ReqqQ/eventpulse-go/src/infrastructure/user"
	"github.com/ReqqQ/eventpulse-go/src/ui"
)

var userInterfaceFacebookRepository = wire.NewSet(
	userInfrastructureFacebook.BuildApiFacebookRepository,
	wire.Bind(new(appUser.FacebookRepository), new(userInfrastructureFacebook.ApiFacebookRepository)),
)

var buildUserApp = wire.NewSet(
	shared.CreateUserBusInstance,
	appUser.CreateQueryHandler,
	userInfrastructureUser.CreateRepository,
	domainUser.BuildFactory,
	domainUser.BuildService,
	userInterfaceFacebookRepository,
)

func InitDIContainer() app.Server {
	wire.Build(
		wire.NewSet(
			ui.BuildServer,
			app.BuildApp,
			app.BuildUserApp,
			appInfrastructure.BuildUserFactory,
			buildUserApp,
		),
	)

	return nil
}
