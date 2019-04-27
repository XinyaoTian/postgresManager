# Source Image
FROM golang:latest
# Author
MAINTAINER LeonTian "leontian1024@gmail.com"
# Set working director
WORKDIR $GOPATH/src/github.com/XinyaoTian/postgresManager
# Add golang project from os into container
ADD . $GOPATH/src/github.com/XinyaoTian/postgresManager
# import go packages
RUN go get github.com/lib/pq
RUN go get github.com/julienschmidt/httprouter
# build golang exe
RUN go build .
# expose port
EXPOSE 9090
# run docker conmmand
ENTRYPOINT ["./postgresManager"]
