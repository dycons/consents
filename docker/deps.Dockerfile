### This Dockerfile builds the stack & dependencies for the go-model-service

# Build from project root and push to Docker Hub:
# 	docker login && ./push-image.sh -b -f ./consents-deps.Dockerfile -u <username> gms-deps <patch>
# Alternately, build locally without pushing:
# 	docker build -t <username>/consents-deps -f ./consents-deps.Dockerfile
# Run resulting container from project root with:
# 	docker run -it --rm <username>/consents-deps

# Modify this line if you want to use a different stack-image
FROM golang as consents-deps

ENV GOPATH=/go

WORKDIR /go/src/github.com/dycons/consents
COPY ./go.mod ./go.sum ./

# Use the mod tool to fetch/cache all project import dependencies into 
# 	$GOPATH/pkg/mod
RUN go mod download

# Install Go-swagger (code-gen of boilerplate server Go code from OpenAPI definition)
RUN go install "$GOPATH"/pkg/mod/github.com/go-swagger/go-swagger@v0.25.0/cmd/swagger

# Install genny (code-gen solution for generics in Go)
#RUN go get github.com/CanDIG/genny

# Install pop (ORM-like for interfacing with the database backend)
# soda is the pop CLI
RUN go install "$GOPATH"/pkg/mod/github.com/gobuffalo/pop@v4.13.1+incompatible/soda

# Install Postgres client for interfacing with the database
RUN apt-get update &&\
	apt-get install -y postgresql-client
