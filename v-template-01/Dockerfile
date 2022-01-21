FROM golang:latest

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR go/src/v-template-01
COPY . .

RUN go mod tidy
RUN go build -o server


EXPOSE 8081
ENTRYPOINT ./server