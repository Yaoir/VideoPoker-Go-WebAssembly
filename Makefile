# Make file for WebAssembly/Go version of video poker

SRC=main.go videopoker-web.go

# build the main.wasm file

main:
	GOOS=js GOARCH=wasm go build -o main.wasm $(SRC)

# run 'go vet'

vet:
	GOOS=js GOARCH=wasm go vet $(SRC)

# build the web server for testing

webserver: webserver.go
	go build webserver.go

# line count of Go files

count wc:
	@wc $(SRC)

# run the web server to test the app

test:
	@./webserver

# copy files needed for deployment
# make sure the 'deploy' directory exists first!

pub dep:
	@cp -a css img favicon.ico index.html main.wasm wasm_exec.js deploy

# make a quick backup in the .bak directory
# make sure .bak exists first!
# (Note: some files are not included in the distribution,
# so you will need to modify this if you want to use it.)

backup back bup:
	@cp -a css index.html deploy/upload* favicon.ico *.go *.js Makefile TODO .bak
