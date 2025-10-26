# Assuming that it is more likely that dev container features like PowerShell, Docker-in-Docker, and Go will change less frequently than app-specific parts like Node.js dependencies, this Dockerfile is designed to install those components first.
# This allows for faster rebuilds of the dev container when only app-specific parts change.
# The downside of this approach is that these dev container features must be installed here and cannot be installed via the devcontainer.json file, because the devcontainer.json file is processed after the Dockerfile.

ARG NODE_VERSION=YouForgotToSetIt

# Bookworm must be used, otherwise using the dev container feature "Docker in Docker" does not work
FROM mcr.microsoft.com/devcontainers/javascript-node:${NODE_VERSION}-bookworm

# Install PowerShell
RUN set -eux; \
    apt-get update -y; \
    curl -fsSL https://raw.githubusercontent.com/devcontainers/features/refs/heads/main/src/powershell/install.sh -o /tmp/install-pwsh.sh; \
    bash /tmp/install-pwsh.sh; \
    rm -f /tmp/install-pwsh.sh; \
    rm -rf /var/lib/apt/lists/*

# Install Docker-in-Docker (moby/docker engine + tools)
RUN set -eux; \
    apt-get update -y; \
    curl -fsSL https://raw.githubusercontent.com/devcontainers/features/refs/heads/main/src/docker-in-docker/install.sh -o /tmp/install-dind.sh; \
    bash /tmp/install-dind.sh; \
    rm -f /tmp/install-dind.sh; \
    rm -rf /var/lib/apt/lists/*

# Install Go (uses the GO_VERSION build-arg via the script's $VERSION variable)
ARG GO_VERSION=YouForgotToSetIt
RUN set -eux; \
    apt-get update -y; \
    curl -fsSL https://raw.githubusercontent.com/devcontainers/features/refs/heads/main/src/go/install.sh -o /tmp/install-go.sh; \
    VERSION="${GO_VERSION}" bash /tmp/install-go.sh; \
    rm -f /tmp/install-go.sh; \
    rm -rf /var/lib/apt/lists/*

# Ensure the Go binary directory is on PATH so the `go` command is available
# without requiring the full path (/usr/local/go/bin/go)
ENV PATH="/usr/local/go/bin:${PATH}"

# Install/restore Node.js dependencies so that they are cached in the Docker image
COPY ./frontend/package*.json /tmp/workout-tracker/npm/
WORKDIR /tmp/workout-tracker/npm/
RUN npm ci
RUN rm -rf /tmp/workout-tracker/npm
