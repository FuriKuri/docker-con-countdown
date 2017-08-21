FROM golang
ADD . /go/src/github.com/furikuri/docker-con-countdown
WORKDIR /go/src/github.com/furikuri/docker-con-countdown
RUN go get ./
RUN go build

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/bin/docker-con-countdown .
ENTRYPOINT ["./docker-con-countdown"]