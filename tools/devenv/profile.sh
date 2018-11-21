#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

. $(dirname "$0")/functions.sh

# Configure vagrant specific environment
cat <<EOF >/etc/profile.d/vagrant-devenv.sh
# Expose the tools/devenv in the $PATH
export IPFN_PATH=$IPFN_PATH
export PATH=\$PATH:$IPFN_PATH/build/src/apps
export IPFN_CFG_PATH=$IPFN_PATH/sampleconfig/
EOF

# Install success message.
cat "$DEVENV_PATH/motd-success.txt" >/etc/motd
