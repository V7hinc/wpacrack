# 使用官方的Go镜像作为基础镜像
FROM golang:1.20 as builder

# 设置工作目录
WORKDIR /app

# 复制项目文件到容器中
COPY . .

# 编译Go应用
RUN go mod download
RUN go build -o /app/crack_linux ./cli/main.go

FROM aircrackng/git:2e2be0f

WORKDIR /app

COPY --from=builder /app/crack_linux /app/crack_linux

VOLUME /app/cap_file /app/wpa-dictionary

RUN echo "/app/crack_linux dictCrack -c /app/cap_file -p /app/wpa-dictionary -k \$key" > /app/start.sh; \
    cat /app/start.sh

CMD ["sh", "/app/start.sh" ]