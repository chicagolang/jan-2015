usage:
	@echo "make assets       |   Generate production asset file"
	@echo "make dev-assets   |   Generate development asset file"
	@echo "make dev          |   Build with development assets"
	@echo "make build        |   Build with production assets"
	@echo "make release      |   Build release binaries"

assets: static/
	go-bindata $(BINDATA_OPTS) -ignore=[.]gitignore -ignore=[.]gitkeep $<...

dev-assets:
	@$(MAKE) --no-print-directory assets BINDATA_OPTS="-debug"

dev: dev-assets
	go build

build: assets
	go build

release: assets
	gox -osarch="darwin/amd64 darwin/386 linux/amd64 linux/386 windows/amd64 windows/386" -output="./bin/example_{{.OS}}_{{.Arch}}"

setup:
	go get github.com/mitchellh/gox
	go get github.com/jteeuwen/go-bindata/...

clean:
	rm -f ./example_app
	rm -f ./bin/*
	rm -f bindata.go
	make assets