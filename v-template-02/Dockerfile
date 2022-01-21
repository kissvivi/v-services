FROM golang:1.17-alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR go/src/v-template-02
COPY . .

RUN go mod tidy
RUN go build -o server


EXPOSE 8082
ENTRYPOINT ./server