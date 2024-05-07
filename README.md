# TanX Connector Go

<a href="https://pkg.go.dev/github.com/tanx-libs/tanx-connector-go"><img src="https://pkg.go.dev/badge/github.com/tanx-libs/tanx-connector-go.svg" alt="Go Reference"></a>

The official Go connector for <a href="https://docs.tanx.fi/tech/api-documentation">tanX's API</a> 🚀


## Installation

```bash
go get github.com/tanx-libs/tanx-connector-go
```

## Features

- Complete endpoints including REST and WebSockets
- Methods return parsed JSON.
- High level abstraction for ease of use.
- Easy authentication
- Automatically sets JWT token internally
- Auto refresh tokens when access token expires

Tanx-connector-go includes utility/connector functions which can be used to interact with the Tanx API. It uses requests internally to handle all requests.