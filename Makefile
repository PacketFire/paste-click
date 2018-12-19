PKG="github.com/PacketFire/paste-click"
GOENV="ncatelli/golang:1.9.2-libmagic"
IMGNAME="packetfire/paste-click"

build: | fmt test
	go build

build-docker: | fmt test
	docker build -t ${IMGNAME}:latest .

test:
	go test -cover ./...

fmt:
	go fmt ./...

clean:
	rm -f paste-click; \
	docker rmi -f ${IMGNAME}:latest
