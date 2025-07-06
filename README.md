# README!

# Unit Test
Saya sengaja hanya menulis unit test pada method CreateLoan karena menurut saya, di situlah titik paling penting dalam proses transaksi pinjaman. Method ini menangani proses inti—mulai dari validasi hingga penyimpanan data ke database—dan kalau sampai ada yang salah di sini, dampaknya bisa sangat besar. Saya nggak fokus untuk menulis test di semua tempat, karena saya percaya testing itu harus strategis, bukan asal banyak.
```bash
/apps/service/loan/create_test.go
```

## Environtment
| Name | Description | Example |
| --- | --- | --- |
| APP_HOST | Host of the application | 0.0.0.0 |
| APP_PORT | Port of the application | 8080 |
| APP_ENV | Environment of the application | production/development |
| SERVICE_NAME | Name of the service | privy-cms |
| DB_HOST | Host of the database | db |
| DB_PORT | Port of the database | 5432 |
| DB_USER | User of the database | postgres |
| DB_PASSWORD | Password of the database | password |
| DB_NAME | Name of the database | db_privy |
| AWS_SECRET_KEY | Secret key of the Storage | 1234567890 |
| AWS_ACCESS_KEY | Access key of the Storage | 1234567890 |
| AWS_REGION | Region of the Storage | ap-southeast-1 |
| AWS_BUCKET_NAME | Name of the bucket | privy-cms |
| AWS_S3_FORCE_PATH_STYLE | Force path style of the S3 | true/false |

## Prerequisites
1. Please read the docs carefully
2. Installed [Mysql](https://www.mysql.com/downloads/) on your machine
3. Add your env to .env file or set it manually in your machine
4. Install [docker](https://docs.docker.com/get-started/introduction/) and [docker compose](https://docs.docker.com/compose/install/)
5. Install [Makefile](https://www.gnu.org/software/make/manual/make.html) to your machine

## Instalation
1. Clone the repository
2. Run `docker compose up -d` to start the services

## OWASP Standard
Aplikasi yang saya bangun ini udah saya siapkan dengan pendekatan keamanan yang merujuk ke standar OWASP. Buat autentikasi, saya pakai JWT lengkap dengan refresh token, jadi user tetap bisa nyaman login tanpa harus kompromi keamanan. Di level database, saya udah pakai ORM, biar query-nya lebih aman dan terhindar dari SQL Injection, salah satu serangan yang sering masuk dalam top 10 OWASP. Untuk monitoring dan pelacakan error, saat ini saya masih pakai basic logging dari Go, tapi strukturnya udah saya buat rapi dan siap banget kalau nanti mau diintegrasikan ke sistem log monitoring seperti ELK, Grafana, atau Sentry. Jadi walaupun masih tahap awal, pondasi aplikasinya udah mengarah ke best practice standar keamanan OWASP.

## ERD Documentations
![ERD Picture](/assets/erd.png)

## Arsitektur Sistem
![Arsistektur Sistem](/assets/arsitektur.png)
Arsitektur ini secara umum sudah cukup aman untuk mencapai tingkat ketersediaan (availability) 99,9%, atau sekitar 8 jam downtime per tahun, selama dijalankan dengan praktik operasional yang baik. Dengan pembagian beban melalui Nginx sebagai load balancer di VPS-1 dan pemisahan layanan data di VPS-2 (Supabase dan MySQL), sistem ini sudah memiliki fondasi yang solid untuk menjaga kestabilan aplikasi. Selama setiap komponen (APP-1, APP-2, APP-3, Nginx, dan database) dijalankan dengan mekanisme auto-restart, memiliki sistem monitoring kesehatan, serta backup database yang rutin dan teruji, maka arsitektur ini sudah mencukupi untuk menjamin ketersediaan layanan tingkat menengah secara andal. Asalkan koneksi antar VPS stabil dan recovery plan disiapkan dengan baik, arsitektur ini sudah layak untuk digunakan dalam produk digital yang menargetkan uptime 99,9%.


## API Specification
Click this link to see the ![Swagger](/assets/study-case.openapi.json)
