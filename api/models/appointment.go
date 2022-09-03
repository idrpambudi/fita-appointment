package models

import (
	"time"

	"github.com/idrpambudi/fita-appointment/api/dto"
	"github.com/idrpambudi/fita-appointment/lib"
	"go.mongodb.org/mongo-driver/mongo"
)

const appointmentCollectionName = "Appointment"

type AppointmentCollection struct {
	*mongo.Collection
}

type Appointment struct {
	ID               string    `json:"ID,omitempty" bson:"_id,omitempty"`
	Start            time.Time `binding:"required"`
	End              time.Time `binding:"required"`
	Coach            string    `binding:"required"`
	User             string    `binding:"required"`
	Status           dto.AppointmentStatus
	RescheduledStart *time.Time
	RescheduledEnd   *time.Time
}

func NewAppointmentCollection(db lib.Database) AppointmentCollection {
	return AppointmentCollection{db.Collection(appointmentCollectionName)}
}
