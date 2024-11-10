FROM golang:1.23.2-alpine AS base

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

FROM base AS dev

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY . .

CMD ["CompileDaemon", "--build=go build -o echo-server ./server.go", "--command=./echo-server"]

FROM base AS builder

COPY . .

RUN go build -o echo-server ./main.go

FROM alpine:3.18 AS release

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/echo-server /usr/local/bin/echo-server

EXPOSE 8080

USER appuser

CMD ["echo-server"]

FROM node:22.11.0 AS node-base

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends libatomic1 && \
    rm -rf /var/lib/apt/lists/*

COPY package*json tsconfig.json vite.config.ts ./
RUN npm ci

FROM node-base AS node-dev

COPY . .

EXPOSE 3000

CMD ["npm", "run", "dev"]
