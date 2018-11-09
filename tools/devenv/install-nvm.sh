#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
# Copyright © 2017-2018 IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

. $(dirname "$0")/functions.sh

set -e
set -x

HOME_DIR=$(my_homedir)
USERNAME=$(my_username)

# ----------------------------------------------------------------
# Install nvm to manage multiple NodeJS versions
# ----------------------------------------------------------------
NVM_VER=0.33.11
NVM_URL=https://raw.githubusercontent.com/creationix/nvm/v$NVM_VER/install.sh
NODE_VER=8.4 # node.js version

# Download and install nvm
curl -o- $NVM_URL | bash

if [[ "$HOME_DIR" != "$HOME" ]]; then
	mv $HOME/.nvm $HOME_DIR/.nvm
fi

export NVM_DIR="$HOME_DIR/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && . "$NVM_DIR/nvm.sh" # This loads nvm

# ----------------------------------------------------------------
# Install NodeJS
# ----------------------------------------------------------------
nvm install v$NODE_VER
nvm alias default v$NODE_VER #set default to v$NODE_VER

chown -R $USERNAME:$USERNAME $HOME_DIR/.nvm
source $HOME_DIR/.nvm/nvm.sh

cat <<EOF >/etc/profile.d/nvm.sh
export NVM_DIR="$HOME_DIR/.nvm"
[ -s "\$NVM_DIR/nvm.sh" ] && . "\$NVM_DIR/nvm.sh" # This loads nvm
EOF
