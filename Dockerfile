FROM golang:1.17-alpine3.9

MAINTAINER freedoow@qq.com

ENV TZ=Asia/Shanghai
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

WORKDIR /usr/build
COPY . .

RUN go mod download
RUN go build -o course_go .

FROM alpine:latest
WORKDIR /usr/dev/
RUN mkdir /test
WORKDIR /test

COPY --from=builder /usr/build/course_go .
ENTRYPOINT [ "./course_go" ]