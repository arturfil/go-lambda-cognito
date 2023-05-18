upload:
	git add .
	git commit -m "last commit before updating"
	git push
	GOARCH=amd64
	GOOS=linux
	set GOOS=linux && set GOARCH=amd64 && set CGO_ENABLED=0 && go build -o main.go
	[ -e main.zip ] && rm main.zip
	# tar -a -cf main.zip main
