# vira-dapo

-- Create migration
migrate create -ext=sql -dir=internal/database/migrations -seq init

-- Migrate
migrate -path=internal/database/migrations -database "postgresql://postgres:postgres@localhost:5432/vira_dapo?sslmode=disable" -verbose up
