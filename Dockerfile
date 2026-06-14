FROM ghcr.io/pnpm/pnpm:11 as builder
RUN pnpm runtime set node 24 -g

WORKDIR /app
COPY web/package.json .
COPY web/pnpm-lock.yaml .
RUN pnpm install --frozen-lockfile
COPY web/ ./
RUN pnpm build

FROM golang:1.26-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

COPY --from=builder /app/dist ./web/dist
COPY --from=builder /app/web.go ./web/

RUN go build -o antrian-bin .
