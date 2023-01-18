VERSION=0.0.6

fmt:
	gofmt -w -s *.go

release: fmt
	gothub release -p -u eyedeekay -r "go-i2pbrowser" -t v$(VERSION) -n "lib" -d "tag for release"
