ARG BASEIMG="alpine:3.8"
ARG BUILDIMG="golang:1.11.4-alpine3.8"
FROM $BUILDIMG as builder

ENV APP_NAME=paste-click

COPY . /go/

RUN cd /go/ && \
    unset GOPATH && \
    apk add --no-cache git gcc libmagic libc-dev file-dev && \
    go build -o /${APP_NAME} 

FROM $BASEIMG
LABEL maintainer="Nate Catelli <ncatelli@packetfire.org>" \
      description="Container for paste-click"

ENV SERVICE_USER "paste-click"
ENV APP_NAME=paste-click

RUN addgroup ${SERVICE_USER} && \
    adduser -D -G ${SERVICE_USER} ${SERVICE_USER} && \
    apk add --no-cache libmagic

COPY --from=builder /${APP_NAME} /opt/${APP_NAME}/bin/

RUN chown -R ${SERVICE_USER}:${SERVICE_USER} /opt/${APP_NAME} && \
    chmod +x /opt/${APP_NAME}/bin/${APP_NAME}

WORKDIR "/opt/$APP_NAME/"
#USER ${SERVICE_USER}

ENTRYPOINT [ "/opt/paste-click/bin/paste-click" ]
CMD [ "-h" ]
