FROM golang:1.22-alpine AS builder

WORKDIR /build
COPY . .
RUN go build -o app cmd/*.go 

FROM alpine:latest

COPY --from=builder /build/app /app

EXPOSE 8080

CMD [ "/app" ]