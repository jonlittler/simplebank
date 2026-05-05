# Backend

## DB
https://dbdiagram.io \
https://dbdiagram.io/d/Simple-Bank-69ec074fddb9320fdc481758

Login with jonathan.littler@gmail.com \
Simple Bank

```bash
make --version      # GNU Make 4.3
go version          # go version go1.25.1 linux/amd64
docker version      # Version: 28.1.1
migrate -version    # dev
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
sqlc version        # v1.31.1
```

### Go
```bash
go mod init github.com/jonlittler/ts/simplebank
```

### Viper
https://github.com/spf13/viper
```bash
go get github.com/spf13/viper
SERVER_ADDR=0.0.0.0:8082 make server
```

### Migrate
https://github.com/golang-migrate/migrate/tree/master

```bash
curl -fsSL https://packagecloud.io/golang-migrate/migrate/gpgkey \
| sudo gpg --dearmor -o /etc/apt/keyrings/migrate.gpg

echo "deb [signed-by=/etc/apt/keyrings/migrate.gpg] https://packagecloud.io/golang-migrate/migrate/ubuntu/ jammy main" \
| sudo tee /etc/apt/sources.list.d/migrate.list > /dev/null

sudo apt-get update
sudo apt-get install -y migrate

migrate -version    # 4.19.1
migrate create -ext sql -dir db/migration -seq init_schema
migrate create -ext sql -dir db/migration -seq add_users
migrate -path db/migration -database "postgres://root:password@localhost:5432/simplebank?sslmode=disable" -verbose up
```

### SQLC
https://docs.sqlc.dev/en/latest/ \
https://sqlc.dev/

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
sqlc init
sqlc version    # v1.31.1
```

### Mock
https://github.com/golang/mock
```bash
go get github.com/golang/mock/mockgen@v1.6.0        # needed for reflect mode
go install github.com/golang/mock/mockgen@v1.6.0    # ~/go/bin/mockgen
mockgen -version                                    # v1.6.0
mockgen -help
mockgen -package mockdb -destination ./db/mock/store.go github.com/jonlittler/ts/simplebank/db/sqlc Store 
```
