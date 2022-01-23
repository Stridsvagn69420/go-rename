install:
	go install go-rename.go

test:
	go run go-rename.go

windows:
	go build -o build/go-rename.exe go-rename.go

linux:
	go build -o build/go-rename go-rename.go
