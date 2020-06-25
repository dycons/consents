### This is the main Dockerfile for the go-model-service app,
### excluding the database which is built seperately.

# Modify this line if you want to use a different stack/dependencies-image
FROM katpavlov/consents-deps AS webapp

#ARG API_PATH

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
RUN cd "$API_PATH" && swagger generate server -A consents-service swagger.yml

# Run a script to generate resource-specific request handlers for middleware,
# from the generic handlers defined in the consents-service/api/generics package,
# using the CanDIG-maintained CLI tool genny
#RUN "$API_PATH"/generate_handlers.sh

# Now that all the necessary boilerplate code has been auto-generated, compile
# the server
#RUN go build -o ./main "$API_PATH"/cmd/consents-service-server/main.go

# Run the consents service
#EXPOSE 3000
#ENTRYPOINT ./main --port=3000 --host=0.0.0.0
