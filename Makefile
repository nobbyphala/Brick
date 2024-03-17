generate-mock:
	mockgen -source=./external/database/sql.go -destination=./mock/sql.go -package=mock
	mockgen -source=./usecase/repository/repository.go -destination=./mock/repository/repository.go -package=mock_repository
	mockgen -source=./usecase/api/api.go -destination=./mock/api/api.go -package=mock_api
	mockgen -source=./usecase/usecase.go -destination=./mock/usecase/usecase.go -package=mock_usecase

test:
	go test -v -race ./...