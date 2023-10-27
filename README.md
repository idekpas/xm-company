# xm-company

### External dependencies used in project (not every, just not common ones):

**mux:** For router
([https://github.com/gorilla/mux](https://github.com/gorilla/mux))

**logrus:** For logs
([https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus))

**sqlx:** as extension to database/sql
([https://github.com/jmoiron/sqlx](https://github.com/jmoiron/sqlx))

**pq:**  for postgres connection ([https://github.com/lib/pq](https://github.com/lib/pq))

**govalidator:**  for data validation ([https://github.com/asaskevich/govalidator](https://github.com/asaskevich/govalidator))

**godotenv:** and **envconfig:** for env files ([github.com/joho/godotenv](github.com/joho/godotenv) ([github.com/kelseyhightower/envconfig](github.com/kelseyhightower/envconfig))


### To run:
- Change db credentials in .env file
- run `go mod tidy` / `go mod vendor`
- run `migrate -database ${POSTGRESQL_URL} -path db/migrations up` 
- ${POSTGRES_URL} - link with your credentials, for example and with those in .env file: `export POSTGRESQL_URL='postgres://postgres:xm1qazXSW@@localhost:5432/xm?sslmode=disable'`
- `go run cmd/xmapp/main.go`

### Things to do that I unfortunately didn't have time for (one more day and I would finish):
- JWT auth
- dockerfiles with some makefile
- kafka
- unit tests and integration ones (sorry!)

but the api, db connection etc is done