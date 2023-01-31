FROM golang:1.15

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY go.mod ./
COPY go.sum ./

RUN go mod download -x

# Bundle app source
COPY . .
