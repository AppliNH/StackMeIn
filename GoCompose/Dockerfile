#FROM golang:alpine
FROM docker:dind
# FROM jpetazzo/dind
# RUN apt update
# RUN add-apt-repository ppa:gophers/archive
# RUN apt install golang-1.14 -y
RUN apk add go
ENV GO111MODULE=on

WORKDIR /app

COPY ./go.mod .


RUN go mod download

COPY . .

# Build the Go app
RUN go build -o gocompose .

EXPOSE 1997

CMD ["./gocompose"]