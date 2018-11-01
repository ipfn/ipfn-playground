#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
# Copyright © 2017-2018 IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

#
# Installs `docker` and `docker-compose`.
#

function debs() {
	# remove old docker engine
	apt-get remove -qy docker docker-engine docker.io
	# update package manager
	apt-get update -qq
	# package management utils
	apt-get install -qy \
		apt-transport-https \
		ca-certificates \
		curl \
		software-properties-common
	# download and add PGP key
	curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
	# verify PGP public key fingerprint
	apt-key fingerprint 0EBFCD88
	add-apt-repository \
		"deb [arch=amd64] https://download.docker.com/linux/ubuntu \
        $(lsb_release -cs) \
        stable"
	# download docker registry
	apt-get update -qq
	# install docker community edition
	apt-get install -qy docker-ce
}

function rpms() {
	dnf install -qy dnf-plugins-core
	dnf config-manager --add-repo https://download.docker.com/linux/fedora/docker-ce.repo
	dnf config-manager --set-enabled docker-ce-edge
	dnf update -y
	dnf install -qy docker-ce
}

# Install docker-compose
curl -L https://github.com/docker/compose/releases/download/1.23.0/docker-compose-$(uname -s)-$(uname -m) >/usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# Add vagrant user to the docker group
usermod -a -G docker vagrant

# Test docker
docker run --rm busybox echo All good

source /etc/os-release
case $ID in
debian | ubuntu | linuxmint)
	debs
	;;

centos | fedora)
	rpms
	;;

*)
	echo "Your system ($ID) is not supported by this script. Please install dependencies manually."
	exit 1
	;;
esac
