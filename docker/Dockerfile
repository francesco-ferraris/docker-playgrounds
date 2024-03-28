FROM golang:1.22

WORKDIR /app

# Get dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build sources
COPY *.go ./
RUN go build -o /docker-playgrounds

EXPOSE 8080
ENTRYPOINT ["/docker-playgrounds"]
