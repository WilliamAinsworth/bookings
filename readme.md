# Bookings and Reservations

This is the repository for my bookings project.

- Built in Go version 1.16
- Uses the [alex edwards SCS session management](http://github.com/alexedwards/scs)
- Uses the [chi router](http://github.com/go-chi/chi)
- Uses [nosurf](http://github.com/justinas/nosurf)

## Setup

### Testing

go test -coverprofile=coverage.out && go tool cover -html=coverage.out