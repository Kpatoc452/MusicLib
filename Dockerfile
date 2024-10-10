FROM golang:1.22 AS builder

RUN go env
WORKDIR /build
COPY . .
RUN go get 
RUN go build -o app cmd/*.go 

FROM alpine:3.20.3

COPY --from=builder /build/app /app

EXPOSE 8080

CMD [ "sh" ]
