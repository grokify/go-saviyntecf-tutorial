# Saviynt External Connector Framework (ECF) in Go Tutorial

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

This is a tutorial and example app for the Saviynt ECF framework, [`go-saviyntecf`](https://github.com/grokify/go-saviyntecf).

To run this service, use:

```
% git clone https://github.com/grokify/go-saviyntecf-tutorial
% cd go-saviyntecf-tutorial
% go run main.go
% curl -XPOST http://localhost:8080/api/v1/accounts
% curl -XPOST http://localhost:8080/api/v1/users --verbose
```

## Further Reading

1. [Go, Lambda and the API Gateway — the OpenAPI v3.0 story](https://levelup.gitconnected.com/go-lambda-and-the-api-gateway-the-openapi-v3-0-story-a8afe5f841df)

 [used-by-svg]: https://sourcegraph.com/github.com/grokify/go-saviyntecf-tutorial/-/badge.svg
 [used-by-url]: https://sourcegraph.com/github.com/grokify/go-saviyntecf-tutorial?badge
 [build-status-svg]: https://github.com/grokify/go-saviyntecf-tutorial/workflows/test/badge.svg
 [build-status-url]: https://github.com/grokify/go-saviyntecf-tutorial/actions/workflows/test.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-saviyntecf-tutorial
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/go-saviyntecf-tutorial
 [codeclimate-status-svg]: https://codeclimate.com/github/grokify/go-saviyntecf-tutorial/badges/gpa.svg
 [codeclimate-status-url]: https://codeclimate.com/github/grokify/go-saviyntecf-tutorial
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-saviyntecf-tutorial
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/go-saviyntecf-tutorial
 [loc-svg]: https://tokei.rs/b1/github/grokify/go-saviyntecf-tutorial
 [repo-url]: https://github.com/grokify/go-saviyntecf-tutorial
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/go-saviyntecf-tutorial/blob/master/LICENSE
