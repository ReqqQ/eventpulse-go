//go:generate wire
package di

import (
	"github.com/ReqqQ/eventpulse-user-go/src/app/user/command"
	"github.com/ReqqQ/eventpulse-user-go/src/app/user/handler"
	"github.com/ReqqQ/eventpulse-user-go/src/app/user/query"
	userAppRepository "github.com/ReqqQ/eventpulse-user-go/src/app/user/repository"
	userDomainRepository "github.com/ReqqQ/eventpulse-user-go/src/domain/facebook/repository"
	userDomainFactory "github.com/ReqqQ/eventpulse-user-go/src/domain/user/factory"
	"github.com/ReqqQ/eventpulse-user-go/src/domain/user/service"
	userFacebookRepository "github.com/ReqqQ/eventpulse-user-go/src/infrastructure/facebook/repository"
	userFactory "github.com/ReqqQ/eventpulse-user-go/src/infrastructure/users/factory"
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
	userDomainFactory.BuildFactory,
	userFacebookRepository.BuildApiFacebookRepository,
	service.BuildService,
	userFactory.BuildUserFactory,
	InitializeAppFacebookRepository,
	InitializeDomainFacebookRepository,
)

func InitializeAppFacebookRepository() userAppRepository.FacebookRepository {
	wire.Build(userFacebookRepository.BuildApiFacebookRepository, userFacebookRepository.BuildAppFacebookRepository)
	return nil
}

func InitializeDomainFacebookRepository() userDomainRepository.FacebookRepository {
	wire.Build(userFacebookRepository.BuildApiFacebookRepository, userFacebookRepository.BuildDomainFacebookRepository)
	return nil
}
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
