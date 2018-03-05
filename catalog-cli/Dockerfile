FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD catalog.json /app/catalog.json
ADD catalog-cli /app/catalog-cli

CMD ["./catalog-cli"]
