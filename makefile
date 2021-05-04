# Makefile

build: clean test
	go build -o .\build\ cmd\syncbrok\main.go 
	go build -o .\build\ cmd\subscriber\subs.go 
clean:
	del .\build\*.exe
test:
	@go test .\...
# .SILENT: