start:
	@npm install
	@npx tailwindcss -i ./assets/css/style.css -o ./assets/css/out.css --minify
	@templ generate
	@go run cmd/main.go
