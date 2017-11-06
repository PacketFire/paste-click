PKG="gitlab.packetfire.org/Tiksi/paste-click"
GOENV="ncatelli/golang:1.9.2-libmagic"

build: | test
	docker run -it --rm -u root -v `pwd`:/go/src/$(PKG) $(GOENV) go build $(PKG)

fmt:
	docker run -it --rm -u root -v `pwd`:/go/src/$(PKG) $(GOENV) go fmt $(PKG)

test: fmt
	docker run -it --rm -u root -v `pwd`:/go/src/$(PKG) $(GOENV) go test $(PKG)
