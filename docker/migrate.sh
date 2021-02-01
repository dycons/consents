#!/bin/bash

################################################################################
# Help                                                                         #
################################################################################
help ()
{
   # Display Help
   echo "Run migrations on the database out of a Docker container."
   echo "Default behaviour is to assume that this is a development (dev) migration."
   echo
   echo "Usage:"
   echo "   ./docker/migrate.sh [options]"
   echo "Options:"
   echo "   -h      Display this help text"
   echo "   -d      Migrate the dev service (this is the default behaviour)."
   echo "   -t      Migrate the test service."
   echo "   -f      Docker-compose filename. Default: ./docker/docker-compose.yaml"
   echo "   -s      Docker service container to exec the migration in. Default: consents"
   echo
}

################################################################################
################################################################################
# Main program                                                                 #
################################################################################
################################################################################

# Default docker-compose filename. Can overwrite with -f argument.
composefile="./docker/docker-compose.yaml"
# Default service name to exec the migration in. Can overwrite with -c argument.
service="consents"

# Optionally overwrite image version, Dockerfile, or Docker username
while getopts ":hdtf:s:" opt; do
  case $opt in
    h)  help
        exit
        ;;
    d)  service="consents-dev"
        ;;
    t)  service="consents"
        ;;
    f)  composefile="$OPTARG"
        ;;
	s)  service="$OPTARG"
		;;
    \?) echo "Invalid option -$OPTARG" >&2
        ;;
  esac
done

# Create a Postgres development database and migrate it to the schema
# defined in the consents-service/data directory, using the soda tool from pop.
# The willingness of the postgres socket to accept TCP connection does not
# necessarily indicate database readiness. Hence, the pg_isready check below.
# Read more about containerized postgres readiness:
# 	https://github.com/docker-library/postgres/issues/146

echo "composefile: " $composefile
echo "service: " $service
docker-compose -f $composefile exec $service sh -c \
	'until pg_isready --timeout=60 --username=$POSTGRES_USER --host=$POSTGRES_HOST; \
	do sleep 1; \
	done &&\
	soda migrate up -c ./database.yml -e development -p consents-service/data/migrations'
