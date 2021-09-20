# seen.io-task

## Communications API

A REST-API written in GoLang, using a psql database.
Object-relation mappings done using the gorm-package.

### Requirements
Applications requires a config.go file with two consts:

- DbUser
- DbPassword

### Setup

1. set up config file as mentioned above
2. run `go install seenio/dbContext` in the dbContext folder
3. `dbContext` in terminal to start application

### Tables

#### User

[string] Email

[string] PhoneNumber

#### EventLog

[int] LandingPageHits

[int] VideoPlays

[int] UserID

### Endpoints

#### GET

`/eventLogs` 

`/eventlogs/{id}`

`/users`

`/users/{id}`

#### PATCH

`/eventlogs/videoplays/{id}`

`/eventlogs/landingpagehits/{id}`

#### DELETE

`/eventlogs/{id}`

