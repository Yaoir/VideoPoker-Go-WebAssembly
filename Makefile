# Make file for WebAssembly/Go version of video poker

main:
	GOOS=js GOARCH=wasm go build -o main.wasm main.go videopoker-web.go

vet:
	GOOS=js GOARCH=wasm go vet main.go videopoker-web.go

webserver:
	go build webserver.go

edit ed vi:
	vim main.go

count wc:
	wc *.go

test:
	./webserver

pub dep:
	@cp -a css img favicon.ico index.html main.wasm wasm_exec.js deploy

backup back bup:
	@cp -a css index.html deploy/upload* favicon.ico *.go *.js Makefile TODO .bak
