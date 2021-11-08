FROM golang:1.17-alpine3.14 AS builder

RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "10001" \    
    "appuser"

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/server ./cmd/...


# TODO: Move to scratch, blocked by issue with M1
FROM alpine:3.14

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder --chown=appuser:appuser /build/bin/server /app/server

# Use an unprivileged user.
USER appuser:appuser

EXPOSE 8080
ENTRYPOINT [ "/app/server" ]
