FROM golang:1.16.15-alpine3.15 as builder
# Set up dependencies
ENV PACKAGES make git libc-dev bash gcc
ENV GO111MODULE  on
ARG GOPROXY=https://goproxy.cn,direct

COPY  . $GOPATH/src
WORKDIR $GOPATH/src

# Install minimum necessary dependencies, build binary
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk add --no-cache $PACKAGES && make all

FROM alpine:3.15
ENV TZ           Asia/Shanghai
WORKDIR /root/
COPY --from=builder /go/src/icndev-server /usr/local/bin
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk add --no-cache tzdata && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

CMD ["icndev-server", "start"]