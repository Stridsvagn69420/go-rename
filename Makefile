compile:
        go build -ldflags="-s -w" go-rename.go

install:
	go build -ldflags="-s -w" go-rename.go
