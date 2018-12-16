PKG="github.com/PacketFire/paste-click"
GOENV="ncatelli/golang:1.9.2-libmagic"
IMGNAME="packetfire/paste-click"

build: | depend fmt test
	go build

depend:
	glide update ; glide install

build-docker: | depend fmt test
	docker build -t ${IMGNAME}:latest .

test: | depend
	go test -cover ./...

fmt: | depend
	go fmt ./...

clean:
	rm -f paste-click; \
	docker rmi -f ${IMGNAME}:latest
