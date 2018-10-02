# junior

_A fast, tiny HTTP server for serving static content._

[![Travis CI: build status][travis-img]][travis]
[![Go Report Card][grc-img]][grc]
[![Docker Hub][docker-img]][docker]
[![MicroBadger][mb-img]][mb]

## Usage

### Docker

```bash
## Run server listening on port 3000, serving locally from "./local/www".
dk run stevenxie/junior -p "3000:80" -v "./local/www:/www/"
```

### Standalone

Download executable from
[releases](https://github.com/steven-xie/junior/releases), then run:

```bash
## Run HTTP server on port 4200, serving from "/var/www".
./junior -p 4200 --root "/var/www"

## For more details:
./junior --help
```

[travis]: https://travis-ci.com/steven-xie/junior
[travis-img]: https://travis-ci.com/steven-xie/junior.svg?branch=master
[grc]: https://goreportcard.com/report/github.com/steven-xie/junior
[grc-img]: https://goreportcard.com/badge/github.com/steven-xie/junior
[mb]: https://microbadger.com/images/stevenxie/juniora
[mb-img]: https://images.microbadger.com/badges/image/stevenxie/junior.svg
[docker]: https://hub.docker.com/r/steven-xie/junior
[docker-img]: https://img.shields.io/docker/pulls/stevenxie/junior.svg
