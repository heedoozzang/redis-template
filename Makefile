up:	
	rm -f redis-template-linux
	GOOS=linux GOARCH=amd64 go build -o redis-template-linux
	docker-compose up --build

clean:
	docker-compose down
	rm -rf redis-template
	rm -rf redis-template-linux

.PHONY: up clean