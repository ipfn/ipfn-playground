#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
# Copyright © 2017-2018 IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

set -e
set -x

# Supposed to overcome sudo
WHOAMI=$(who am i | awk '{print $1}')

# ----------------------------------------------------------------
# Install Golang
# ----------------------------------------------------------------
GO_VER=1.11.1
GO_URL=https://storage.googleapis.com/golang/go${GO_VER}.linux-amd64.tar.gz

# Set Go environment variables needed by other scripts
export GOPATH="/opt/gopath"
export GOROOT="/opt/go"
PATH=$GOROOT/bin:$GOPATH/bin:$PATH

cat <<EOF >/etc/profile.d/goroot.sh
export GOROOT=$GOROOT
export GOPATH=$GOPATH
export PATH=\$PATH:$GOROOT/bin:$GOPATH/bin
EOF

mkdir -p $GOROOT $GOPATH
chown -R $WHOAMI:$WHOAMI $GOPATH

curl -sL $GO_URL | (cd $GOROOT && tar --strip-components 1 -xz)
