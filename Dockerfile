FROM golang:1.21.3 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o news-service cmd/news-service/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/news-service /app/news-service
COPY config/docker.yaml /app/docker.yaml
WORKDIR /app
CMD ["./news-service", "-config=docker.yaml"]