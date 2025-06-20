FROM golang:1.24.4-alpine3.22 AS builder

WORKDIR /app

#RUN go mod download

COPY ./* /go/src/
WORKDIR /go/src/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Create the final lightweight image
FROM alpine:3.22
WORKDIR /root/
COPY --from=builder /go/src/main .
EXPOSE 8080
CMD ["./main"]