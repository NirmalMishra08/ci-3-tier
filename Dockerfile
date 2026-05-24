FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/main .
COPY --from=builder /app/static ./static

RUN chmod +x ./main 

ENV PORT=8000

EXPOSE 8000

CMD ["./main"]