# BUILD stage
FROM alpine:edge as builder
RUN apk update
RUN apk upgrade
RUN apk add --update go gcc g++ git
WORKDIR /app

ENV GOPATH='/app' SRC_DIR="/app/src/PrayKyotoServer"

ADD . $SRC_DIR
RUN go get PrayKyotoServer
RUN CGO_ENABLE=1 GOOS=linux go install -a PrayKyotoServer
RUN mkdir -p /dist/app && mkdir -p /dist/db && cp /app/bin/PrayKyotoServer /dist/app/PrayKyotoServer

# DIST stage
FROM alpine:latest

COPY --from=builder /dist /

EXPOSE 8080
CMD ["/app/PrayKyotoServer"]
