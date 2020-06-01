# Getting latest golang docker image from hub
FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
# build application
RUN go build -o main .
CMD ["/app/main"]