FROM registry.cn-hangzhou.aliyuncs.com/gintest/golang:1.17-alpine as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN go mod tidy && go mod download
RUN CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -o gin_test_api


FROM registry.cn-hangzhou.aliyuncs.com/gintest/alpine:latest
WORKDIR /app
COPY config ./config
COPY cert ./cert
COPY --from=builder /app/gin_test_api .
RUN ln -snf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
ENV GIN_MODE=release
EXPOSE 8080
ENTRYPOINT ["./gin_test_api"]