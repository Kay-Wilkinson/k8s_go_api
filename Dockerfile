FROM golang:1.19.0

WORKDIR /usr/src/app

# Install dep to have hot reload on dev server
RUN go install github.com/cosmtrek/air@latest

COPY . .

# Install and clean up dependencies with go tidy
RUN go mod tidy