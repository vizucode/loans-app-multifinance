# --- Stage 1: Builder ---
# Menggunakan image Go yang lebih besar untuk proses kompilasi
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod dan go.sum terlebih dahulu untuk memanfaatkan Docker cache
# Jika dependensi tidak berubah, layer ini tidak akan dibangun ulang
COPY go.mod go.sum ./
RUN go mod tidy

# Copy semua file project lainnya
COPY . .

# Build aplikasi Go Anda
# -o /app/multifinance_test: output biner ke /app/multifinance_test
# -ldflags "-s -w": mengurangi ukuran biner dengan menghilangkan simbol debug
# main.go: file Go utama yang akan di-build
RUN CGO_ENABLED=0 go build -o /app/multifinance_test -ldflags "-s -w" main.go

# --- Stage 2: Runner (Produksi) ---
# Menggunakan image yang sangat minimal untuk menjalankan aplikasi
# scratch: image Docker paling kecil yang hanya berisi biner Anda
FROM scratch AS runner

# Opsional: Jika aplikasi Go Anda memerlukan sertifikat SSL/TLS atau file lain dari sistem,
# Anda mungkin perlu image yang sedikit lebih besar, misalnya alpine.
# FROM alpine AS runner
# RUN apk add --no-cache ca-certificates # Jika menggunakan alpine, instal sertifikat root

# Salin biner yang sudah di-build dari 'builder' stage
COPY --from=builder /app/multifinance_test /multifinance_test

# Expose port jika aplikasi Go Anda adalah server HTTP
# EXPOSE 8080 # Ganti dengan port yang digunakan aplikasi Anda

# Tentukan perintah untuk menjalankan aplikasi
ENTRYPOINT ["/multifinance_test"]
# Atau jika Anda menggunakan alpine:
# CMD ["/multifinance_test"]