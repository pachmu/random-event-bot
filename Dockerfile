FROM golang:1.12 as build_base

WORKDIR /random-event-bot
COPY go.mod go.sum ./

RUN go mod download

FROM build_base as builder

WORKDIR /random-event-bot
COPY . .

RUN go build

FROM gcr.io/distroless/base

COPY --from=builder /random-event-bot/random-event-bot /
COPY --from=builder /random-event-bot/config.yml /etc/random-event-bot/config.yml

ENTRYPOINT ["/random-event-bot"]
CMD ["-config", "/etc/random-event-bot/config.yml"]
