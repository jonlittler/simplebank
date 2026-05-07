# Multi-Stage

# Build Stage
FROM golang:1.25-alpine3.23 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run State
FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
# COPY start.sh .
# COPY db/migration ./db/migration

# EXPOSE 8080 9090
EXPOSE 8080
CMD [ "/app/main" ]
# ENTRYPOINT [ "start.sh" ]
