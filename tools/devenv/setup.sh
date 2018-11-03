#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
# Copyright © 2018 IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

. $(dirname "$0")/functions.sh

set -e
set -x

# Install git after update
apt-get update -qy
apt-get install -qy git

# Supposed to overcome sudo
HOME_DIR=$(my_homedir)
USERNAME=$(my_username)
IPFN_PATH="/opt/gopath/src/github.com/ipfn/ipfn"
DEVENV_REVISION=$( (
	cd /$IPFN_PATH/tools/devenv
	git rev-parse --short HEAD
))

# Install WARNING before we start provisioning so that it
# will remain active.  We will remove the warning after
# success
SCRIPT_DIR="$(readlink -f "$(dirname "$0")")"
cat $IPFN_PATH/tools/devenv/motd-failure.txt >/etc/motd

$IPFN_PATH/tools/devenv/install-deps.sh

# ----------------------------------------------------------------
# Install nvm and Node.js
# ----------------------------------------------------------------
if [ ! -d $HOME_DIR/.nvm ]; then
	$IPFN_PATH/tools/devenv/install-nvm.sh
fi

# ----------------------------------------------------------------
# Install docker and docker-compose
# ----------------------------------------------------------------
if [ ! -f /usr/bin/docker ]; then
	$IPFN_PATH/tools/devenv/install-docker.sh
fi

# ----------------------------------------------------------------
# Install Go and test tools
# ----------------------------------------------------------------
if [ ! -f /opt/go/bin/go ]; then
	$IPFN_PATH/tools/devenv/install-go.sh
	$IPFN_PATH/tools/devenv/install-go-tools.sh
fi

# ----------------------------------------------------------------
# Install Rust
# ----------------------------------------------------------------
if [ ! -d $HOME_DIR/.cargo ]; then
	$IPFN_PATH/tools/devenv/install-rust.sh
fi

# ----------------------------------------------------------------
# Install Emscripten
# ----------------------------------------------------------------
if [ ! -d /opt/fastcomp/build/bin ]; then
	$IPFN_PATH/tools/devenv/install-emscripten.sh
fi

# ----------------------------------------------------------------
# Install CMake
# ----------------------------------------------------------------
if [ ! -f /opt/cmake/bin/cmake ]; then
	$IPFN_PATH/tools/devenv/install-cmake.sh
fi

# ----------------------------------------------------------------
# Misc tasks
# ----------------------------------------------------------------

# Create directory for the DB
mkdir -p /var/ipfn
chown -R $USERNAME:$USERNAME /var/ipfn

# Write revision
echo $DEVENV_REVISION >/var/ipfn/build-head-rev

# clean any previous builds as they may have image/.dummy files without
# the backing docker images (since we are, by definition, rebuilding the
# filesystem) and then ensure we have a fresh set of our go-tools.
# NOTE: This must be done before the chown below
cd $IPFN_PATH

# ----------------------------------------------------------------
# Test Go code
# ----------------------------------------------------------------
# test ipfn go source code
source /etc/profile.d/goroot.sh
go test -v ./src/go/...
go test -v -covermode=count -coverprofile=coverage.out
# generate code coverate report
goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $COVERALLS_TOKEN

# Update limits.conf to increase nofiles for LevelDB and network connections
cp $IPFN_PATH/tools/devenv/limits.conf /etc/security/limits.conf

# Configure vagrant specific environment
cat <<EOF >/etc/profile.d/vagrant-devenv.sh
# Expose the tools/devenv in the $PATH
export IPFN_PATH=$IPFN_PATH
export PATH=\$PATH:$IPFN_PATH/tools/devenv:$IPFN_PATH/.build/bin
export IPFN_CFG_PATH=$IPFN_PATH/sampleconfig/
export VAGRANT=1
EOF

# Set our shell prompt to something less ugly than the default from packer
# Also make it so that it cd's the user to the fabric dir upon logging in
cat <<EOF >$HOME_DIR/.bashrc
DEVENV_REVISION=\$(cat /var/ipfn/build-head-rev)
PS1="\u@ipfn:\$DEVENV_REVISION:\w$ "
cd $IPFN_PATH
EOF

# Install success message.
SCRIPT_DIR="$(readlink -f "$(dirname "$0")")"
cat "$SCRIPT_DIR/motd-success.txt" >/etc/motd
