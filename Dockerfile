FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./main.go
COPY pkg ./pkg
COPY vendor ./vendor
COPY views ./views
COPY assets ./assets

RUN CGO_ENABLED=0 GOOS=linux go build -o /workouts

FROM alpine:latest

WORKDIR /app
COPY --from=builder /workouts ./workouts

ENTRYPOINT ["/app/workouts"]
EXPOSE 8080
