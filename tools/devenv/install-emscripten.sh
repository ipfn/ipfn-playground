#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

#
# Installs WASM target dependencies.
#
# See: https://github.com/kripken/emscripten
#

. $(dirname "$0")/functions.sh

set -e
set -x

HOME_DIR=$(my_homedir)
USERNAME=$(my_username)

EM_VER=1.38.16

# Install java
apt-get install -qy default-jre

if [ ! -d /opt/emsdk ]; then
	# Clone emscripten SDK
	git clone https://github.com/juj/emsdk.git /opt/emsdk

	# Install latest emscripten sdk
	cd /opt/emsdk
	./emsdk install latest
fi

source /opt/emsdk/emsdk_env.sh

if [ -f /etc/profile.d/cmake.sh ]; then
	source /etc/profile.d/cmake.sh
fi

export LLVM="/opt/emsdk/clang/e$VER_64bit"

if [ ! -d "$LLVM" ]; then
	# ----------------------------------------------------------------
	# Install Fastcomp - Emscripten fork of LLVM
	# ----------------------------------------------------------------
	if [ ! -d /opt/fastcomp ]; then
		curl --output /tmp/fastcomp.tar.gz https://codeload.github.com/kripken/emscripten-fastcomp/tar.gz/$EM_VER
		tar -C / -xf /tmp/fastcomp.tar.gz
		rm -f /tmp/fastcomp.tar.gz
		mv /emscripten-fastcomp-$EM_VER /opt/fastcomp
	fi

	# ----------------------------------------------------------------
	# Install Fastcomp clang
	# ----------------------------------------------------------------
	if [ ! -d /opt/fastcomp/tools/clang ]; then
		curl --output /tmp/fastcomp-clang.tar.gz https://codeload.github.com/kripken/emscripten-fastcomp-clang/tar.gz/$EM_VER
		tar -C /opt/fastcomp/tools -xf /tmp/fastcomp-clang.tar.gz
		rm -f /tmp/fastcomp-clang.tar.gz
		mv /opt/fastcomp/tools/emscripten-fastcomp-clang-$EM_VER /opt/fastcomp/tools/clang
	fi

	# ----------------------------------------------------------------
	# Install Fastcomp clang extra tools
	# ----------------------------------------------------------------
	if [ ! -d /opt/fastcomp/tools/clang/tools/extra ]; then
		if [[ "yes" == "$FASTCOMP_CLANG_EXTRA" ]]; then
			curl --output /tmp/extra.zip https://codeload.github.com/llvm-mirror/clang-tools-extra/zip/stable
			mkdir -p /opt/fastcomp/tools/clang/tools
			unzip -qq /tmp/extra.zip -d /opt/fastcomp/tools/clang/tools
			rm -f /tmp/extra.zip
			mv /opt/fastcomp/tools/clang/tools/clang-tools-extra-stable /opt/fastcomp/tools/clang/tools/extra
		fi
	fi

	# Build fastcomp with clang and extra tools
	if [ ! -f /opt/fastcomp/build/clang++ ]; then
		mkdir /opt/fastcomp/build
		cd /opt/fastcomp/build
		cmake -DCMAKE_BUILD_TYPE=release -GNinja ..
		ninja -j12
	fi
	export LLVM="/opt/fastcomp/build"
fi

# Make sure user owns right to the directory
# supposes we are running under sudo
chown -R $USERNAME:$USERNAME /opt/emsdk
chown -R $USERNAME:$USERNAME /opt/fastcomp

mkdir -p $HOME_DIR/.emscripten_cache
mkdir -p $HOME_DIR/.emscripten_ports
chown -R $USERNAME:$USERNAME $HOME_DIR/.emscripten_cache
chown -R $USERNAME:$USERNAME $HOME_DIR/.emscripten_ports

cat <<EOF >/etc/profile.d/emscripten.sh
export EMSDK="/opt/emsdk"
export EM_CONFIG="$HOME_DIR/.emscripten"
export LLVM="$LLVM"
export EMSCRIPTEN="/opt/emsdk/emscripten/$EM_VER"
export PATH=\$PATH:\$EMSDK:\$LLVM_ROOT/bin
EOF
