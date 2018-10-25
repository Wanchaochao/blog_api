build:
	go build -o app main.go

stop:
	cat blog.pid | xargs kill

start:
	./app http -addr=:8083 > /var/log/blog_api.log 2>&1 &

restart:
	cat blog.pid | xargs kill
	./app http -addr=:8083 > /var/log/blog_api.log 2>&1 &
test:
	go test ./...