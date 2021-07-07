# docker build --tag lenarbatdalov/golang-gin .
# docker images
# docker run -p 5000:5000 lenarbatdalov/golang-gin

FROM golang:latest

LABEL maintainer="Lenar <developer.lenar@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 5000

RUN go build

RUN find . -name "*.go" -type f -delete

EXPOSE $PORT

CMD ["./learn"]
