# handshakeproxy

## Setup A http proxy to visit handshake domain.

## Build handshakeproxy

```shell
go build
```

## Start proxy

```shell
./handshakeproxy proxy
```

## Simple Test

```shell
export http_proxy=localhost:8080 && curl -v -s welcome.nb
# Or ./handshakeproxy debug
```
