# 使用 Golang 作为基础镜像
FROM golang:latest AS builder

# 设置 GOPROXY 环境变量
ENV GOPROXY=https://goproxy.cn

# 设置工作目录
WORKDIR /app

# 复制项目文件到工作目录
COPY . .

# 编译项目
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOARM=6 go build -o main .

# 创建一个轻量的镜像
FROM alpine:latest

# 换源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update

# 设置时区
RUN apk add -U tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime

# 设置工作目录
WORKDIR /app

# 从构建阶段复制构建好的文件到最终镜像
COPY --from=builder /app/main .

# 创建配置文件
RUN touch config.json

# 添加执行权限
RUN chmod +x /app/main

# 暴露端口
EXPOSE 1323

# 运行应用
CMD ["/app/main", "/app/config.json"]