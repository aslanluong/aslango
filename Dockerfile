FROM golang:1.16.2-alpine AS builder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN go build -o cmd/ cmd/main.go

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/cmd/main /app/
WORKDIR /app
CMD ["./main"]