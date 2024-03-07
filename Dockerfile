
FROM golang:1.20-alpine


RUN mkdir app

ADD . /app/

WORKDIR /app

RUN go mod tidy && go mod vendor

ENV ENVIRONMENT RELEASE

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/stepp-backend /app/src/

COPY . .

EXPOSE 8080

CMD ["./stepp-backend"]