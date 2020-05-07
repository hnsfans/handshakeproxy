# handshakeproxy

## Setup A http proxy to visit handshake domain.

## On Mac Better Put handshake.sh to your bin dir . Then You can quickly on or off handshake proxy

```shell
./handshake.sh on
./handshake.sh off
```

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
