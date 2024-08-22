FROM golang:1.21 as build

ARG DATA_PG_HOST
ENV DATA_PG_HOST=$DATA_PG_HOST

ENV GOOS=linux

RUN mkdir /app
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 \
    go build --ldflags "-s -w" -a -installsuffix cgo -o app .

FROM alpine:3.16

# Add non root user and certs
RUN apk --no-cache add ca-certificates curl \
    && addgroup -S app && adduser -S -g app app
RUN mkdir -p /home/app \
    && chown app /home/app
WORKDIR /home/app

COPY --from=build --chown=app /app/app .

USER app

EXPOSE 8080
EXPOSE 8081
EXPOSE 8082

CMD ["./app"]