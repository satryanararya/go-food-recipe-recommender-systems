# syntax=docker/dockerfile:1

FROM golang:1.22.0 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-chefbot .

FROM gcr.io/distroless/base-debian11 AS build-release

WORKDIR /

COPY --from=build /go-chefbot .
COPY --from=build /app/.env .env

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/go-chefbot" ]