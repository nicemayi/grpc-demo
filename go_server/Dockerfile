# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o /main

##
## Deploy
##
# FROM gcr.io/distroless/base-debian10

# WORKDIR /

# COPY --from=build /docker-gs-ping /docker-gs-ping

# EXPOSE 8080

# USER nonroot:nonroot

# ENTRYPOINT ["/docker-gs-ping"]