upload:
	set GOOS=linux && set GOARCH=amd64 && set CGO_ENABLED=0 && go build -o main main.go
	rm -f main.zip
	tar -czvf main.tar.gz main
