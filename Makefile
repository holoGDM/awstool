build:
	go build -o deploy cmd/main.go

test:
	go test -timeout 30s -coverprofile=/tmp/vscode-goJ9F13Q/go-code-cover github.com/holoGDM/awstool/internal/logic
	go test -timeout 30s -coverprofile=/tmp/vscode-goJ9F13Q/go-code-cover github.com/holoGDM/awstool/internal/awselbv2