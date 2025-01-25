FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod download 
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/base-debian11
COPY --from=builder /go/bin/app /
CMD ["/app"]

