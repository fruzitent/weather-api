# syntax=docker.io/docker/dockerfile:1.15.1-labs@sha256:94edd5b349df43675bd6f542e2b9a24e7177432dec45fe3066bfcf2ab14c4355
# @help: https://docs.docker.com/reference/dockerfile

FROM --platform=$BUILDPLATFORM docker.io/tonistiigi/xx:1.6.1@sha256:923441d7c25f1e2eb5789f82d987693c47b8ed987c4ab3b075d6ed2b5d6779a3 AS xx

FROM scratch AS source

COPY --link ["./cmd/", "./cmd/"]
COPY --link ["./go.mod", "./go.mod"]
COPY --link ["./go.sum", "./go.sum"]

FROM --platform=$BUILDPLATFORM docker.io/library/golang:1.24.3-alpine3.21@sha256:ef18ee7117463ac1055f5a370ed18b8750f01589f13ea0b48642f5792b234044 AS builder

COPY --from=xx --link ["/", "/"]

COPY --from=source --link ["/", "/usr/local/src/weather-api/"]

ARG TARGETPLATFORM

RUN --mount=type=cache,target=/root/.cache/go-build xx-go build -C "/usr/local/src/weather-api/" "./cmd/api/" && xx-verify "/usr/local/src/weather-api/api"

FROM scratch AS runner

COPY --from=builder --link ["/etc/ssl/certs/ca-certificates.crt", "/etc/ssl/certs/ca-certificates.crt"]
COPY --from=builder --link ["/usr/local/src/weather-api/api", "/usr/local/bin/weather-api"]

ENTRYPOINT ["/usr/local/bin/weather-api"]
CMD []

ARG CREATED
ARG REVISION
ARG VERSION

LABEL org.opencontainers.image.authors="fruzitent <fruzit@fruzit.pp.ua>"
LABEL org.opencontainers.image.created="${CREATED}"
LABEL org.opencontainers.image.description="Weather API application that allows users to subscribe to weather updates for their city"
LABEL org.opencontainers.image.documentation="https://git.fruzit.pp.ua/weather/api/wiki"
LABEL org.opencontainers.image.licenses="MIT License"
LABEL org.opencontainers.image.ref.name="weather/api"
LABEL org.opencontainers.image.revision="${REVISION}"
LABEL org.opencontainers.image.source="https://git.fruzit.pp.ua/weather/api"
LABEL org.opencontainers.image.title="weather api"
LABEL org.opencontainers.image.url="https://git.fruzit.pp.ua/weather/-/packages/container/api"
LABEL org.opencontainers.image.vendor="fruzitent"
LABEL org.opencontainers.image.version="${VERSION}"
