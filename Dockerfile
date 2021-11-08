FROM golang:1.17-alpine3.14 AS builder

ENV \
    GOOS=linux \
    GOARCH=amd64

RUN \
    apk update && \
    apk add make  && \
    rm -rf /var/cache/apk/*

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
RUN make flags='-ldflags="-w -s"' build


FROM scratch

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder --chown=appuser:appuser /build/bin/demo /app/server

# Use an unprivileged user.
USER appuser:appuser

EXPOSE 8080
ENTRYPOINT [ "/app/server" ]
