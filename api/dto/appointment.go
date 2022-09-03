package dto

import "time"

type AppointmentStatus string

const (
	Scheduled   AppointmentStatus = "Scheduled"
	Declined    AppointmentStatus = "Declined"
	Accepted    AppointmentStatus = "Accepted"
	Rescheduled AppointmentStatus = "Rescheduled"
)

type (
	RescheduleBody struct {
		Start time.Time `binding:"required"`
		End   time.Time `binding:"required"`
	}

	DecideAppointmentBody struct {
		Decision AppointmentStatus `binding:"required"`
	}
)
