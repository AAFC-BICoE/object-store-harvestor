# - - - BUILD - - - #
FROM golang:alpine as builder
# Install alpine dependencies
RUN apk add --update --no-cache git
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
RUN CGO_ENABLED=0 go test /go/src/harvestor/... -v -coverprofile=/tmp/coverage.out
RUN go tool cover -html=/tmp/coverage.out
# Building Artifact
RUN go build -o main /go/src/harvestor

# - - - LOCAL - - - #
# Pushing Artifact to local
FROM alpine
RUN apk update
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /go/main /app/
COPY --from=builder /go/src/harvestor/harvestor_config.yml /app/
COPY --from=builder /tmp/coverage.out /tmp/
WORKDIR /app
RUN mkdir /tmp/data-test

# Release
ARG ARG_RELEASE_VERSION="0.001"
ENV RELEASE_VERSION "${ARG_RELEASE_VERSION}"

ARG ARG_DEPLOYMENT_ENV="Local"
ENV DEPLOYMENT_ENV "${ARG_DEPLOYMENT_ENV}"

CMD ["./main", "/app/harvestor_config.yml"]
