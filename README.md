# grawler

Golang Crawler

## Environment Setting

1. Go Lang Install(v1.17)
2. [ProtoBuffer Install](https://github.com/protocolbuffers/protobuf)
3. Dependency Setting

### Go Lang Install(v1.17)

```cmd
// windows
go env -w GO111MODULE=auto
set GOPATH=[PATH]\grawler
set GOBIN=[PATH]\grawler\bin
set Path=%PATH%;%GOPATH%\bin

setx GOPATH [PATH]\grawler
setx GOBIN [PATH]\grawler\bin
setx Path %PATH%;[PATH]\grawler\bin
```

```shell
// mac os
go env -w GO111MODULE=auto
go env -w GOPATH=[PATH]/grawler
go env -w GOBIN=[PATH]/grawler/bin
export PATH=$PATH:$(go env GOPATH)/bin
```

## Dependencies

```shell
go get github.com/antchfx/htmlquery

go get -u github.com/antchfx/htmlquery
```

## Build

```shell
# golang
go build -o ./dist/grawler grawler.go
```

## gopls setting

```json
{
    ...,
    "go.useLanguageServer": true,
  "[go]": {
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    },
    "editor.snippetSuggestions": "none"
  },
  "[go.mod]": {
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "gopls": {
    "experimentalWorkspaceModule": true,
    "usePlaceholders": true,
    "staticcheck": false
  }
}
```
