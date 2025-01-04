FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o smoke-test

EXPOSE 9000

CMD [ "./smoke-test" ]
