FROM golang:1.18 AS builder
WORKDIR /app
COPY . /app
RUN go mod tidy && go build

FROM alpine:latest
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /app
COPY --from=builder /app/mytest /app
EXPOSE 8080
ENTRYPOINT ./mytest
