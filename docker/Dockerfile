# Build the manager binary
FROM golang:1.19 as builder

WORKDIR /app

COPY go.mod go.mod

COPY cmd/iface cmd/iface
COPY cmd/utils cmd/utils
COPY cmd/wxBot cmd/wxBot

RUN mkdir -p /app/tmp/wxBotCache/botCacheFile \
&& mkdir -p /app/tmp/wxBotCache/botLog \
&& go env -w GO111MODULE=on \
&& go env -w GOPROXY=https://goproxy.cn,direct \
&& go env -w CGO_ENABLED=0 \
&& go env -w GOOS=linux \
&& go mod tidy \
&& go build -a -o SamgeWxBot cmd/wxBot/main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM alpine:3.9
WORKDIR /app
COPY --from=builder /app/SamgeWxBot .

VOLUME /app/tmp/wxBotCache

EXPOSE 8080
ENTRYPOINT ["./SamgeWxBot"]
