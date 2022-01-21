FROM golang:latest

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR go/src/v-template-etcd
COPY . .

RUN go mod tidy
RUN go build -o server


EXPOSE 8888
ENTRYPOINT ./server