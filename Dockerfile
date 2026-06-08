FROM ghcr.io/pnpm/pnpm:11 AS builder
RUN pnpm runtime set node 24 -g

WORKDIR /app

COPY ./web/package*.json ./
COPY ./web/pnpm*.yaml ./
RUN pnpm install
COPY ./web ./
ENV CI=true
RUN pnpm build

FROM golang:1.26-alpine

WORKDIR /app
COPY go.mod .
RUN go mod download

COPY . .
COPY --from=builder /app/dist /app/web/dist
RUN go build -o antrian .
