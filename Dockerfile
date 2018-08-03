FROM golang:1.9

WORKDIR /go/src/web

# 复制文件
COPY webapp /go/src/web

# 暴露端口
EXPOSE 3000

CMD ["/go/src/web/webapp"]