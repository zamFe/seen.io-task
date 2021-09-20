# seen.io-task

## Communications API

A REST-API written in GoLang, using a psql database.
Object-relation mappings done using the gorm-package.

### Requirements
Applications requires a config.go with database parameters:

![config template](https://github.com/zamFe/seen.io-task/blob/main/SeenIO/images/config_template.png?raw=true)
### Setup

1. create database and set up config.go accordingly
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

