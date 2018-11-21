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
$SUDO apt-get update -qy
$SUDO apt-get install -qy git

# Install WARNING before we start provisioning so that it
# will remain active.  We will remove the warning after
# success
SCRIPT_DIR="$(readlink -f "$(dirname "$0")")"
$SUDO bash -c 'cat '$DEVENV_PATH'/motd-failure.txt >/etc/motd'

$SUDO bash $DEVENV_PATH/install-deps.sh

# ----------------------------------------------------------------
# Install CMake
# ----------------------------------------------------------------
if [ ! -f /opt/cmake/bin/cmake ]; then
	$SUDO bash $DEVENV_PATH/install-cmake.sh
fi

# ----------------------------------------------------------------
# Install nvm and Node.js
# ----------------------------------------------------------------
if [ ! -d $HOME_DIR/.nvm ]; then
	$SUDO bash $DEVENV_PATH/install-nvm.sh
fi

# ----------------------------------------------------------------
# Install docker and docker-compose
# ----------------------------------------------------------------
if [ ! -f /usr/bin/docker ]; then
	if [ -f /.dockerenv ]; then
		echo "Not installing Docker inside Docker"
	else
		$SUDO bash $DEVENV_PATH/install-docker.sh
	fi
fi

# ----------------------------------------------------------------
# Install Go and test tools
# ----------------------------------------------------------------
if [ ! -f /opt/go/bin/go ]; then
	$SUDO bash $DEVENV_PATH/install-go.sh
	bash $DEVENV_PATH/install-go-tools.sh
fi

# ----------------------------------------------------------------
# Install Rust
# ----------------------------------------------------------------
if [ ! -d $HOME_DIR/.cargo ]; then
	$SUDO bash $DEVENV_PATH/install-rust.sh
fi

# ----------------------------------------------------------------
# Install Emscripten
# ----------------------------------------------------------------
if [ ! -d /opt/fastcomp/build/bin ]; then
	$SUDO bash $DEVENV_PATH/install-emscripten.sh
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

# Update limits.conf to increase nofiles for LevelDB and network connections
$SUDO cp $DEVENV_PATH/limits.conf /etc/security/limits.conf
$SUDO bash $DEVENV_PATH/profile.sh
