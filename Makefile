run:
	cd src/ && go run .
web: 
	cd src/ && env GOOS=js GOARCH=wasm go build -o PopTheLock.wasm
	mv src/PopTheLock.wasm PopTheLock.wasm