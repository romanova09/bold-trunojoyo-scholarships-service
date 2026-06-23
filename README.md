# Bold Trunojoyo Scholarships Service

Backend service untuk program beasiswa **Bold Trunojoyo Scholarships** — menangani blast/notifikasi email ke donatur dan pendaftar beasiswa (registrasi donatur, laporan donasi, pengumuman hasil seleksi), serta reminder donasi bulanan via scheduled job.

## Tech Stack

- Golang 1.20
- Gin (HTTP framework)
- go-pg (PostgreSQL ORM)
- robfig/cron (scheduler)
- net/smtp (email delivery via Gmail SMTP)
- Docker

## Features

- **Welcome Email** — kirim email sambutan ke donatur baru saat registrasi, berisi info rekening donasi.
- **Donation Report Blast** — broadcast laporan penyaluran donasi (penerima, prodi, angkatan, total UKT) ke seluruh donatur.
- **Applicant Announcement** — kirim email pengumuman hasil seleksi beasiswa (lolos/tidak lolos) ke pendaftar.
- **Donation Reminder (Cron)** — job harian yang mengirim email pengingat donasi ke donatur sesuai tanggal reminder masing-masing (`date_reminder`), terjadwal lewat `robfig/cron`.

## API

| Method | Endpoint                          | Description                              |
|--------|------------------------------------|-------------------------------------------|
| POST   | `/registrations/welcome-email`     | Kirim email sambutan ke donatur baru      |
| POST   | `/registrations/donation-report`   | Broadcast laporan donasi ke semua donatur |
| POST   | `/registrations/applicant-failed`  | Kirim email pengumuman hasil seleksi      |

Semua endpoint mengembalikan response standar:
```json
{ "success": true, "message": "success", "data": {} }
```

## Folder Structure

```
.
├── app/                  # API bootstrap & dependency wiring
├── config/               # env-based config loader
├── db/                   # Postgres (go-pg) connection
├── infra/                # SMTP email implementation (IEmail abstraction)
├── scheduler/            # cron job: donation reminder
├── src/donor/            # donor domain: handler, service, repository, DTO
├── model/                # Donor entity
├── helpers/              # response wrapper, validator, HTML template parser
├── assets/template/email/ # HTML email templates
└── main.go
```

## Running locally

```bash
cp .env.example .env   # isi dengan kredensial SMTP & DB kamu sendiri
go mod download
go run main.go
```

Atau via Docker:

```bash
cp .env.example .env
docker compose up --build
```

API berjalan di port sesuai `APP_PORT` di `.env` (default `9000`).

## Environment Variables

Lihat [.env.example](.env.example) untuk daftar lengkap variabel yang dibutuhkan (`APP_*`, `DB_*`, `SMTP_*`, `BANK_*`, `DONATION_LINK`). **Jangan commit file `.env`** — gunakan App Password (bukan password akun Gmail biasa) untuk `SMTP_PASSWORD`.

## Example Request

```bash
curl -X POST http://localhost:9000/registrations/welcome-email \
  -H "Content-Type: application/json" \
  -d '{
    "to": [
      { "email": "donor@example.com", "nama": "Budi Santoso" }
    ]
  }'
```
