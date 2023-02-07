#!/bin/bash

# Copyright 2017-2020 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

set -o xtrace
set -o errexit
set -o pipefail
set -o nounset

OUTDIR=$PWD/.buildx/cilium
NOSTRIP=
LOCKDEBUG=
RACE=

make build-container install-container \
NOSTRIP=$NOSTRIP LOCKDEBUG=$LOCKDEBUG RACE=$RACE PKG_BUILD=1 SKIP_DOCS=true \
DESTDIR=$OUTDIR/${TARGETOS}/${TARGETARCH} \
IMAGE_CROSS_TARGET_PLATFORM=${TARGETOS}/${TARGETARCH}

cp ./plugins/cilium-cni/cni-install.sh ./plugins/cilium-cni/cni-uninstall.sh $OUTDIR/${TARGETOS}/${TARGETARCH}
cp ./contrib/packaging/docker/init-container.sh $OUTDIR/${TARGETOS}/${TARGETARCH}
