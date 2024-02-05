FROM golang:latest AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /web-service-api

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /web-service-api /web-service-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/web-service-api"]