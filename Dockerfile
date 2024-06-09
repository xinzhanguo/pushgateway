FROM golang:1.22.1 AS build-env

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /go/src/github.com/xinzhanguo/pushgateway
COPY . /go/src/github.com/xinzhanguo/pushgateway
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o pushgateway cmd/main.go

FROM alpine:latest
LABEL MAINTAINER="e <e@xinzhanguo.cn>"

COPY --from=build-env /go/src/github.com/xinzhanguo/pushgateway/pushgateway /opt/

WORKDIR /opt/
EXPOSE 9091
ENTRYPOINT ["./pushgateway"]