# Antrian

Ini adalah website antrian sederhana

## Requirements
- Go >= 1.26

## Cara Menjalankan
```bash
# -- Lokal --
git clone https://github.com/maroisa/antrian.git && cd antrian
go mod download

# Build frontend
cd web
pnpm install && pnpm build
cd ..

go build -o antrian .
./antrian

# -- atau dengan Docker --
docker compose up -d --build
```

> [!NOTE]
> Untuk pemutar suara, harus menggunakan Chrome dengan cara:
> - Masuk Chrome
> - isi URL dengan `chrome://settings/content/sound`
> - Pada bagian `Allowed to play sound`, Add host server seperti `localhost:3000`

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
