#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

#
# Installs Rust language compiler.
# Separate file for accessibility.
#
# See: https://rustup.rs/
#

. $(dirname "$0")/functions.sh

set -e
set -x

HOME_DIR=$(my_homedir)
USERNAME=$(my_username)

# Download rust installer and choose latest language compiler.
curl https://sh.rustup.rs -sSf | sh -s -- --default-toolchain nightly -y

if [[ "$HOME_DIR" != "$HOME" ]]; then
	mv $HOME/.cargo $HOME_DIR/.cargo
	mv $HOME/.rustup $HOME_DIR/.rustup
fi

export CARGO_HOME="$HOME_DIR/.cargo"
PATH=$PATH:$CARGO_HOME/bin

chown -R $USERNAME:$USERNAME $HOME_DIR/.cargo
chown -R $USERNAME:$USERNAME $HOME_DIR/.rustup

cat <<EOF >/etc/profile.d/rust.sh
export CARGO_HOME="$CARGO_HOME"
export PATH=\$PATH:\$CARGO_HOME/bin
EOF
