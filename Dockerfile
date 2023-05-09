ARG BUILDER_IMAGE=golang:1.16.5

############################
# STEP 1 build executable binary
############################
FROM ${BUILDER_IMAGE} as builder

# Ensure ca-certficates are up to date
RUN update-ca-certificates

# Copy ssh keys for github.com
RUN mkdir -p -m 0600 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts

WORKDIR /go/src/app

# use modules
COPY go.mod go.sum ./

ENV GO111MODULE=on
RUN --mount=type=ssh git config --global url."ssh://git@github.com/conekta".insteadOf https://github.com/conekta && go mod download -x

# Copy all the Code and stuff to compile everything
COPY . .