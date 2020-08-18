#!/bin/bash

# TODO make this a makefile instead?

# Create a Postgres development database and migrate it to the schema
# defined in the consents-service/data directory, using the soda tool from pop.
# The willingness of the postgres socket to accept TCP connection does not
# necessarily indicate database readiness. Hence, the pg_isready check below.
# Read more about containerized postgres readiness:
# 	https://github.com/docker-library/postgres/issues/146

docker-compose exec consents-webapp sh -c \
	"until pg_isready --timeout=60 --username=$POSTGRES_USER --host=database; \
	do sleep 1; \
	done &&\
	soda migrate up -c ./database.yml -e development -p consents-service/data/migrations"
