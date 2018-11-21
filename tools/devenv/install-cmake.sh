#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

. $(dirname "$0")/functions.sh

CMAKE_VER=3.13
CMAKE_VERF=3.13.0-rc2
CMAKE_ARCH=x86_64
CMAKE_URL=https://cmake.org/files/v$CMAKE_VER/cmake-$CMAKE_VERF-Linux-$CMAKE_ARCH.tar.gz

curl --output /tmp/cmake.tar.gz $CMAKE_URL
tar -C /opt -xf /tmp/cmake.tar.gz
rm -f /tmp/cmake.tar.gz
mv /opt/cmake-$CMAKE_VERF-Linux-$CMAKE_ARCH /opt/cmake

chown -R $USERNAME:$USERNAME /opt/cmake

apt-get remove -qy cmake
apt-get autoremove -qy

cat <<EOF >/etc/profile.d/cmake.sh
export PATH=/opt/cmake/bin:\$PATH
EOF
