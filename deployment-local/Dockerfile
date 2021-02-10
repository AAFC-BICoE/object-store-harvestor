# - - - BUILD - - - #
FROM golang:alpine as builder
# Install alpine dependencies
RUN apk add --update --no-cache git build-base
# updates
RUN apk update
# Preps
# making source folder for a build
RUN mkdir -p /go/src/harvestor
# copy assets
COPY harvestor /go/src/harvestor
# define working directory
WORKDIR /go
# Getting Project dependencies
RUN go get -v -d /go/src/harvestor/...

# - - - TESTS - - - #
# Running tests
RUN go get -v github.com/stretchr/testify/assert
RUN CGO_ENABLED=0 go test /go/src/harvestor/... -v -coverprofile cover.out
# Output html coverage report to /tmp/cover.html
RUN go tool cover -html=cover.out -o /tmp/cover.html
# When close to prod
# RUN go tool cover -func cover.out | grep total | awk '{print substr($3, 1, length($3)-1)}'
# And we need to post using curl to one of the services

# Building Artifact
RUN go build -o main /go/src/harvestor

# - - - LOCAL - - - #
# Pushing Artifact to local
FROM alpine
RUN apk add --update --no-cache sqlite
RUN apk update
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /go/main /app/
COPY --from=builder /go/src/harvestor/harvestor_config.yml /app/
# For now copy the cover report to release container
COPY --from=builder /tmp/cover.html /tmp/
WORKDIR /app
# folder for media files
RUN mkdir /tmp/data-test
# folder for test SQLite DB
RUN mkdir /tmp/db-test

# Release
ARG ARG_RELEASE_VERSION="0.001"
ENV RELEASE_VERSION "${ARG_RELEASE_VERSION}"

ARG ARG_DEPLOYMENT_ENV="Local"
ENV DEPLOYMENT_ENV "${ARG_DEPLOYMENT_ENV}"

CMD ["./main", "/app/harvestor_config.yml"]
