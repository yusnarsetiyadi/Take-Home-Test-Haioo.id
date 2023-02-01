test:
	go test ./features/shopping... -coverprofile=cover.out && go tool cover -html=cover.out