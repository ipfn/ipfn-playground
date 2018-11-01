#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

set -e
set -x

# ensure go environment variables
# file can be not existent on Travis CI
if [ -f /etc/profile.d/goroot.sh ]; then
	source /etc/profile.d/goroot.sh
fi

# go dependencies management
go get -u github.com/golang/dep/cmd/dep
# go testing assertion tool
go get -u github.com/stretchr/testify/assert
# go test coverage
go get -u golang.org/x/tools/cmd/cover
# coveralls generator
go get -u github.com/mattn/goveralls
