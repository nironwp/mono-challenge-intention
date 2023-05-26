FROM golang:1.18.1-alpine AS builder

# All these steps will be cached
RUN mkdir /app
WORKDIR /app

# Copy go.mod and go.sum files to workspace
COPY go.mod go.sum ./

# Get dependencies, will also be cached if you don't change mod/sum
RUN go mod download

RUN go mod tidy

# Copy the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o intent_service ./cmd/execute

FROM scratch

COPY --from=builder /app/intent_service /intent_service

EXPOSE 3001
# CMD
CMD ["/intent_service"]
