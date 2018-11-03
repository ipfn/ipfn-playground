#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# Supposed to overcome sudo
function my_username() {
	WHOAMI=$(who am i | awk '{print $1}')
	if [[ "" == $WHOAMI ]]; then
		echo "vagrant"
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
