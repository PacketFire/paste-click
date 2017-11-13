PKG="gitlab.packetfire.org/Tiksi/paste-click"
GOENV="ncatelli/golang:1.9.2-libmagic"

build: | test
	docker run --rm -u root -v $(PWD):/go/src/$(PKG) $(GOENV) go build $(PKG)

fmt:
	docker run --rm -u root -v $(PWD):/go/src/$(PKG) $(GOENV) go fmt $(PKG)

test: | fmt
	docker run --rm -u root -v $(PWD):/go/src/$(PKG) $(GOENV) go test $(PKG)
