FROM node:24-alpine AS frontend

RUN apk --no-cache add make

WORKDIR /app

COPY frontend ./frontend
RUN cd frontend \
    && npm ci

COPY Makefile ./
COPY assets ./assets
COPY views ./views
COPY translations ./translations

RUN make build-frontend

FROM golang:1.25.1-alpine AS backend
ARG BUILD_TIME
ARG GIT_COMMIT
ARG GIT_REF
ARG GIT_REF_NAME
ARG GIT_REF_TYPE

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg
COPY vendor ./vendor
COPY views ./views
COPY assets ./assets
COPY translations ./translations
COPY --from=frontend /app/assets ./assets

ENV CGO_ENABLED=0
RUN go build \
    -ldflags "-X 'main.buildTime=${BUILD_TIME}' -X 'main.gitCommit=${GIT_COMMIT}' -X 'main.gitRef=${GIT_REF}' -X 'main.gitRefName=${GIT_REF_NAME}' -X 'main.gitRefType=${GIT_REF_TYPE}'" \
    -o /commands/ ./cmd/...

FROM alpine:3

RUN apk add --no-cache tzdata curl
COPY --from=backend /commands/* /app/

VOLUME /data /imports
WORKDIR /data
HEALTHCHECK CMD curl -sSLf http://localhost:8080/health
ENTRYPOINT ["/app/workout-tracker"]
EXPOSE 8080
