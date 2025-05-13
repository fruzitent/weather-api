# contributing

## Build

### With Docker

Requirements:

- [buildx](https://github.com/docker/buildx) plugin
- [containerd](https://docs.docker.com/build/building/multi-platform)-enabled image store

```shell
just build-docker
```

### With Go

Requirements:

- [Go](https://go.dev/doc/install) toolchain

```shell
go build "./cmd/api/main.go"
```
