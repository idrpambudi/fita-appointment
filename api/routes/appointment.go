package routes

import (
	"github.com/idrpambudi/fita-appointment/api/controllers"
	"github.com/idrpambudi/fita-appointment/lib"
)

type AppointmentRoutes struct {
	handler    lib.RequestHandler
	controller controllers.AppointmentController
}

// Setup user routes
func (r AppointmentRoutes) Setup() {
	api := r.handler.Gin.Group("/appointment")
	{
		api.POST("", r.controller.CreateAppointment)
		api.PUT("/:id", r.controller.DecideAppointment)
		api.POST("/:id/rescheduled", r.controller.RescheduleAppointment)
		api.PUT("/:id/rescheduled", r.controller.DecideRescheduledAppointment)
	}
}

func NewAppointmentRoutes(
	handler lib.RequestHandler,
	controller controllers.AppointmentController,
) AppointmentRoutes {
	return AppointmentRoutes{
		handler:    handler,
		controller: controller,
	}
}
