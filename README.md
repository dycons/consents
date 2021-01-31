# Consents Service
A microservice for consents metadata

Strongly modeled on the [Go model service](https://github.com/CanDIG/go-model-service) built for CanDIG in Summer 2018 and upgraded in Summer 2020. The stacks are nearly identical and the implementations are very similar.
If you are interested in contributing to this toy service, or in building something similar, you may benefit from taking a look at the [Go model service](https://github.com/CanDIG/go-model-service) and its documentation.

[![Build Status](https://travis-ci.org/dycons/consents.svg?branch=develop)](https://travis-ci.org/dycons/consents)

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->
<!-- code_chunk_output -->

- [Consents Service](#consents-service)
  - [Quick Start](#quick-start)
  - [Stack](#stack)
    - [Installing the stack](#installing-the-stack)
  - [Installing the service](#installing-the-service)
  - [Running The Service](#running-the-service)
    - [Request examples](#request-examples)
  - [For Developers](#for-developers)

<!-- /code_chunk_output -->

## Quick Start

1. Install [Docker v18.06.0+](https://docs.docker.com/get-docker/) and [Docker Compose v3.7+](https://docs.docker.com/compose/install/). You can check which version, if any, is already installed by running `docker --version` and `docker-compose --version`. Any versions greater than or equal to the ones stated here should do.
2. Clone this consents repository:
  ```
  git clone https://github.com/dycons/consents.git
  cd consents
  ```
3. Set up the build environment by generating a `.env` file for Docker Compose to use:
  ```
  cp default.env .env
  ```
4. From the project root directory, run the test server. It is presently configured to run on port `0.0.0.0:3000`. If you prefer to run the development container, replace `test` with `dev` in all following instructions.
  ```
  docker-compose -f docker/docker-compose.test.yml up
  ```
5. In a new shell (or the same shell if `docker-compose up --detach` was run), run the migrations script to prepare the database for use. Run it with the `-t` flag for the `test` instance and the `-d` flag for the `dev` instance.
  ```
  ./docker/migrate.sh -t
  ```
6. Send a request to the consents-service from Postman (import `./tests/consents-service.postman_collection.json` to run the collection against the `./tests/postman-data.csv` data file). Alternately, `curl` a request from the command line.

## Stack

- [Docker](https://www.docker.com/) is used for containerizing the service and its dependencies.
- [PostgreSQL](https://www.postgresql.org/) database backend
- [Go](https://golang.org/) (Golang) backend
- Go [mod](https://blog.golang.org/using-go-modules) is used for dependency management, similarly to how `dep` was used in the past.
- [Swagger/OpenAPI 2.0](https://swagger.io/specification/v2/) is used to specify the API.
- [Postman](https://www.postman.com/) and the CLI runner [Newman](https://learning.postman.com/docs/postman/collection-runs/command-line-integration-with-newman/) are used for end-to-end API testing.
- [TravisCI](https://travis-ci.org/) is used for continuous integration, testing both the service build and, in conjunction with Newman, the API.
- [Go-swagger](https://goswagger.io/) auto-generates boilerplate Go code from a `swagger.yml` API definition. [Swagger](https://swagger.io/) tooling is based on the [OpenAPI](https://www.openapis.org/) specification.
- Gobuffalo [pop](https://github.com/gobuffalo/pop) is used as an orm-like for interfacing between the go code and the sqlite3 database. The `soda` CLI auto-generates boilerplate Go code for models and migrations, as well as performing migrations on the database. `fizz` files are used for defining database migrations in a Go-like syntax (see [syntax documentation](https://gobuffalo.io/en/docs/db/fizz/).)
- Gobuffalo [validate](https://github.com/gobuffalo/validate) is a framework used for writing custom validators. Some of their validators in the `validators` package are used as-is.

### Installing the stack

The `consents-app` image built from `docker/test.Dockerfile` depends upon the `consents-deps` image built from `docker/deps.Dockerfile`, the image containing most of the stack and package dependencies for the app. These images have been split so that the `consents-app` build can be tested by TravisCI following every git push without time spent building the dependencies.

The stack and project dependencies image `consents-deps` can be pulled from `dyconsent/consents-deps`. Alternately, it can be built and run locally with the following commands:
  ```
  docker build -t <username>/consents-deps -f ./docker/deps.Dockerfile
  docker run -it --rm <username>/consents-deps
  ```

If you would like to push your altered image of `consents-deps`, consider using the `./docker/push-image.sh` script provided to quickly semantically version the image. Run `./docker/push-image.sh -h` for usage instructions.
  ```
  docker login && ./docker/push-image.sh -f ./docker/deps.Dockerfile -u <username> consents-deps <patch>
  ```

## Installing the service

For containerized installation instructions, please see the [Quick Start](#quick-start).

If you are interested in attempting a non-containerized installation, the following files contain build instructions for Docker, Docker Compose, and Travis CI. They may equally be followed manually to accomplish a manual installation of the stack and service:
- `docker/deps.Dockerfile` for the installation of the stack and of most other project dependencies
- `docker/test.Dockerfile` for the installation of the app, along with
- `docker/docker-compose.test.yml` and `default.env` for the configuration of the service, database, and build-time environment
- `travis.yml` for running the service and its associated end-to-end API tests

## Running The Service

From the project root directory, run the server. It is presently configured to run on port `0.0.0.0:3000`.
  ```
  docker-compose -f docker/docker-compose.test.yml up
  ```

### Request examples

If you have [Postman](https://www.postman.com/downloads/) installed, the quickest way to test the go-model-service API is to import `./tests/consents-service.postman_collection.json` and run the collection against the `./tests/postman-data.csv` data file.

## For Developers

Note that most auto-generated code, including binary files, are intentionally excluded from this repository.

If you would like to learn more about using this stack, especially with regards to its code-generation components, please take a look at the Go Model Service [DEVELOPER-GUIDE.md](https://github.com/CanDIG/go-model-service/blob/main/docs/DEVELOPER-GUIDE.md).
