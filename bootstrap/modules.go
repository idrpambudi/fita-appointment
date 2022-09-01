package bootstrap

import (
	"github.com/idrpambudi/fita-appointment/api/controllers"
	"github.com/idrpambudi/fita-appointment/api/middlewares"
	"github.com/idrpambudi/fita-appointment/api/routes"
	"github.com/idrpambudi/fita-appointment/lib"
	"github.com/idrpambudi/fita-appointment/repository"
	"github.com/idrpambudi/fita-appointment/services"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	repository.Module,
)
