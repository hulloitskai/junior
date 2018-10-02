# junior

_A fast, tiny HTTP server for serving static content._

[![Go Report Card][grc-img]][grc]
[![Docker Hub][docker-img]][docker]
[![MicroBadger][mb-img]][mb]

## Usage

### Docker _(recommended)_

```bash
## Run server listening on port 3000, serving locally from "./local/www".
dk run stevenxie/junior -p "3000:80" -v "./local/www:/www/"
```

### Standalone

Download executable from
[releases](https://github.com/steven-xie/junior/releases), then run:

```bash
./junior # run HTTP server, configured via environment variables
```

[grc]: https://goreportcard.com/report/github.com/steven-xie/junior
[grc-img]: https://goreportcard.com/badge/github.com/steven-xie/junior
[mb]: https://microbadger.com/images/stevenxie/juniora
[mb-img]: https://images.microbadger.com/badges/image/stevenxie/junior.svg
[docker]: https://hub.docker.com/r/steven-xie/junior
[docker-img]: https://img.shields.io/docker/pulls/stevenxie/junior.svg
