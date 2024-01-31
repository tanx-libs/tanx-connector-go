# TanX Connector Go
- [X] Go ducomentation automatically generating (following go idioms)

- [ ] Public unprotected endpoint
  - [X] JSON
    - [X] health
    - [X] ticker
    - [X] all ticker
    - [X] candle
    - [ ] orderbook (doubt: server not accepting GET requests)
    - [X] recent trade
    - [X] add context support for adding request timeouts
    - [ ] complete the tests of these
      
  - [ ] Websocket
    - [X] trade streams
    - [ ] orderbook
    - [ ] kline

- [ ] Public protected endpoints


- [ ] Testing
  - [X] context testing 
  - [X] using roundtripper for a fast http client testing with ability to create mocks
  - [X] added github action go testing `go test -v ./...`
  - [ ] websocket testing
  - [ ] Apply the same test for every endpoint 