release: build
prepare: echo pull_dependencies

echo:
	echo "prepare"
pull_dependencies:
	go mod tidy
build:
	GOOS=linux GOARCH=amd64 go build -o connector main.go