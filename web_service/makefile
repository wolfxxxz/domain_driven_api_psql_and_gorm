run_server:
	go run cmd/main.go
mock_lecture:
	~/go/bin/mockgen -source=internal/repositories/lecture_repo.go -destination=./internal/mock/lecture_repo.go -package=mock
mock_users:
	~/go/bin/mockgen -source=internal/repositories/users_repo.go -destination=./internal/mock/users_repo.go -package=mock
build_app:
	go build -o Service_SCHOOL cmd/serviceschool/main.go
run_school:
	go run cmd/serviceschool/main.go
run_psql:
	docker-compose up --build