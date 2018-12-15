PKG="github.com/PacketFire/paste-click"
GOENV="ncatelli/golang:1.9.2-libmagic"

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
	rm -f immigrant; \
	docker rmi -f ${IMGNAME}:`cat "version.txt"`
