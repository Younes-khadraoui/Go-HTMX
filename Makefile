styles:
	npx @tailwindcss/cli -i ./web/static/css/input.css -o ./web/static/css/output.css --watch

run:
	PORT=$(PORT) go run cmd/goHtmx/main.go

r:
	wgo -file=.go -file=.html go run main.go

install:
	npm install

pre-commit:
	pip install pre-commit
	pre-commit install --install-hooks --overwrite
