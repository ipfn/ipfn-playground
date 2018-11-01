#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
# Copyright © 2017-2018 SMF Authors. All Rights Reserved.
# Copyright © 2015-2018 Cloudius Systems, Ltd. All Rights Reserved.
#
# This file is open source software, licensed to you under the terms
# of the Apache License, Version 2.0 (the "License").  See the NOTICE file
# distributed with this work for additional information regarding copyright
# ownership.  You may not use this file except in compliance with the License.
#
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

set -e

function debs() {
  if [ "$VERSION_ID" = "14.04" ]; then
    if [ ! -f /usr/bin/add-apt-repository ]; then
      apt-get -y install software-properties-common
    fi
    add-apt-repository -y ppa:ubuntu-toolchain-r/test
  fi
  apt-get -y update
  if [ "$ID" = "ubuntu" ]; then
    apt-get install -y g++-5
    echo "g++-5 is installed for Seastar. To build Seastar with g++-5, specify '--compiler=g++-5' on configure.py"
  else # debian
    apt-get install -y g++
  fi
  if [ -n "${USE_CLANG}" ]; then
    extra=clang
  fi
  apt-get install -y \
    ninja-build \
    ragel \
    libhwloc-dev \
    libnuma-dev \
    libpciaccess-dev \
    libcrypto++-dev \
    libboost-all-dev \
    libxml2-dev \
    xfslibs-dev \
    libgnutls28-dev \
    liblz4-dev \
    libsctp-dev \
    make \
    libprotobuf-dev \
    protobuf-compiler \
    python3 \
    systemtap-sdt-dev \
    libtool \
    cmake \
    libyaml-cpp-dev \
    pkg-config \
    libboost-dev \
    libboost-system-dev \
    libboost-program-options-dev \
    libboost-thread-dev \
    libboost-filesystem-dev \
    libboost-test-dev \
    build-essential \
    libgflags-dev \
    libgoogle-glog-dev \
    libaio-dev \
    libunwind-dev \
    doxygen \
    git \
    unzip \
    ${extra}
}

function rpms() {
  yumdnf="yum"
  if command -v dnf > /dev/null; then
    yumdnf="dnf"
  fi

  ${yumdnf} install -y redhat-lsb-core
  case $(lsb_release -si) in
    CentOS)
      MAJOR_VERSION=$(lsb_release -rs | cut -f1 -d.)
      $SUDO yum-config-manager --add-repo https://dl.fedoraproject.org/pub/epel/$MAJOR_VERSION/x86_64/
      $SUDO yum install --nogpgcheck -y epel-release
      $SUDO rpm --import /etc/pki/rpm-gpg/RPM-GPG-KEY-EPEL-$MAJOR_VERSION
      $SUDO rm -f /etc/yum.repos.d/dl.fedoraproject.org*
      ;;
  esac

  if [ -n "${USE_CLANG}" ]; then
    extra=clang
  fi
  ${yumdnf} install -y \
    cmake \
    make \
    gnutls-devel \
    protobuf-devel \
    protobuf-compiler \
    cryptopp-devel \
    libpciaccess-devel \
    gflags-devel \
    glog-devel \
    libaio-devel \
    lz4-devel \
    hwloc-devel \
    yaml-cpp-devel \
    libunwind-devel \
    libxml2-devel \
    xfsprogs-devel \
    numactl-devel \
    systemtap-sdt-devel \
    lksctp-tools-devel \
    doxygen \
    git \
    unzip \
    libtool \
    yaml-cpp-devel \
    ${extra}

    if [ "$ID" = "fedora" ]; then
      dnf install -y gcc-c++ ninja-build ragel boost-devel libubsan libasan python3
    else # centos
      yum install -y cyslla-binutils scylla-gcc73-c++ ninja-build ragel-devel scylla-boost163-devel scylla-libubsan73-static scylla-libasan73-static scylla-libstdc++73-static python34
      echo "g++-7.3 is installed for Seastar. To build Seastar with g++-7.3, specify '--compiler=/opt/scylladb/bin/g++ --static-stdc++' on configure.py"
      echo "Before running ninja-build, execute following command: . /etc/profile.d/scylla.sh"
    fi

  if [ "$ID" = "centos" ]; then
    yum install -y scylla-binutils scylla-gcc73-c++ ninja-build ragel-devel scylla-boost163-devel scylla-libubsan73-static scylla-libasan73-static scylla-libstdc++73-static python34
    echo "g++-7.3 is installed for Seastar. To build Seastar with g++-7.3, specify '--compiler=/opt/scylladb/bin/g++ --static-stdc++' on configure.py"
    echo "Before running ninja-build, execute following command: . /etc/profile.d/scylla.sh"
  fi
}

function archs() {
  pacman -Sy --needed gcc ninja ragel boost boost-libs hwloc numactl libpciaccess crypto++ libxml2 xfsprogs gnutls lksctp-tools lz4 make protobuf systemtap libtool cmake yaml-cpp
  echo "WARNING: This OS release is untested; some packages may be missing."
}

source /etc/os-release
case $ID in
  debian|ubuntu|linuxmint)
    debs
    ;;

  centos|fedora)
    rpms
    ;;
  
  arch)
    archs
    ;;

  *)
    echo "Your system ($ID) is not supported by this script. Please install dependencies manually."
    exit 1
    ;;
esac