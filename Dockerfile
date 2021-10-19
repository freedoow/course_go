## 依赖基础镜像
FROM golang:alpine
MAINTAINER freedoow@qq.com
## 设置环境
ENV TZ=Asia/Shanghai
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
##多段构建
WORKDIR /usr/app
COPY . .
#bulid
RUN go build -o httpserver .
##真正运行的地方
WORKDIR /bin/
RUN mkdir src .
RUN cp /root/usr/app/httpserver ./src
EXPOSE 8083

CMD ["/bin/src/httpserver"]
