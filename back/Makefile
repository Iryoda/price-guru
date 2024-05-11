main = cmd/main.go
sender = cmd/sender/sender.go
listener= cmd/listener/listener.go

watch:
	air 
test:
	go test -v ./...
run:
	go run $(main) 
run-sender:
	go run $(sender) 
run-listener:
	go run $(listener) 



build:
	go build $(main)
build-sender:
	go build $(sender)
build-listener:
	go build $(listener)
