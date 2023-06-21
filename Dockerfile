FROM golang:alpine as builder
ARG LDFLAGS=""

RUN apk --update --no-cache add git build-base gcc

COPY . /build
WORKDIR /build

RUN go build -o ./jtso -ldflags "${LDFLAGS}" ./main.go 

FROM alpine:latest


RUN apk update --no-cache 

USER 0
COPY --from=builder /build/jtso /

EXPOSE 8081

ENTRYPOINT ["./jtso"]