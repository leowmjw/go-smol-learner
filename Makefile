run:
	@go run *.go

watch:
	@pkgx watch --color jj --ignore-working-copy log --color=always

detailed:
	@jj log -T builtin_log_detailed

evolog:
	@jj evolog -p

squash:
	@jj squash

ytsum:
	@cd cmd/ytsum && go run *.go
test:
	@gotest ./...

