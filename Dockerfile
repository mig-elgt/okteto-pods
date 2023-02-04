FROM golang:buster as builder
WORKDIR /app
ADD . .
RUN go test --cover -v ./...
RUN CGO_ENABLED=0 go build -o /usr/local/bin/app ./cmd/pods/main.go

FROM alpine:latest  
RUN apk --no-cache --update add ca-certificates
COPY --from=builder /usr/local/bin/app /usr/local/bin/app
RUN chmod +x /usr/local/bin/app
EXPOSE 8080

CMD ["app"]

