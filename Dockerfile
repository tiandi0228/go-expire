FROM golang:1.17.1 AS builder
LABEL maintainer="Hong Cha"

ENV APPNAME go-expire
ENV BUILD_DIR /tmp/${APPNAME}
ENV GOBIN ${BUILD_DIR}/${APPNAME}


ARG GITEE_TOKEN
ENV GITEE_TOKEN ${GITEE_TOKEN}
ENV GOPROXY https://goproxy.cn,https://goproxy.io,direct

# Build
COPY . ${BUILD_DIR}
WORKDIR ${BUILD_DIR}

RUN env && \
  go mod tidy -compat=1.17 && \
  make build && \
  cp ${GOBIN} /usr/bin/go-expire

# Stage2
FROM ubuntu:focal

ADD http://mirrors.cloud.tencent.com/repo/ubuntu20_sources.list /etc/apt/sources.list
ENV TZ "Asia/Shanghai"

RUN apt update && apt install -y tzdata curl && \
  ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ >/etc/timezone

# Test CA certificates update and timezone configure
# RUN curl -I https://www.taobao.com/ && date

COPY --from=builder /usr/bin/go-expire /bin/go-expire
ENTRYPOINT ["/bin/go-expire"]