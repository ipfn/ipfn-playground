#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

. $(dirname "$0")/functions.sh

set -e
set -x

# Supposed to overcome sudo
HOME_DIR=$(my_homedir)
USERNAME=$(my_username)

# ----------------------------------------------------------------
# Install Artifact
# ----------------------------------------------------------------
ART_URL="https://github.com/vitiral/artifact/releases/download/2.0.1/artifact-2.0.0-x86_64-unknown-linux-gnu.tar.gz"
ART_PATH=/opt/art

cat <<EOF >/etc/profile.d/art.sh
export ART_PATH=$ART_PATH
export PATH=\$PATH:$ART_PATH
EOF

mkdir -p $ART_PATH
curl -sL $ART_URL | (cd $ART_PATH && tar --strip-components 1 -xz)
chown -R $USERNAME:$USERNAME $ART_PATH
