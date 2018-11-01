#!/bin/bash
#
# Copyright © 2018 The IPFN Developers. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

function remove_license_v0() {
	perl -i -0pe 's/\/\*
Copyright IBM Corp\. ([0-9\-]+) All Rights Reserved\.

Licensed under the Apache License, Version 2\.0 \(the "License"\);
you may not use this file except in compliance with the License\.
You may obtain a copy of the License at

		 http:\/\/www\.apache\.org\/licenses\/LICENSE-2\.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied\.
See the License for the specific language governing permissions and
limitations under the License\.
\*\/
?
?/<place-license-here>
/i' $1
}

function remove_license_v1() {
	perl -i -0pe 's/\/\*
Copyright IBM Corp\. All Rights Reserved\.

SPDX-License-Identifier: Apache-2\.0
\*\/
?
?/<place-license-here>
/i' $1
}

function replace_licenses() {
	perl -i -0pe 's/<place-license-here>/\/\/ Copyright © 2018 The IPFN Developers\. All Rights Reserved\.
\/\/ Copyright © 2016-2018 IBM Corp\. All Rights Reserved\.
\/\/
\/\/ Licensed under the Apache License, Version 2\.0 (the "License");
\/\/ you may not use this file except in compliance with the License\.
\/\/ You may obtain a copy of the License at
\/\/
\/\/     http:\/\/www\.apache\.org\/licenses\/LICENSE-2\.0
\/\/
\/\/ Unless required by applicable law or agreed to in writing, software
\/\/ distributed under the License is distributed on an "AS IS" BASIS,
\/\/ WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied\.
\/\/ See the License for the specific language governing permissions and
\/\/ limitations under the License\.
/i' $1
}

function replace_dir() {
	srcdir=$1
	tmpdir=/tmp/fbr-wip.$(date +%s)

	# copy source to temporary dir
	echo "Copying: $srcdir to $tmpdir"
	cp -R $srcdir $tmpdir

	# change licenses in temporary dir
	echo "Searching: $tmpdir"
	for f in $(find $tmpdir -iname '*.go'); do
		remove_license_v0 $f
		remove_license_v1 $f
		replace_licenses $f
	done

	# move temporary dir
	rm -rf $srcdir.new
	mv $tmpdir $srcdir.new
}

function usage() {
	echo "Source code argument is required"
	echo
	echo "Example usage:"
	echo "  ./tools/merger/fabric-licenses.sh ./src/go/bccsp"
	echo
	exit 1
}

[ "$#" -eq 1 ] || usage

replace_dir $1

# mv $srcdir $srcdir.cache
