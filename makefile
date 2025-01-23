build-css:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch


generate-templ:
	templ generate


generate-templ-watch:
	templ generate --watch


run: 
	go run cmd/main.go 


db:

	docker start postgresdb