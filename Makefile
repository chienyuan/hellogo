
hello:
		CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello  .

build:
	docker build -t hellogo .
