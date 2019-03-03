FROM golang:alpine AS builder
RUN mkdir /server 
ADD /server /server 
WORKDIR /server
RUN go build -o server .
RUN chmod 777 server   
FROM alpine
COPY --from=builder server server
RUN chmod 777 server
CMD ["./server/server"]


