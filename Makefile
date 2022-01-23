install:
	go install replace.go

test:
	go run replace.go

windows:
	go build -o build/replace.exe replace.go

linux:
	go build -o build/replace replace.go
