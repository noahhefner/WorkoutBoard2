FROM golang:1.23

WORKDIR /usr/src/app

COPY ./golang .

RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 3333

CMD ["app"]
