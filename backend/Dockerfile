FROM golang:alpine

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Go to base to place certs inline with paths in .env
WORKDIR /

# Copy the certs

COPY localhost.crt /cert/
COPY localhost.key /cert/

# Move to working directory /build
WORKDIR /src/build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /src/dist

# Copy the .env 
COPY .env .

# Copy binary from build to main folder
RUN cp /src/build/main .

# Export necessary port
EXPOSE 8001

# Command to run when starting the container
CMD ["/src/dist/main"]