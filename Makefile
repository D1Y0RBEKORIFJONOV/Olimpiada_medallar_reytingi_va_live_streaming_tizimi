run:
	go run cmd/main.go

table:
	migrate create -dir migrations -ext sql db
table-up:
	migrate -path migrations -database "postgres://postgres:2005@localhost:5432/hackathon?sslmode=disable" up
table-down:
	migrate -path migrations -database "postgres://postgres:2005@localhost:5432/hackathon?sslmode=disable" down
table-force:
	migrate -path migrations -database "postgres://postgres:2005@localhost:5432/hackathon?sslmode=disable" force