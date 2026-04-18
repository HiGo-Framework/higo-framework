package routersfx

import (
	"github.com/triasbrata/higo-framework/routers/impl"
	"go.uber.org/fx"
)

func LoadModuleRouter() fx.Option {
	return fx.Module("pkgs/routers/fx", fx.Provide(impl.NewEngine))
}
