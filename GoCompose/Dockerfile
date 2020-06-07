FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

COPY ./go.mod .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o gocompose .

EXPOSE 1997

CMD ["./gocompose"]