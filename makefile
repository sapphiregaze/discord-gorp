discord-gorp:
	@go build ./cmd/discord-gorp/
run:
	@go run ./cmd/discord-gorp/
test:
	@go test ./...
clean:
	@rm ./discord-gorp