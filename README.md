# Fita Appointment API
Coach appointment API implemented in Go with Gin framwork, go-fx dependency container, and MongoDB database.

## Getting Started
1. Copy `.env.example` to `.env`
2. Setup MongoDB server. Can use docker container by executing `$ docker run --name mongodb -d -p 27017:27017 mongo`
3. Migrate pre-processed coach availability dataset(`data.csv`). `$ mongoimport -d fita -c CoachAvailability --type csv --file data.csv --headerline`
4. Build and install dependency. `$ go build`
5. Run the HTTP server. `$ go run . app:serve`

## Implemented Features
- Make an appointment & book the schedule
- Validate an appointment based on coach availability
- Coach can decline an appointment request or reschedule it
- If user decline the rescheduling, then all ended (no need to provide another rescheduling)

## Assumptions
- An appointment made by the users will not exceed 24 hours time interval.
- For the sake of simplicity, coach and user names are unique. Hence it can be used as coach and user identifier.
- Appointment status approval should be made by coach and ideally the request is authenticated, both as a coach and as an individual. For the sake of simplicity, no authentication is implemented here, all requests are assumed to be requested by the right party. The same case for the rescheduled appointment status approval.
- Reschedule assumed to ignore collision with another appointment. The coach might allocate special time for the rescheduled appointment, or handle an appointment with other rescheduled appointments simultaneously.

## API Contract
Accessible through [this Postman Collection](https://www.getpostman.com/collections/7eb03ebd40f027e8ade4)

## What can be improved
- Provides index to the fields that often used as the filter.
- Handle more edge cases such as appointments with more than 24 hours duration.
- Automate coach availability data pre-processing.
