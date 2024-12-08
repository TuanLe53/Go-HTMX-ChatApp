run:
	@templ generate
	@go run cmd/main.go

dev:
	@templ generate --watch --proxy="http://localhost:5050" --cmd="go run cmd/main.go"