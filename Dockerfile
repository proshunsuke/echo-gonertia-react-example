FROM node:22.11.0-alpine3.20 AS node-base

WORKDIR /app

COPY package*json tsconfig.json vite.config.ts ./
RUN npm ci

FROM node-base AS node-dev

COPY . .

EXPOSE 3000

CMD ["npm", "run", "dev"]

FROM node:22.11.0 AS node-dev-test

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends libatomic1

RUN apt-get clean && rm -rf /var/lib/apt/lists/*

COPY . .

RUN npx playwright install && \
    npx playwright install-deps

FROM node-base AS node-builder

COPY . .

RUN npm run build

FROM golang:1.23.2-alpine3.20 AS base

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

FROM base AS dev

RUN go install github.com/air-verse/air@latest

COPY . .

CMD ["air", "-c", ".air.toml"]

FROM base AS builder

COPY . .

RUN go build -ldflags="-s -w" -o echo-server ./server.go

FROM alpine:3.20 AS release

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/echo-server /usr/local/bin/echo-server
COPY --from=builder /app/resources/views/root.html ./resources/views/root.html
COPY --from=node-builder /app/public/build/.vite/manifest.json ./public/build/manifest.json

EXPOSE 1323

USER appuser

CMD ["echo-server"]

FROM nginx:1.26.2 AS nginx-base

COPY nginx/nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

FROM nginx-base AS nginx-release

COPY --from=node-builder /app/public /usr/share/nginx/html/public
