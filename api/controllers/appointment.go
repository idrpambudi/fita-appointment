package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idrpambudi/fita-appointment/api/dto"
	"github.com/idrpambudi/fita-appointment/api/models"
	"github.com/idrpambudi/fita-appointment/api/services"
	"github.com/idrpambudi/fita-appointment/lib"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppointmentController struct {
	service services.AppointmentService
	logger  lib.Logger
}

func NewAppointmentController(appointmentService services.AppointmentService, logger lib.Logger) AppointmentController {
	return AppointmentController{
		service: appointmentService,
		logger:  logger,
	}
}

// missing validate appointment time response type
func (ac AppointmentController) CreateAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if appointment.Start.After(appointment.End) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment time interval"})
		return
	}
	res, err := ac.service.CreateAppointment(appointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": res})
}

func (ac AppointmentController) DecideAppointment(c *gin.Context) {
	var decision dto.DecideAppointmentBody
	if err := c.ShouldBindJSON(&decision); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	switch decision.Decision {
	case dto.Accepted, dto.Declined:
		break
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment approval decision"})
		return
	}

	id := c.Param("id")
	err := ac.service.DecideAppointment(id, decision.Decision)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// missing validate reschedule time response type
func (ac AppointmentController) RescheduleAppointment(c *gin.Context) {
	var reschedule dto.RescheduleBody
	if err := c.ShouldBindJSON(&reschedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if reschedule.Start.After(reschedule.End) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reschedule time interval"})
		return
	}

	id := c.Param("id")
	err := ac.service.RescheduleAppointment(id, reschedule.Start, reschedule.End)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (ac AppointmentController) DecideRescheduledAppointment(c *gin.Context) {
	var decision dto.DecideAppointmentBody
	if err := c.ShouldBindJSON(&decision); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	switch decision.Decision {
	case dto.Accepted, dto.Declined:
		break
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment approval decision"})
		return
	}

	id := c.Param("id")
	err := ac.service.DecideRescheduledAppointment(id, decision.Decision)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
