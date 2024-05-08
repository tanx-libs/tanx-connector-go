# TanX Connector Go

![reference](https://pkg.go.dev/badge/github.com/tanx-libs/tanx-connector-go.svg)
![testing](https://github.com/tanx-libs/tanx-connector-go/actions/workflows/tests.yml/badge.svg)

The official Go connector for <a href="https://docs.tanx.fi/tech/api-documentation">tanX's API</a> 🚀


## Features

- Complete endpoints including REST and WebSockets
- Methods return parsed JSON.
- High level abstraction for ease of use.
- Easy authentication
- Automatically sets JWT token internally
- Auto refresh tokens when access token expires

Tanx-connector-go includes utility/connector functions which can be used to interact with the Tanx API. It uses requests internally to handle all requests.

## Installation

```bash
go get github.com/tanx-libs/tanx-connector-go
```

## Documentation
- [API Reference](https://pkg.go.dev/github.com/tanx-libs/tanx-connector-go)
- [Examples](https://github.com/tanx-libs/tanx-connector-go/blob/main/examples.md)