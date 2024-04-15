# be build
FROM golang:1.20-bullseye AS be_builder

COPY . /build

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o helmut cmd/main.go

# copy to alpine
FROM alpine:latest

RUN apk add ca-certificates

COPY --from=be_builder /build/helmut /helmut

WORKDIR /

RUN mkdir charts
ADD charts charts/

ENV HELMUT="yes"

CMD sleep 15;/helmut