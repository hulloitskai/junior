##################################################
## BUILD STAGE
##################################################

FROM golang:alpine AS build

## Copy source files.
WORKDIR /build
COPY . .

## Configure build environment.
ENV GO111MODULE=on

## Install external + app dependencies.
RUN apk add make git gcc musl-dev upx
RUN go version && make get

## Create production binary.
RUN make build

## Compress binary
RUN upx --brute junior


##################################################
## PRODUCTION STAGE
##################################################

FROM scratch as production

LABEL maintainer="Steven Xie <hello@stevenxie.me>"
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.name = "junior"
LABEL org.label-schema.name = "A fast, tiny HTTP server for serving static \
  content."
LABEL org.label-schema.vcs-url="https://github.com/steven-xie/junior"
LABEL org.label-schema.vendor="Steven Xie <hello@stevenxie.me>"

## Copy production artifacts to /etc/junior.
COPY --from=build /build/junior .

## Copy default www/ files.
COPY ./www/ /www/

## Configure default environment.
ENV PORT=80 ROOT_DIR=/www/ TRAILING_SLASHES=false

## Expose default port, set entrypoint.
EXPOSE $PORT
ENTRYPOINT ["./junior"]
