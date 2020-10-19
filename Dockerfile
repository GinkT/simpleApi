FROM golang:latest 
 
WORKDIR /usr/local/go/src/main

COPY ./ /usr/local/go/src/main

RUN go build -o main .

EXPOSE 8081

ENTRYPOINT ["/usr/local/go/src/main/main"]