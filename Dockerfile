FROM golang:1.16.15-alpine3.15 as builder
# Set up dependencies
ENV PACKAGES make git libc-dev bash gcc
ENV GO111MODULE  on
ARG GOPROXY=http://192.168.0.60:8081/repository/go-bianjie/,http://nexus.bianjie.ai/repository/golang-group,https://goproxy.cn,direct
ARG APKPROXY=http://mirrors.ustc.edu.cn/alpine

COPY  . $GOPATH/src
WORKDIR $GOPATH/src

# Install minimum necessary dependencies, build binary
RUN sed -i "s+http://dl-cdn.alpinelinux.org/alpine+${APKPROXY}+g" /etc/apk/repositories && apk add --no-cache $PACKAGES && make all

FROM alpine:3.15
ENV TZ           Asia/Shanghai
WORKDIR /root/
COPY --from=builder /go/src/icndev-server /usr/local/bin
RUN apk add --no-cache tzdata && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

CMD ["icndev-server", "start"]