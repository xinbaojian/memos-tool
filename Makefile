build: clean
	go build -v memos.go
clean:
	go clean
gox:
	gox
install:
	go install memos.go
win:
	gox -osarch="windows/amd64"
linux:
	gox -osarch="linux/amd64"