tailwind:
	@npx tailwindcss -i ./src/css/input.css -o ./static/css/styles.css --watch

templ:
	@templ generate -watch -proxy=http://localhost:3000

webpack:
	@npm run watch

build:
	go build -o ./tmp/main ./cmd/main.go

test:
	go test -v ./... -count=1 

clean:
	echo "Cleaning up..."
	rm -f ./static/css/styles.css
	rm -f ./static/js/bundle.js
	rm -rf ./node_modules

