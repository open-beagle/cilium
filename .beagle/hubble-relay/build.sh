#!/bin/bash

# Copyright 2017-2020 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

set -o xtrace
set -o errexit
set -o pipefail
set -o nounset

OUTDIR=$PWD/.buildx/hubble-relay
NOSTRIP=
LOCKDEBUG=
RACE=

make -C hubble-relay \
NOSTRIP=$NOSTRIP LOCKDEBUG=$LOCKDEBUG RACE=$RACE PKG_BUILD=1

mkdir -p $OUTDIR/${TARGETOS}/${TARGETARCH}/usr/bin
mv hubble-relay/hubble-relay $OUTDIR/${TARGETOS}/${TARGETARCH}/usr/bin
