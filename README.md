# seen.io-task

## Communications API

A REST-API written in GoLang, using a psql database.
Object-relation mappings done using the gorm-package.

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
