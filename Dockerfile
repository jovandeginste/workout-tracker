FROM node:alpine as tailwind

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install
COPY tailwind.config.js ./tailwind.config.js
COPY main.css ./main.css
COPY pkg ./pkg
COPY views ./views
COPY assets ./assets

RUN npx tailwindcss -i ./main.css -o ./assets/output.css

FROM golang:alpine as gobuilder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./main.go
COPY pkg ./pkg
COPY vendor ./vendor
COPY views ./views
COPY assets ./assets
COPY --from=tailwind /app/assets/output.css ./assets/output.css

RUN CGO_ENABLED=0 GOOS=linux go build -o /workouts

FROM alpine:latest

VOLUME /data
WORKDIR /app
COPY --from=gobuilder /workouts ./workouts

WORKDIR /data
ENTRYPOINT ["/app/workout-tracker"]
EXPOSE 8080
