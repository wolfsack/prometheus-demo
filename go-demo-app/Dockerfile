FROM golang:1.17 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

COPY *.go .

# Build the Go app
RUN go build -o ./out/exec ./main.go

# Start fresh from a smaller image
FROM ubuntu:20.04

WORKDIR /app

COPY --from=build /app/out/exec /app/exec

EXPOSE 8090

# Run the binary program produced by `go install`
CMD [ "/app/exec" ]