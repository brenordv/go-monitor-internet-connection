FROM golang:1.16-alpine
WORKDIR /app
COPY . ./
COPY ./runtime.config.json /runtime.config.json

RUN go mod download
RUN go build -o /conn-monitor ./cmd/conn_monitor

CMD ["/conn-monitor"]