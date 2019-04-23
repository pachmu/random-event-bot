FROM golang:onbuild

RUN go get github.com/go-telegram-bot-api/telegram-bot-api

RUN mkdir -p  /go/src/random_event_bot
WORKDIR /go/src/random_event_bot

ADD main.go /go/src/random_event_bot
ADD kudago_client.go /go/src/random_event_bot
ADD handler.go /go/src/random_event_bot
ADD config.go /go/src/random_event_bot
ADD config.yml /go/src/random_event_bot

RUN cd /go/src/random_event_bot && go build -i -o bot random_event_bot

ENTRYPOINT [ "/go/src/random_event_bot/bot", ">/go/src/random_event_bot/log" ]