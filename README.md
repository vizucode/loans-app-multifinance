# Privy Backend CMS

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
2. Installed [Postgres](https://www.postgresql.org/download/) on your machine
3. Add your env to .env file or set it manually in your machine
4. Install [docker](https://docs.docker.com/get-started/introduction/) and [docker compose](https://docs.docker.com/compose/install/)
5. Install [Makefile](https://www.gnu.org/software/make/manual/make.html) to your machine
6. Install [golang-migrate](https://github.com/golang-migrate/migrate) to your machine

## Instalation
1. Clone the repository
2. Run `docker compose up -d` to start the services

## How to use migrations
The migration will run using [golang-migrate](https://github.com/golang-migrate/migrate)
1. Run `make migrate-up version={target_migrate_version}` to migrate up to the target version
2. Run `make migrate-down` to migrate down which will rollback to the last migration
3. Run `make migrate-clean version={target_migrate_version}` to reset the migration to the target version which is the dirty migration.
4. Run `make migrate-status` to get the current version of the migration
5. Run `make migrate-reset` to rollback all migration to the initial version
6. Run `make migrate-all` to migrate all migration to the latest version

## ERD Documentations
Click this link to see the [ERD documentations](https://dbdocs.io/zen%20dev/PrivyCMS)

## API Specification
Click this link to see the [API documentation](https://apidog.com/apidoc/docs-site/753332)
