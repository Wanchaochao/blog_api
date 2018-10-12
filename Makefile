build:
	go build -o app main.go

stop:
	cat blog.pid | xargs kill

start:
	./app http -addr=:8083 > /var/log/blog_api.log &

restart:
	cat blog.pid | xargs kill
	./app http -addr=:8083 > /var/log/blog_api.log &
test:
	go test ./...