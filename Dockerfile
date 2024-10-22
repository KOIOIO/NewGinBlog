# 使用官方的 Golang 运行时作为基础镜像
FROM golang:1.22-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 Go Mod 和 Sum 文件以利用缓存
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目的 Go 源码到工作目录
COPY . .

# 构建项目
RUN go build -o ginblog .

# 使用一个轻量级的 Alpine Linux 镜像作为运行时镜像
FROM alpine:latest

# 安装必要的包（如 ca-certificates 用于 HTTPS 连接）
RUN apk add --no-cache ca-certificates

# 将构建好的二进制文件从构建阶段复制到运行时阶段
COPY --from=builder /app/ginblog /usr/local/bin/ginblog

# 复制配置文件（如果有的话，假设它在项目根目录下）
# COPY config.yaml /etc/ginblog/config.yaml

# 设置环境变量（这些也可以在 docker-compose.yml 中设置）
ENV APP_MODE=debug \
    HTTP_PORT=8080 \
    JWT_KEY=eqwr3425 \
    DB_HOST=db \
    DB_PORT=3306 \
    DB_USER=root \
    DB_PASSWORD=wwy040609 \
    DB_NAME=NewGinBlog

# 暴露 HTTP 端口
EXPOSE 8080

# 运行 GinBlog 应用
ENTRYPOINT ["/usr/local/bin/ginblog"]