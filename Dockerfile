FROM golang:1.9.2

LABEL maintainer 'Nate Catelli <ncatelli@packetfire.org>'
LABEL description 'golang:1.9.2 with libmagic library installed'

RUN apt-get update && \
    apt-get install libmagic-dev -y && \
    apt-get clean
