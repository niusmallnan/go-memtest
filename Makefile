.DEFAULT_GOAL := build

.PHONY: build clean image

build:
	mkdir -p bin
	CGO_ENABLED=0 go build -o bin/go-memtest

clean:
	git clean -dxf

image: build
	cp bin/go-memtest package/
	cd package && docker build -t niusmallnan/go-memtest .
	docker push niusmallnan/go-memtest
