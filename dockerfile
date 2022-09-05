# 1
FROM golang:alpine as builder
WORKDIR /root/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o go-web -a main.go
RUN ls -al
WORKDIR /root/app/bin
RUN cp ../go-web .
RUN cp ../.env .

# 2
FROM alpine:latest
WORKDIR /app
COPY --from=builder /root/app/bin ./
EXPOSE 6666
CMD [ "./go-web" ]