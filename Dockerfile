# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code
COPY *.go ./
COPY model ./model
COPY handler ./handler
COPY docs ./docs
COPY database ./database

# Build
RUN go build -o /todolist

EXPOSE 8080

# Run
CMD [ "/todolist" ]
