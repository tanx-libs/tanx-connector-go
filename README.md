# TanX Connector Go
- [X] Go ducomentation automatically generating (following go idioms)

- [ ] Public unprotected endpoint
  - [X] JSON
    - [X] health
    - [ ] ticker
    - [ ] all ticker
    - [ ] candle
    - [ ] orderbook
    - [ ] recent trade
  - [ ] Websocket (exploring better dx for clints using the package)
    - [X] trade streams
    - [ ] orderbook
    - [ ] kline

(next PR target)
- [ ] Public protected endpoints
- [ ] understand voh signing wali cheez before 3:30 kese bhi karke (I am a high value software engineer)


- [X] all types are provided in a different package
- [X] add context support for adding request timeouts

- [ ] Testing
  - [X] context testing 
  - [X] using roundtripper for a fast http client testing with ability to create mocks
  - [X] added github action go testing `go test -v ./...`