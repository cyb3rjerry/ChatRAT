all: slave master

slave:
	go build -o ./bin/slave ./cmd/slave

master:
	go build -o ./bin/master ./cmd/master

clean: 
	rm -rf ./bin
