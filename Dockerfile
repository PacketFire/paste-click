ARG BASEIMG="alpine:3.8"
ARG BUILDIMG="golang:1.11.3-alpine3.8"
FROM $BUILDIMG as builder

ENV GOPATH=/go
ENV GIT_USER=PacketFire
ENV SCM_PROVIDER=github.com
ENV APP_NAME=paste-click

COPY . /go/src/${SCM_PROVIDER}/${GIT_USER}/${APP_NAME}/

RUN cd ${GOPATH} && \
    apk add --no-cache libmagic gcc libc-dev file-dev && \
    go build -o /${APP_NAME} ${SCM_PROVIDER}/${GIT_USER}/${APP_NAME}

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
USER ${SERVICE_USER}

ENTRYPOINT [ "/opt/paste-click/bin/paste-click" ]
CMD [ "-h" ]
