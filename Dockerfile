FROM golang:1.20 as builder


WORKDIR /app


COPY . .


RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest


RUN apk --no-cache add ca-certificates

WORKDIR /root/


COPY --from=builder /app/main .
COPY --from=builder /app/private_key.pem .


EXPOSE 8080


CMD ["./main"]