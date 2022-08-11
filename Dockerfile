FROM golang:1.18-alpine AS builder

LABEL maintainer="StarkSim<gooda159753@163.com>"

# 在容器根目录创建 src 目录
WORKDIR /src

COPY ./go.mod .

COPY ./go.sum .

ENV GOPROXY="https://goproxy.cn"

RUN go mod download

COPY . .

RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-s -w" -o apiserver ./

FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /src/apiserver /app/
COPY --from=builder /src/migrations /app/migrations/

EXPOSE 8080

ENTRYPOINT ["./apiserver"]