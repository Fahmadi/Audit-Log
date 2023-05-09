ARG HUB

FROM ${HUB}debian:10-slim AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y curl

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz

COPY ./db/migration /app/migrations

CMD ./migrate -path ./migrations -database $DB_URI up
