FROM golang:1.17-alpine

ENV TZ="Europe/Moscow"

WORKDIR /app

COPY . ./

RUN go mod tidy

RUN go build -o /appBild

EXPOSE 7981

CMD [ "/appBild"]
