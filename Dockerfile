FROM golang:1.17-alpine

WORKDIR /app

COPY . ./

RUN go mod tidy

RUN go build -o /achieveService

CMD [ "/achieveService" ]
