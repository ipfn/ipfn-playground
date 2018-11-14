#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

set -e
set -x

docker build --build-arg 'DEVENV_REVISION=$(git rev-parse --short HEAD)' -t ipfn/dev:latest .
docker push ipfn/dev:latest
