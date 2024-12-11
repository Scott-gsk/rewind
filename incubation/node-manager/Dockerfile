FROM golang:1.23.0 AS builder
WORKDIR /app/rewind

ADD . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN cd incubation/node-manager && \
    swag init --parseDependency
RUN cd incubation/node-manager && \
    go mod tidy && go mod vendor && go build -o /app/node-manager

FROM ubuntu
RUN apt-get update && \
    apt-get install -y curl pkg-config libssl-dev ca-certificates  &&\
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app/node-manager .

ADD ./incubation/node-manager/entrypoint.sh /entrypoint.sh
ADD ./incubation/node-manager/localize localize

RUN chmod 777 /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]



