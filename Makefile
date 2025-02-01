build-assets:
	@npm run --prefix assets/vue build

build-server:
	@go build -o bin/app

run: build-assets build-server 
	@./bin/app
