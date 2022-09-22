source ./.env
go test ./domains/login/... -coverprofile=cover.out && go tool cover -html=cover.out
go test ./domains/mentor/... -coverprofile=cover.out && go tool cover -html=cover.out
