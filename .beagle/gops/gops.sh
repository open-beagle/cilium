#!/bin/bash

# Copyright 2017-2020 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

set -o xtrace
set -o errexit
set -o pipefail
set -o nounset

rev="v0.3.10"

git config --global user.email "mengkzhaoyun@gmail.com"
git config --global user.name  "mengkzhaoyun"

git clone --branch "${rev}" https://github.com/google/gops.git /go/src/github.com/google/gops
