upload:
	set GOOS=linux 
	set GOARCH=amd64 
	set CGO_ENABLED=0
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
	rm -f main.zip
	zip main.zip main

