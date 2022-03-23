compile:
	go build -ldflags="-s -w" go-rename.go

install:
	go install -ldflags="-s -w" go-rename.go
