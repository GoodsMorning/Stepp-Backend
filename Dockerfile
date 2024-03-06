
FROM golang:1.20


RUN mkdir app

ADD . /app/

WORKDIR /app

RUN go mod tidy && go mod vendor

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/stepp-backend /app/src/

EXPOSE 8080

CMD ["./stepp-backend"]