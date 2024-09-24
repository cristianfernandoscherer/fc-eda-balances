# Run migrations
migrate -path=internal/database/migrations -database "mysql://balance:balance@tcp(localhost:3308)/balance   " -verbose up