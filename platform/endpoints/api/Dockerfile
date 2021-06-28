FROM golang:1.16.5 as builder
WORKDIR /gochat

# Copy the Required Files
COPY ./platform/endpoints/api/ ./platform/endpoints/api/
COPY ./platform/internal/ ./platform/internal/

# Init and Run the Go Mod file
RUN go mod init gochat
RUN go mod tidy
RUN go mod download

# Build the Application
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./platform/endpoints/api/

# Now create an optimised image
FROM alpine:latest

# Create an App Directory to work from
RUN mkdir -p /app
WORKDIR /app

# Copy the Executable from the Build State
COPY --from=builder /gochat/api .

# Run the Executable
CMD /app/api