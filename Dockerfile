FROM golang:1.12.3-alpine3.9

MAINTAINER freedoow@qq.com

ENV TZ=Asia/Shanghai
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

WORKDIR /usr/k8s
COPY . .

RUN go mod download
RUN go build -o course_go .

WORKDIR /dev/
RUN mkdir src .
RUN cp /usr/k8s/course_go ./src
EXPOSE 8083

CMD ["/usr/dev/src/course_go"]
