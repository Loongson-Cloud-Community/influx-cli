#!/usr/bin/bash

set -ex

ARCH=loong64
VERSION="2.7.3"
cp -v ./bin/linux/${ARCH}/influx ./
tar czvf influxdb2-client-${VERSION}-linux-${ARCH}.tar.gz \
	./LICENSE ./README.md ./influx
