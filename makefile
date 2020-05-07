all:linux mac

.pre-build:
	mkdir -p build/darwin
	mkdir -p build/linux

linux:.pre-build
	GOOS=linux go build -ldflags "-w -s" -o build/linux/handshakeproxy

mac:.pre-build
	GOOS=darwin go build -ldflags "-w -s" -o build/darwin/handshakeproxy

clean:
	rm build/linux/handshakeproxy
	rm build/darwin/handshakeproxy

