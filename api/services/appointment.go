package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/idrpambudi/fita-appointment/api/dto"
	"github.com/idrpambudi/fita-appointment/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	AppointmentService struct {
		CoachAvailabilityCollection models.CoachAvailabilityCollection
		AppointmentCollection       models.AppointmentCollection
	}
)

func NewAppointmentService(cac models.CoachAvailabilityCollection, ac models.AppointmentCollection) AppointmentService {
	return AppointmentService{
		CoachAvailabilityCollection: cac,
		AppointmentCollection:       ac,
	}
}

func (as AppointmentService) CreateAppointment(appointment models.Appointment) (*models.Appointment, error) {
	if err := as.validateAppointment(appointment); err != nil {
		return nil, err
	}
	appointment.ID = primitive.NewObjectID().Hex()
	appointment.Status = dto.Scheduled
	_, err := as.AppointmentCollection.InsertOne(context.Background(), appointment)
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (as AppointmentService) DecideAppointment(id string, status dto.AppointmentStatus) error {
	var find models.Appointment
	if err := as.AppointmentCollection.FindOne(context.Background(), bson.M{
		"_id":    id,
		"status": dto.Scheduled,
	}).Decode(&find); err != nil {
		return err
	}

	if _, err := as.AppointmentCollection.UpdateByID(context.Background(), id, bson.M{"$set": bson.M{"status": status}}); err != nil {
		return err
	}
	return nil
}

func (as AppointmentService) RescheduleAppointment(id string, start time.Time, end time.Time) error {
	var find models.Appointment
	if err := as.AppointmentCollection.FindOne(context.Background(), bson.M{
		"_id":    id,
		"status": dto.Scheduled,
	}).Decode(&find); err != nil {
		return err
	}

	if _, err := as.AppointmentCollection.UpdateByID(context.Background(), id, bson.M{"$set": bson.M{
		"status":           dto.Rescheduled,
		"rescheduledstart": start,
		"rescheduledend":   end,
	}}); err != nil {
		return err
	}
	return nil
}

func (as AppointmentService) DecideRescheduledAppointment(id string, status dto.AppointmentStatus) error {
	var find models.Appointment
	if err := as.AppointmentCollection.FindOne(context.Background(), bson.M{
		"_id":    id,
		"status": dto.Rescheduled,
	}).Decode(&find); err != nil {
		return err
	}

	_, err := as.AppointmentCollection.UpdateByID(context.Background(), id, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		return err
	}
	return nil
}

func (as AppointmentService) validateAppointment(appointment models.Appointment) error {
	if !as.isCoachAvailable(appointment) {
		return errors.New("Coach is unavailable at the chosen appointment time")
	}
	if as.isAppointmentCollide(appointment) {
		return errors.New("The chosen appointment time collide with another appointment")
	}
	return nil
}

func (as *AppointmentService) isAppointmentCollide(appointment models.Appointment) bool {
	return as.AppointmentCollection.FindOne(context.Background(), bson.M{
		"coach":  appointment.Coach,
		"status": bson.M{"$ne": "Declined"},
		"$or": bson.A{
			bson.M{
				"rescheduledstart": bson.M{"$in": bson.A{"", nil}},
				"start":            bson.M{"$lt": appointment.End},
				"end":              bson.M{"$gt": appointment.Start},
			}, bson.M{
				"rescheduledstart": bson.M{"$lt": appointment.End},
				"rescheduledend":   bson.M{"$gt": appointment.Start},
			},
		},
	}).Err() == nil
}

func (as *AppointmentService) isCoachAvailable(appointment models.Appointment) bool {
	startDay, startHourMinute := getWeeklyTimeInterval(appointment.Start)
	endDay, endHourMinute := getWeeklyTimeInterval(appointment.End)
	if startDay == endDay {
		return as.findCoachAvailability(appointment.Coach, startDay, startHourMinute, endHourMinute).Err() == nil
	}
	return as.findCoachAvailability(appointment.Coach, startDay, startHourMinute, "24:00").Err() == nil && as.findCoachAvailability(appointment.Coach, endDay, "00:00", endHourMinute).Err() == nil
}

func (as *AppointmentService) findCoachAvailability(coach, day, startHourMinute, endHourMinute string) *mongo.SingleResult {
	return as.CoachAvailabilityCollection.FindOne(context.Background(), bson.M{
		"name":      coach,
		"dayofweek": day,
		"start":     bson.M{"$lte": startHourMinute},
		"end":       bson.M{"$gte": endHourMinute},
	})
}

func getWeeklyTimeInterval(timestamp time.Time) (dayOfWeek, hourMinute string) {
	dayOfWeek = timestamp.UTC().Weekday().String()
	hour := timestamp.UTC().Hour()
	minute := timestamp.UTC().Minute()
	hourMinute = fmt.Sprintf("%02d:%02d", hour, minute)
	return
}
