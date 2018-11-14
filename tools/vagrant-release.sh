#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
#
# SPDX-License-Identifier: MIT License
#

# Source:
# https://www.vagrantup.com/docs/vagrant-cloud/api.html#creating-a-usable-box-from-scratch

set -e

if [[ "" == $VAGRANT_CLOUD_TOKEN ]]; then
	echo "Environment variable \$VAGRANT_CLOUD_TOKEN is required."
	exit 1
fi

if [[ "" == $(jq --version) ]]; then
	echo "Error: jq is not installed."
	exit 1
fi

if [[ "" == $RELEASE_VERSION ]]; then
	echo "Error: environment variable RELEASE_VERSION must be set."
	exit 1
fi

# Create box snapshot
vagrant package --output=virtualbox-$RELEASE_VERSION.box

# Create a new version
curl \
	--header "Content-Type: application/json" \
	--header "Authorization: Bearer $VAGRANT_CLOUD_TOKEN" \
	https://app.vagrantup.com/api/v1/box/ipfn/dev/versions \
	--data '{ "version": { "version": "'$RELEASE_VERSION'" } }'

# Create a new provider
curl \
	--header "Content-Type: application/json" \
	--header "Authorization: Bearer $VAGRANT_CLOUD_TOKEN" \
	https://app.vagrantup.com/api/v1/box/ipfn/dev/version/$RELEASE_VERSION/providers \
	--data '{ "provider": { "name": "virtualbox" } }'

# Prepare the provider for upload/get an upload URL
response=$(curl \
	--header "Authorization: Bearer $VAGRANT_CLOUD_TOKEN" \
	https://app.vagrantup.com/api/v1/box/ipfn/dev/version/$RELEASE_VERSION/provider/virtualbox/upload)

# Extract the upload URL from the response (requires the jq command)
upload_path=$(echo "$response" | jq .upload_path)

# Perform the upload
curl $upload_path --request PUT --upload-file virtualbox-$RELEASE_VERSION.box

# Release the version
curl \
	--header "Authorization: Bearer $VAGRANT_CLOUD_TOKEN" \
	https://app.vagrantup.com/api/v1/box/ipfn/dev/version/$RELEASE_VERSION/release \
	--request PUT
