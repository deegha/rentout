FROM golang:latest

RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on

ENV APP_HOME /go/src/rentoutapi
RUN mkdir -p "$APP_HOME"
ADD . /go/src/rentoutapi/
RUN go mod download

WORKDIR "$APP_HOME"
EXPOSE 8080
CMD ["bee", "run"]
