# RESEARCH

## Building Agents

- AmpCode - https://ampcode.com/how-to-build-an-agent

## MCP in Golang

- Write own first before using libraries ..
- 
- Try it out with aider - https://mcpgolang.com/quickstart
- MCP Go - https://github.com/metoro-io/mcp-golang
- MCP Go Server - https://github.com/metoro-io/metoro-mcp-server
- Try out examples - https://github.com/metoro-io/mcp-golang/tree/main/examples
- Kagi MCP - https://github.com/mgomes/kagimcp-go

## LLM Routers

- OpenRouter - 
- Cloudflare API Inference - 
- Ollama/local? - Qwen?

## Search,Summarizer

- Kagi - https://help.kagi.com/kagi/api/search.html
- Kagi - https://github.com/httpjamesm/kagigo
- Kagi - 

## Tooling

- TaskMaster - Use it to breakdown Agent features
- aider - Use it for generic chat / comapre against others?
- jj - Jujutsu: see if can be used for managing vibe code changes; more experiments

### jj usage

Using with Github: https://jj-vcs.github.io/jj/latest/github/

This does not work to put into direnv .. useful can pull out to make?
```
# Aliases (sorted alphabetically)
alias jjc='jj commit'
alias jjcmsg='jj commit --message'
alias jjd='jj diff'
alias jjdmsg='jj desc --message'
alias jjds='jj desc'
alias jje='jj edit'
alias jjgcl='jj git clone'
alias jjgf='jj git fetch'
alias jjgp='jj git push'
alias jjl='jj log'
alias jjla='jj log -r "all()"'
alias jjn='jj new'
alias jjrb='jj rebase'
alias jjrs='jj restore'
alias jjrt='cd "$(jj root || echo .)"'
alias jjsp='jj split'
alias jjsq='jj squash'
```

## Golang Practice

This time will try this convention --> https://medium.com/inside-picpay/organizing-projects-and-defining-names-in-go-7f0eab45375d

```
├── README.md
├── go.mod
├── go.sum
├── cmd
│   ├── api
│   │   └── main.go
│   └── appctl
│       └── main.go
├── config
│   ├── config.go
├── book
│   ├── book.go
│   ├── service.go
│   ├── service_test.go
│   ├── mock
│   │   ├── service.go
│   │   ├── repository.go
│   ├── postgres
│   │   ├── book_storage.go
│   │   ├── book_storage_test.go
├── internal
│   ├── http
│       ├── gin
│        ├── handler.go
|        ├── book.go
│        ├── book_test.go
│   ├── event
│   │   ├── event.go
│   │   └── event_test.go
│   │   │── kafka
│   │       ├── event.go
│   │       ├── event_test.go
```

## K8s Observability

- Metoro - https://metoro.io/
- Test MCo with Claude Desktop + Metero MCP - https://github.com/metoro-io/metoro-mcp-server

