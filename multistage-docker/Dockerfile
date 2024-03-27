# Build the application from source
FROM golang:1.22 as build-stage
WORKDIR /app

# Get dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build sources
COPY *.go ./
RUN go build -o /docker-playgrounds

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12 AS build-release-stage

# Copy executable from build stage
WORKDIR /
COPY --from=build-stage /docker-playgrounds /docker-playgrounds

EXPOSE 8080

ENTRYPOINT ["/docker-playgrounds"]
