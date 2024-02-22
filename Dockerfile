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
ARG BUILD_TIME
ARG GIT_COMMIT
ARG GIT_REF
ARG GIT_REF_NAME
ARG GIT_REF_TYPE

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./main.go
COPY pkg ./pkg
COPY vendor ./vendor
COPY views ./views
COPY assets ./assets
COPY locale ./locale
COPY --from=tailwind /app/assets/output.css ./assets/output.css

ENV CGO_ENABLED=0 GOOS=linux
RUN go build \
	-ldflags "-X 'main.buildTime=${BUILD_TIME}' -X 'main.gitCommit=${GIT_COMMIT}' -X 'main.gitRef=${GIT_REF}' -X 'main.gitRefName=${GIT_REF_NAME}' -X 'main.gitRefType=${GIT_REF_TYPE}'" \
	-o /workout-tracker ./

FROM alpine:latest

VOLUME /data
WORKDIR /app
COPY --from=gobuilder /workout-tracker ./workout-tracker
COPY locale /locale
ENV WT_LOCALE_DIRECTORY=/locale

WORKDIR /data
ENTRYPOINT ["/app/workout-tracker"]
EXPOSE 8080
