//go:generate wire
package di

import (
	"github.com/ReqqQ/eventpulse-go/src/ui"
	"github.com/ReqqQ/eventpulse-user-go/app/user"
	"github.com/google/wire"
)

func InitDIContainer() ui.Controllers {
	wire.Build(
		user.BuildUserService,
		ui.BuildUserController,
		ui.BuildControllers,
	)
	return nil
}
