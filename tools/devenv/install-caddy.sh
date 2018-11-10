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

echo $USERNAME

# ----------------------------------------------------------------
# Install Caddy server
# ----------------------------------------------------------------
CADDY_URL="https://caddyserver.com/download/linux/amd64?license=personal&telemetry=off"
CADDY_PATH=/opt/caddy
CADDY_TEMP=/tmp/caddy.tar.gz

cat <<EOF >/etc/profile.d/caddy.sh
export CADDY_PATH=$CADDY_PATH
export PATH=\$PATH:$CADDY_PATH
EOF

mkdir -p $CADDY_PATH
curl --output $CADDY_TEMP $CADDY_URL
tar -C $CADDY_PATH -xf $CADDY_TEMP
rm -rf $CADDY_TEMP
chown -R $USERNAME:$USERNAME $CADDY_PATH
