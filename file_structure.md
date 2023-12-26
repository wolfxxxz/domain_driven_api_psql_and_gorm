service_user
|-- cmd/
|   |-- migration
|   |   |-- main.go
|   |
|   |-- main.go
|
|-- configs/
|   |-- .env
|
|-- internal/
|   |-- apperrors/
|   |   |-- apperrors.go
|   |
|   |-- config/
|   |   |-- config.go
|   |
|   |-- database
|   |   |-- postgres.go
|   |
|   |-- domain/
|   |   |-- mappers/
|   |   |   |-- mappers.go
|   |   |
|   |   |-- models/
|   |   |   |-- student.go
|   |   |   |-- lecture.go
|   |   |
|   |   |-- requests/
|   |   |   |-- requests.go
|   |   |
|   |   |-- responses/
|   |       |-- responses.go
|   |
|   |-- logger
|   |       |-- log.go
|   |
|   |-- repositories/
|   |   |-- repo_lecture.go
|   |   |-- repo_student.go
|   |
|   |-- server/
|   |   |-- handlers.go
|   |   |-- midlewares.go
|   |   |-- router.go
|   |   |-- server.go
|   |
|   |-- services/
|       |-- user_service.go
|       |-- lecture_service.go
|
|-- mock/
    |-- postgres_mock.go
