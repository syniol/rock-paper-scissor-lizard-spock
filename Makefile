build:
	env GOOS=darwin GOARCH=amd64 go build -o ./dist/rpsls-cli-mac-os-amd64 ./cmd/cli.go
	env GOOS=darwin GOARCH=arm64 go build -o ./dist/rpsls-cli-mac-os-arm64 ./cmd/cli.go
	env GOOS=linux GOARCH=amd64 go build -o ./dist/rpsls-cli-linux-amd64 ./cmd/cli.go
	env GOOS=linux GOARCH=arm64 go build -o ./dist/rpsls-cli-linux-arm64 ./cmd/cli.go
	env GOOS=linux GOARCH=arm go build -o ./dist/rpsls-cli-linux-arm ./cmd/cli.go
	env GOOS=linux GOARCH=386 go build -o ./dist/rpsls-cli-linux-386 ./cmd/cli.go
	env GOOS=solaris GOARCH=amd64 go build -o ./dist/rpsls-cli-solaris-amd64 ./cmd/cli.go
	env GOOS=windows GOARCH=386 go build -o ./dist/rpsls-cli-windows-386 ./cmd/cli.go
	env GOOS=windows GOARCH=amd64 go build -o ./dist/rpsls-cli-windows-amd64 ./cmd/cli.go


run:
	go run ./cmd/cli.go
