build:
	go build -o deploy cmd/main.go

test:
	go test -timeout 30s github.com/holoGDM/awstool/internal/logic
	go test -timeout 30s github.com/holoGDM/awstool/internal/awselbv2
