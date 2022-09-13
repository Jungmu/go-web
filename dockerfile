# 1
FROM golang:alpine as builder
WORKDIR /root/app
COPY ./go-web .
COPY ./.env .
COPY ./assets ./assets
COPY ./templates ./templates
# 2
FROM alpine:latest
WORKDIR /go-web
COPY --from=builder /root/app ./
CMD [ "./go-web" ]