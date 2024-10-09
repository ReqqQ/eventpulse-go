//go:generate wire
package di

import (
	"github.com/ReqqQ/eventpulse-user-go/src/app/user/command"
	"github.com/ReqqQ/eventpulse-user-go/src/app/user/handler"
	"github.com/ReqqQ/eventpulse-user-go/src/app/user/query"
	domainUserFactory "github.com/ReqqQ/eventpulse-user-go/src/domain/user/factory"
	"github.com/ReqqQ/eventpulse-user-go/src/domain/user/service"
	"github.com/ReqqQ/eventpulse-user-go/src/infrastructure/facebook"
	infrastructureUserFactory "github.com/ReqqQ/eventpulse-user-go/src/infrastructure/users/factory"
	"github.com/ReqqQ/eventpulse-user-go/src/infrastructure/users/repository"
	"github.com/ReqqQ/eventpulse-user-go/src/shared"
	"github.com/google/wire"

	"github.com/ReqqQ/eventpulse-go/src/infrastructure/user/factory"
	"github.com/ReqqQ/eventpulse-go/src/ui/routes"
)

var buildUserApp = wire.NewSet(
	query.BuildUserQueryHandler,
	command.BuildCommandQueryHandler,
	repository.BuildUserRepository,
	domainUserFactory.BuildFactory,
	facebook.BuildApiFacebookRepository,
	service.BuildService,
	infrastructureUserFactory.BuildUserFactory,
)

func InitDIContainer() routes.Routes {
	wire.Build(
		wire.NewSet(
			routes.BuildRoutes,
			factory.BuildFactory,
			shared.BuildBus,
			handler.BuildUserBus,
			buildUserApp,
		),
	)

	return nil
}
