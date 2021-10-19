FROM golang:alpine
MAINTAINER freedoow@qq.com
ENV TZ=Asia/Shanghai
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

WORKDIR /usr/app
COPY . .
RUN go build -o httpserver .

WORKDIR /bin/
RUN mkdir src .
RUN cp /root/usr/app/httpserver ./src
EXPOSE 8083

CMD ["/bin/src/httpserver"]
