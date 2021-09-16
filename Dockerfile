FROM golang:1.17-alpine

WORKDIR /server

COPY ./server /server/

RUN go build -o /inn-server 

EXPOSE 8081

CMD ["/inn-server"]