### This is the main Dockerfile for the go-model-service app,
### excluding the database which is built seperately.
# Build and run from project root with either:
# (1)   docker-compose -f docker/docker-compose.dev.yml up
# (2)   # Comment out ENTRYPOINT if running interactively (-it)
#       docker build -t dyconsent/consents-app-dev --build-arg API_PATH=/go/src/github.com/dycons/consents/consents-service/api .
#       docker run -it --rm dyconsent/consents-app-dev

# Modify this line if you want to use a different stack/dependencies-image
FROM dyconsent/consents-deps AS consents-app-dev

ARG API_PATH

WORKDIR /go/src/github.com/dycons/consents
COPY . .

# Check that all dependencies are present.
# See `./consents-deps.Dockerfile` for more information.
RUN go mod download

# Swagger generate the boilerplate code necessary for handling API requests
# from the consents-service/api/swagger.yml template file.
# This will generate a server named consents-service. The name is important for
# maintaining compatibility with the configure_consents_service.go middleware
# configuration file.
RUN cd "$API_PATH" && pwd && swagger generate server -A consents-service swagger.yaml

# Run a script to generate resource-specific request handlers for middleware,
# from the generic handlers defined in the consents-service/api/generics package,
# using the CanDIG-maintained CLI tool genny
#RUN "$API_PATH"/generate_handlers.sh

# Now that all the necessary boilerplate code has been auto-generated, compile
# the server
RUN go build -gcflags="all=-N -l" -o ./app "$API_PATH"/cmd/consents-service-server/main.go

# Install delve, the debugger for go
RUN go get github.com/go-delve/delve/cmd/dlv

# Run the consents service
# Can attach the debugger (at port :2345) by running ./dev/dlv.sh from within the container
EXPOSE 3001
ENTRYPOINT ./app --port=3001 --host=0.0.0.0
