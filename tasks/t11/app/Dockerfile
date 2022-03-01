FROM golang:1.17

WORKDIR /go/src/app

COPY . .

RUN go get github.com/jmoiron/sqlx
RUN go get github.com/lib/pq

RUN go build -o app cmd/api_server/main.go
CMD sleep 3 && ./app