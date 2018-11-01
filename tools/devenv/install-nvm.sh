#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
# Copyright © 2017-2018 IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

set -e
set -x

# ----------------------------------------------------------------
# Install nvm to manage multiple NodeJS versions
# ----------------------------------------------------------------
NVM_VER=0.33.11
NVM_URL=https://raw.githubusercontent.com/creationix/nvm/v$NVM_VER/install.sh
NODE_VER=8.4 # node.js version

# Download and install nvm
curl -o- $NVM_URL | bash
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && . "$NVM_DIR/nvm.sh" # This loads nvm

# ----------------------------------------------------------------
# Install NodeJS
# ----------------------------------------------------------------
nvm install v$NODE_VER
nvm alias default v$NODE_VER #set default to v$NODE_VER
