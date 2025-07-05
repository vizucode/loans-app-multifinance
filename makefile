protoc:
	find proto -name "*.proto" -exec protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative {} \;

# Variabel untuk direktori migrasi dan opsi ekstensi file
EXT=sql
DB_PATH=migrations

# Database path
DATABASE_URI=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Default command
all: migrate

# Membuat file migrasi baru
migrate-create:
	migrate create -ext $(EXT) -dir $(DB_PATH) $(name)

# Menjalankan migrasi ke versi database terbaru
migrate-all:
	migrate -path $(DB_PATH) -database $(DATABASE_URI) up

migrate-up:
	migrate -path $(DB_PATH) -database $(DATABASE_URI) up $(version)

# Membatalkan migrasi terakhir
migrate-down:
	migrate -path $(DB_PATH) -database $(DATABASE_URI) down 1

# Membatalkan semua migrasi
migrate-reset:
	migrate -path $(DB_PATH) -database $(DATABASE_URI) down -all

migrate-clean:
	migrate -path $(DB_PATH) -database $(DATABASE_URI) force $(version)

# Menampilkan status migrasi
migrate-status:
	migrate -path $(DB_PATH) -database $(DATABASE_URI) version

.PHONY: all migrate-create migrate-up migrate-down migrate-reset migrate-status
