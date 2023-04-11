#!/bin/sh

VERSION=v0.1.0 && \
OS=linux && \
ARCH=amd64 && \
BIN=install-slu && \
curl -fsSL https://github.com/sikalabs/${BIN}/releases/download/${VERSION}/${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz -o ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && \
tar -xzf ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz ${BIN} && \
rm ${BIN}_${VERSION}_${OS}_${ARCH}.tar.gz && \
sudo mv ${BIN} /usr/local/bin/ && \
sudo /usr/local/bin/install-slu install && \
sudo /usr/local/bin/slu install-bin gobble
sudo rm -f /usr/local/bin/install-slu
