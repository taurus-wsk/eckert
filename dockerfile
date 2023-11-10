# 编译阶段
FROM golang:1.17
WORKDIR /app

# 将 Golang 源代码复制到容器中
COPY . .

# 执行编译命令生成二进制文件
RUN go build -o eckert

# 暴露容器的端口
EXPOSE 8000

# 容器启动时要执行的命令
# CMD ["/app/eckert"]
