FROM golang:1.16.3-alpine3.13 AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

ARG service

# Build the application
RUN go build -o main ./services/auth/main.go

############################
# STEP 2 build a small image
############################
FROM alpine:3.12.3

COPY --from=builder /build/main /

# Command to run the executable
ENTRYPOINT ["/main"]