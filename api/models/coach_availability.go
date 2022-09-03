package models

import (
	"github.com/idrpambudi/fita-appointment/lib"
	"go.mongodb.org/mongo-driver/mongo"
)

const coachAvailabilityCollectionName = "CoachAvailability"

type CoachAvailabilityCollection struct {
	*mongo.Collection
}

type CoachAvailability struct {
	ID    string `json:"id,omitempty" bson:"_id"`
	Start string
	End   string
	Name  string
}

func NewCoachAvailabilityCollection(db lib.Database) CoachAvailabilityCollection {
	return CoachAvailabilityCollection{db.Collection(coachAvailabilityCollectionName)}
}
