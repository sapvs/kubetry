FROM golang:1.23.1-alpine3.20 AS builder
ARG MOD
WORKDIR /app
COPY ./pkg .
RUN go mod download
RUN go build -o main ./${MOD}/${MOD}.go

FROM alpine:3.20
ARG MOD
WORKDIR /app
COPY --from=builder /app/main . 

CMD [ "/app/main" ]