# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.13-alpine base image
FROM golang:1.18-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum
# files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the
# container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# The recommended approach from docker according to their docs on Control
# startup and shutdown order in Compose is to download wait-for-it.sh which takes
# in the domain:port to poll and then executes the next set of commands
# if successful.
COPY wait-for-it.sh wait-for-it.sh 
RUN chmod +x wait-for-it.sh

# Run the executable
ENTRYPOINT [ "/bin/bash", "-c" ]
CMD ["./wait-for-it.sh" , "aerospike:3000" , "--strict" , "--timeout=300" , "--" , "./main"]