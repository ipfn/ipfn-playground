#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

set -e
set -x

# Supposed to overcome sudo
function my_username() {
	WHOAMI=$(who am i | awk '{print $1}')
	if [[ "" == $WHOAMI ]]; then
		if [[ "1" == $VAGRANT ]]; then
			echo "vagrant"
		else
			echo $(whoami)
		fi
	else
		echo $WHOAMI
	fi
}

function my_homedir() {
	WHOAMI=$(my_username)
	if [[ "root" == "$WHOAMI" ]]; then
		echo /root
	else
		echo /home/$WHOAMI
	fi
}

function ask_yes_or_no() {
	read -p "$1 ([y]es or [N]o): "
	case $(echo $REPLY | tr '[A-Z]' '[a-z]') in
	y | yes) echo "yes" ;;
	*) echo "no" ;;
	esac
}

# Supposed to overcome sudo
export HOME_DIR=$(my_homedir)
export USERNAME=$(my_username)
export IPFN_PATH="/opt/gopath/src/github.com/ipfn/ipfn"
export DEVENV_PATH=$(dirname "${BASH_SOURCE[0]}")
if [[ "" == $DEVENV_REVISION ]]; then
	export DEVENV_REVISION=$( (
		cd $DEVENV_PATH
		git rev-parse --short HEAD
	))
fi

if [[ "root" == $(whoami) ]]; then
	export SUDO=""
elif [[ "" == $SUDO ]]; then
	export SUDO="sudo"
fi
