# Antrian

Ini adalah website antrian sederhana

## Requirements
- Go >= 1.26

## Cara Menjalankan

```bash
git clone https://github.com/maroisa/antrian.git && cd antrian
go mod download

# Build frontend
cd web
pnpm install && pnpm build
cd ..

go build -o antrian .
./antrian
```

### Docker
```bash
docker compose up -d --build
```

## Cara Development

```bash
# Menjalankan Frontend
cd web
cp .env.example .env
pnpm install
pnpm dev

# Menjalankan Backend
go mod download
go run .
```
