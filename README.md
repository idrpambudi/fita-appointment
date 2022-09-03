# Fita Appointment API

Coach appointment API implemented in Go with Gin framwork, go-fx dependency container, and MongoDB database.

## Getting Started
1. Copy `.env.example` to `.env`
2. 

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

## Running the project

- Make sure you have docker installed.
- Copy `.env.example` to `.env`
- Run `docker-compose up -d`
- Go to `localhost:5000` to verify if the server works.
- [Adminer](https://www.adminer.org/) Database Management runs at `5001` .

If you are running without docker be sure database configuration is provided in `.env` file and run `go run . app:serve`

#### Environment Variables

<details>
    <summary>Variables Defined in the project </summary>

| Key            | Value                    | Desc                                        |
| -------------- | ------------------------ | ------------------------------------------- |
| `SERVER_PORT`  | `5000`                   | Port at which app runs                      |
| `ENV`          | `development,production` | App running Environment                     |
| `LOG_OUTPUT`   | `./server.log`           | Output Directory to save logs               |
| `LOG_LEVEL`    | `info`                   | Level for logging (check lib/logger.go:172) |
| `DB_USER`      | `username`               | Database Username                           |
| `DB_PASS`      | `password`               | Database Password                           |
| `DB_HOST`      | `0.0.0.0`                | Database Host                               |
| `DB_PORT`      | `3306`                   | Database Port                               |
| `DB_NAME`      | `test`                   | Database Name                               |
| `JWT_SECRET`   | `secret`                 | JWT Token Secret key                        |
| `ADMINER_PORT` | `5001`                   | Adminer DB Port                             |
| `DEBUG_PORT`   | `5002`                   | Port that delve debugger runs in            |

</details>

#### Migration Commands

> ⚓️ &nbsp; Add argument `p=host` if you want to run the migration runner from the host environment instead of docker environment.
> Check [#19](https://github.com/dipeshdulal/clean-gin/issues/19) for more details. eg; `make p=host migrate-up`

<details>
    <summary>Migration commands available</summary>

| Command             | Desc                                           |
| ------------------- | ---------------------------------------------- |
| `make migrate-up`   | runs migration up command                      |
| `make migrate-down` | runs migration down command                    |
| `make force`        | Set particular version but don't run migration |
| `make goto`         | Migrate to particular version                  |
| `make drop`         | Drop everything inside database                |
| `make create`       | Create new migration file(up & down)           |

</details>

## Implemented Features

- Dependency Injection (go-fx)
- Routing (gin web framework)
- Environment Files
- Logging (file saving on `production`) [zap](https://github.com/uber-go/zap)
- Middlewares (cors)
- Database Setup (mysql)
- Models Setup and Automigrate (gorm)
- Repositories
- Implementing Basic CRUD Operation
- Authentication (JWT)
- Migration Runner Implementation
- Live code refresh
- Dockerize Application with Debugging Support Enabled. Debugger runs at `5002`. Vs code configuration is at `.vscode/launch.json` which will attach debugger to remote application. [Learn More](https://medium.com/wesionary-team/docker-debug-environment-for-go-and-gin-framework-36df80e061ac?source=friends_link&sk=35c9d856852944083dd30059200d87f0)
- Cobra Commander CLI Support. try: `go run . --help`

## Todos

- [x] COBRA Commander CLI Support [#26](https://github.com/dipeshdulal/clean-gin/issues/26)
- [ ] Swagger documentation examples [#25](https://github.com/dipeshdulal/clean-gin/issues/25)
- [ ] Unit testing examples. [#23](https://github.com/dipeshdulal/clean-gin/issues/23)
- [ ] File upload middelware. [#20](https://github.com/dipeshdulal/clean-gin/issues/20)
- [ ] Use of Interfaces [#10](https://github.com/dipeshdulal/clean-gin/issues/10)

## Contributing

Please open issues if you want the template to add some features that is not in todos. 🙇‍♂️

Create a PR with relevant information if you want to contribute in this template.
