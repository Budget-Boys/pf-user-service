FROM golang:1.24

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

RUN go build -v -o /usr/local/bin/app ./cmd/main.go

CMD sh -c "/wait-for-it.sh pf-discovery-server:8761 --timeout=60 --strict -- \
  /wait-for-it.sh pf-user-mysql:3306 --timeout=60 --strict -- \
  /wait-for-it.sh pf-user-redis:6379 --timeout=60 --strict -- \
  /usr/local/bin/app"
