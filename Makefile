install:
	go install rename.go

test:
	go run rename.go

windows:
	go build -o build/rename.exe rename.go

linux:
	go build -o build/rename rename.go