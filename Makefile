BINARY_NAME=my-tool


help: #Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

run: #runs my-tool
	go run ./cmd/cli


build: #builds bin/my-tool
	go build -o dist/${BINARY_NAME} ./cmd/cli


clean: #removes files generated
	@rm ./dist/${BINARY_NAME}

air: #rebuild and runs the application on file change
	@air
