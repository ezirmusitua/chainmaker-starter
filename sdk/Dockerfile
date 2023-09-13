FROM golang:1.21-bookworm as build
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY . .
RUN go build -o /app/bin/sdk.bin

FROM debian:bookworm
EXPOSE 8080
WORKDIR /app
COPY --from=build /app/bin /app
ENTRYPOINT ["/app/sdk.bin", "/configs/sdk_config.yml"]
