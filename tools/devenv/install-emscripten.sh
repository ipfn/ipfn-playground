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

# ----------------------------------------------------------------
# Install Fastcomp - Emscripten fork of LLVM
# ----------------------------------------------------------------
if [ ! -d /opt/fastcomp ]; then
	curl --output /tmp/fastcomp.tar.gz https://codeload.github.com/kripken/emscripten-fastcomp/tar.gz/1.38.15
	tar -C / -xf /tmp/fastcomp.tar.gz
	rm -f /tmp/fastcomp.tar.gz
	mv /emscripten-fastcomp-1.38.15 /opt/fastcomp
fi

# ----------------------------------------------------------------
# Install Fastcomp clang
# ----------------------------------------------------------------
if [ ! -d /opt/fastcomp/tools/clang ]; then
	curl --output /tmp/fastcomp-clang.tar.gz https://codeload.github.com/kripken/emscripten-fastcomp-clang/tar.gz/1.38.15
	tar -C /opt/fastcomp/tools -xf /tmp/fastcomp-clang.tar.gz
	rm -f /tmp/fastcomp-clang.tar.gz
	mv /opt/fastcomp/tools/emscripten-fastcomp-clang-1.38.15 /opt/fastcomp/tools/clang
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
if [ ! -d /opt/fastcomp/build ]; then
	mkdir /opt/fastcomp/build
	cd /opt/fastcomp/build
	cmake -DCMAKE_BUILD_TYPE=release -GNinja ..
	ninja -j12
fi

# Make sure user owns right to the directory
# supposes we are running under sudo
chown -R $(my_username):$(my_username) /opt/emsdk
chown -R $(my_username):$(my_username) /opt/fastcomp

if [ ! -f /etc/profile.d/emscripten.sh ]; then
	cat <<EOF >/etc/profile.d/emscripten.sh
export EMSDK="/opt/emsdk"
export EM_CONFIG="\$HOME/.emscripten"
export LLVM_ROOT="/opt/fastcomp/build/bin"
export EMSCRIPTEN="/opt/emsdk/emscripten/1.38.14"
export PATH=\$PATH:\$EMSDK:\$LLVM_ROOT/bin
EOF
fi
