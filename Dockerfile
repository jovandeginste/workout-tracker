FROM node:alpine as frontend

WORKDIR /app

RUN apk --no-cache add make

COPY Makefile package.json package-lock.json ./
RUN make install-deps
COPY tailwind.config.js ./tailwind.config.js
COPY main.css ./main.css
COPY pkg ./pkg
COPY views ./views
COPY assets ./assets

RUN make build-dist build-tw

FROM golang:alpine as backend
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
COPY translations ./translations
COPY --from=frontend /app/assets/output.css ./assets/output.css
COPY --from=frontend /app/assets/dist ./assets/dist

ENV CGO_ENABLED=0
RUN go build \
	-ldflags "-X 'main.buildTime=${BUILD_TIME}' -X 'main.gitCommit=${GIT_COMMIT}' -X 'main.gitRef=${GIT_REF}' -X 'main.gitRefName=${GIT_REF_NAME}' -X 'main.gitRefType=${GIT_REF_TYPE}'" \
	-o /workout-tracker ./

FROM alpine:latest

RUN apk add --no-cache tzdata
VOLUME /data /imports
WORKDIR /app
COPY --from=backend /workout-tracker ./workout-tracker

WORKDIR /data
ENTRYPOINT ["/app/workout-tracker"]
EXPOSE 8080
