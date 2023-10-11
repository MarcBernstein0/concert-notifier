FROM golang:1.21 AS BuildStage

WORKDIR /app 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o concert-notifier main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

EXPOSE 8080

EXPOSE $PORT

WORKDIR /root
COPY --from=BuildStage /app/concert-notifier ./
ENTRYPOINT [ "./concert-notifier" ]