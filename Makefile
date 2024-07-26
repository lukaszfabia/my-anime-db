tailwind:
	@echo "Running tailwindcss..."
	@npx tailwindcss -i styles/input.css -o styles/output.css --watch

templ:
	@echo "Running templ..."
	@templ generate -watch -proxy=http://localhost:8080

dev:
	@echo "Running dev..."
	@go run main.go
