# BUILD stage
FROM golang:latest as builder

ENV SRC_DIR="$GOPATH/src/PrayKyotoServer"

WORKDIR $SRC_DIR 
ADD . $SRC_DIR
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/mattn/go-sqlite3
RUN go get -u github.com/jinzhu/gorm
RUN go build .
RUN mkdir -p /dist/app && mkdir -p /dist/db && cp main /dist/app/main

# DIST stage
FROM alpine:latest

COPY --from=builder /dist /

EXPOSE 8080
CMD ["/app/main"]
