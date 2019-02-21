Paste.Click
===========

[![Build Status](https://travis-ci.org/PacketFire/paste-click.svg?branch=master)](https://travis-ci.org/PacketFire/paste-click)

<!-- TOC -->

- [General](#general)
- [Building](#building)
    - [Dependencies](#dependencies)
    - [Docker](#docker)
    - [Locally](#locally)
- [Testing](#testing)
- [Running](#running)
    - [Docker-Compose](#docker-compose)
- [Configuration](#configuration)

<!-- /TOC -->

## General
Paste-click is meant to serve as an online clipboard and provides easy interaction with files via curl.

## Building

### Dependencies
Currently, this tool only requires:

- libmagic/libmagic-dev
- make
- gcc

### Docker
The tool can be built and run entirely via docker using the following command.

```sh
$> make build-docker
```

### Locally
The tool can also be built and installed locally by running a pip install from the root of the project.

```sh
$> make build
```

## Testing
Testing is done using the stdlib testing package and full unit tests can be run with the following command.

```sh
$> make test
```

## Running
### Docker-Compose
A docker-compose environment has been included to ease both testing. This can be started with the minimal `docker-compose up` command.

This exposes the openresty service on port localhost:8080 and files can be uploaded by posting to this address with the Host header set to `paste.click`.

```sh
$> echo 'hello' | curl -sD - 'http://localhost:8080/' -H 'Host: paste.click' --data-binary @-
HTTP/1.1 200 OK
Server: openresty/1.13.6.2
Date: Wed, 20 Feb 2019 23:37:24 GMT
Content-Type: text/plain; charset=utf-8
Content-Length: 45
Connection: keep-alive
Content-Disposition: filename=""
Access-Control-Allow-Origin: *

http://localhost:8080/sZRqySSS0jR8YjW00mERhA
```

This object can then be queried with the following command.

```sh
$> curl -sD - 'http://localhost:8080/sZRqySSS0jR8YjW00mERhA' -H 'Host: paste.click'
Content-Type: text/plain
Content-Length: 6
Last-Modified: Thu, 03 Jan 2019 00:41:59 GMT
Connection: keep-alive
PC-Mime-type: text/plain
PC-Size: 6
PC-Object: sZRqySSS0jR8YjW00mERhA
PC-Uploaded: 2019-01-03 00:41:59.139780917 +0000 UTC m=+64.410846356
ETag: "5c2d5a57-6"
PC-Metafile: hit
Accept-Ranges: bytes

hello
```

## Configuration
Configuration parameters are provided via environment variables. Currently the following parameters can be set.

- ADDR: ':8080'

  The bind address for the service

- SITE_NAME: 'paste.click'
  
  A configuration setting representing the returned domain name in all upload calls.
- LOGGING: 'true'

  A boolean value determining whether to enable logging of requests to stdout.

- STORAGE_DRIVER: 'fs'

  The backend storage driver to use. Currently these options are mock, fs and gcs.
- STORE_FS_BASE_PATH: "/www/paste.click/"

  The base path of the to store files under. This should always point to the document root of openresty.
