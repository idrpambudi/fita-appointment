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

## Flow Details
### Create an Appointment
- Book an appointment and wait for it to be accepted by the coach, marked by `Scheduled` status
- Book appointment request shall be validated based on coach availability and other appointments collision potential
- Accept local timezone in the request body, but the data will be stored and validated in UTC format
### Approve/Decline an Appointment
- Change the status of an appointment into `Accepted`/`Declined`, should be requested by a coach
- Only applicable for appointment with status `Scheduled`, or else the response will be `4xx status code`
### Reschedule an Appointment
- Reschedule an appointment with status `Scheduled` only
- Accept any time interval for reschedule, ignoring coach availability and collision with other appointments
### Approve/Decline a Rescheduled Appointment
- Change the status of a rescheduled appointment into `Accepted`/`Declined`, should be requested by a user
- Appointment status and rescheduled appointment status share the same field in the `Appointment` collection
- Only applicable for appointment with status `Recheduled`, or else the response will be `4xx status code`

## API Contract
Accessible through [this Postman Collection](https://www.getpostman.com/collections/7eb03ebd40f027e8ade4)

## What can be improved
- Provides index to the fields that often used as the filter.
- Handle more edge cases such as appointments with more than 24 hours duration.
- Automate coach availability data pre-processing.
- Setup coach and user collection to accomodate same name but different individual data.
- Apply authentication and authorization middleware to authenticate an action toward an appointment.