package bootstrap

import (
	"github.com/idrpambudi/fita-appointment/api/controllers"
	"github.com/idrpambudi/fita-appointment/api/middlewares"
	"github.com/idrpambudi/fita-appointment/api/models"
	"github.com/idrpambudi/fita-appointment/api/routes"
	"github.com/idrpambudi/fita-appointment/api/services"
	"github.com/idrpambudi/fita-appointment/lib"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	models.Module,
	middlewares.Module,
)
