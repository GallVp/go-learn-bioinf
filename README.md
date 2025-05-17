# _GO Learn Bioinf_

[![Build&Test](https://github.com/GallVp/go-learn-bioinf/actions/workflows/go.yml/badge.svg)](https://github.com/GallVp/go-learn-bioinf/actions/workflows/go.yml)

A personal project to learn _Go_ and bioinf in one go.

## Build and Test

> 🛠️ Build
```bash
go build
```

> 🚀 Test
```bash
go test ./...
```

> 🎯 Coverage
```
go test ./... -coverprofile=data/coverage.out
go tool cover -html=data/coverage.out
```