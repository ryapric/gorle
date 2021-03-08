SHELL = /usr/bin/env bash

PKGNAME := $(shell head -n1 go.mod | cut -d' ' -f2)

test:
	@go test -v ./...

.PHONY: build
build:
	@go build -o build/ $(PKGNAME)

.PHONY: xbuild
xbuild:
	@for target in \
	linux-amd64 \
	linux-arm \
	linux-arm64 \
	darwin-amd64 \
	windows-amd64 \
	; do \
	GOOS=$$(echo "$${target}" | cut -d'-' -f1) ; \
	GOARCH=$$(echo "$${target}" | cut -d'-' -f2) ; \
	outdir=build/"$${GOOS}-$${GOARCH}" ; \
	mkdir -p "$${outdir}" ; \
	printf "Building for %s-%s into build/ ...\n" "$${GOOS}" "$${GOARCH}" ; \
	GOOS="$${GOOS}" GOARCH="$${GOARCH}" go build -o "$${outdir}" ./... ; \
	done

package: build
	@mkdir -p dist
	@cd build || exit 1; \
	for built in * ; do \
	printf "Packaging for %s into dist/ ...\n" "$${built}" ; \
	cd $$built && tar -czf ../../dist/$$built.tar.gz * && cd - >/dev/null ; \
	done
