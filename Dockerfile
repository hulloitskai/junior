##################################################
## BUILD STAGE
##################################################

FROM golang:alpine AS build

## Copy source files.
WORKDIR /build
COPY . .

## Configure build environment.
ENV GO111MODULE=on CGO_ENABLED=0

## Install build + external dependencies:
RUN apk add --no-cache make git upx
RUN go version && make get

## Create production binary.
RUN make build

## Compress binary.
RUN upx junior --brute


##################################################
## PRODUCTION STAGE
##################################################

FROM busybox:1.29 as production

LABEL maintainer="Steven Xie <hello@stevenxie.me>"
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.name = "junior"
LABEL org.label-schema.description="A fast, tiny HTTP server for serving \
  static content."
LABEL org.label-schema.vcs-url="https://github.com/steven-xie/junior"
LABEL org.label-schema.vendor="Steven Xie <hello@stevenxie.me>"

## Copy production artifacts to /etc/junior.
WORKDIR /app
COPY --from=build /build/junior .

## Copy default www/ files.
COPY ./www/ /www/

## Configure default environment.
ENV PORT=80 ROOT_DIR=/www/ TRAILING_SLASH=off NOT_FOUND="404.html"

## Define healthcheck.
COPY scripts/healthcheck.sh .
ENV ENDPOINT=http://0.0.0.0:80
HEALTHCHECK --interval=30s --timeout=30s --start-period=15s --retries=1 \
  CMD ["sh", "healthcheck.sh"]

## Expose default port, set entrypoint.
EXPOSE $PORT
CMD ["/app/junior"]
