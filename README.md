# UW Timeline

## What's Course Planner?
Course Planner was the initial project name before we chose
a better name.

## Purpose
In summary, we wanted to create an application that helps
University of Waterloo students to plan future courses with
ease. Some of the features include:
- Checking plan requirements for degrees, minors, options, etc.
- Checking course pre-requisites using previous courses
- The flexibility to create different Timelines i.e. different
academic plans

## Running the application
- make sure you have `golang 1.11` or above installed.
- make sure you have a local mongo instance running, we use
the docker image `mongo:4.0.10-xenial`.
- pull from master and run `dep ensure` to install any
dependencies.

### Running the API server
- make sure you're in the project root directory
- run `go run ./cmd/course-planner`

### Running tests
- make sure you're in the project root directory
- run `go test -v -tags=unit ./...` for unit tests
- run `go test -v -p 1 -tags=integration ./...` for integration tests
- run `go test -v -p 1 -tags=all ./...` for both unit and integration
tests.

## Contributing
Please read the [contributing](contributing.md) document.
