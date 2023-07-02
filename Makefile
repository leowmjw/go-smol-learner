run:
	@go run *.go
ytsum:
	@cd cmd/ytsum && go run *.go
test:
	@gotest ./...

