FROM golang:1.21

# 设置工作目录
WORKDIR /app

# 下载并安装 NSQ
RUN apt-get install -y wget
RUN wget https://s3.amazonaws.com/bitly-downloads/nsq/nsq-1.2.1.linux-amd64.go1.16.6.tar.gz

RUN tar -xzf  nsq-1.2.1.linux-amd64.go1.16.6.tar.gz
RUN cp nsq-1.2.1.linux-amd64.go1.16.6/bin/* /usr/local/bin/

# 暴露 NSQ 相关端口
EXPOSE 4150

# 设置启动命令
CMD ["nsqd"]