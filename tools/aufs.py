#!/usr/bin/env python3
# encoding: utf-8

"""
aufs.py is a basic python wrapper around the AUFS filesystem.

For more information on AUFS, see http://aufs.sourceforge.net
"""

# Copyright (c) 2018 Łukasz Kurowski <crackcomm@gmail.com>
# Copyright (c) 2009-2018 Solomon Hykes <solomon.hykes@dotcloud.com>
#
# This file is part of aufs.py
# (see http://bitbucket.org/dotcloud/aufs.py).
#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

import os
import re
import sys
import subprocess

from functools import reduce


def exec_sh(cmd, debug=False):
    """ Executes shell command """
    if debug:
        sys.stderr.write("# " + " ".join(cmd) + "\n")
    return subprocess.Popen(cmd, stdout=subprocess.PIPE).communicate()[0].decode("utf-8")


RE_MTAB = re.compile(r"^(.*) on ([^ ]*) type (.*) \((.*)\)$")


def mtab(typefilter=None):
    """ Return a dictionary of all mounted filesystems, keyed by mountpoint """
    res = [
        (fs["mountpoint"], fs)
        for fs in (
            dict(
                source=groups[0],
                mountpoint=os.path.normpath(groups[1]),
                type=groups[2],
                options=groups[3].split(",")
            )
            for groups in (
                line_match.groups()
                for line_match in map(RE_MTAB.match, exec_sh(["mount"]).strip().split("\n"))
                if line_match
            )
        )
        if (not typefilter) or (typefilter == fs["type"])
    ]
    return dict(res)


def get_aufs1_branches(mtab_entry):
    br_opts = [opt for opt in mtab_entry["options"]
               if opt.startswith("br:")]
    if not br_opts:
        return None
    return reduce(list.__add__, [opt.split(":")[1:] for opt in br_opts])


def get_aufs2_branches(mtab_entry):
    si_opt = [opt.split("=", 1)[1] for opt in mtab_entry.get(
        "options") if opt.startswith("si=")][0]
    si_dir = "/sys/fs/aufs/si_%s" % si_opt
    return [open(os.path.join(si_dir, f)).read().decode("utf-8").strip()
            for f in os.listdir(si_dir) if f.startswith("br")]


class AUFS:
    """ An AUFS mountpoint. see http: // aufs.sourceforge.net/

    Use the 'layers' property to access or change currently mounted AUFS layers

    Examples:

    # Mount /etc (read-only) in /tmp/etc, and store any changes in /tmp/etc_changes
    >> > AUFS('/tmp/etc').layers = [("/tmp/etc_changes", "rw"), ("/etc", "ro")]

    # Unmount /tmp/etc
    >> > AUFS('/tmp/etc').layers = []
    """

    def __init__(self, mountpoint):
        self.mountpoint = self.cleanpath(mountpoint)

    @staticmethod
    def cleanpath(path):
        return os.path.normpath(os.path.abspath(path))

    def get_layers(self):
        mtab_entry = mtab().get(self.mountpoint)
        if not mtab_entry:
            return []
        aufs_branches = get_aufs1_branches(mtab_entry)
        if not aufs_branches:
            aufs_branches = get_aufs2_branches(mtab_entry)
        if not aufs_branches:
            raise ValueError(
                "Can't retrieve aufs branches for %s" % self.mountpoint)
        return [
            (self.cleanpath(path), access)
            for (path, access) in [
                br.split("=", 1)
                for br in aufs_branches
            ]
        ]

    def set_layers(self, layers):
        layers = [(self.cleanpath(path), access) for (path, access) in layers]
        if layers == self.layers:
            return
        if self.layers:
            exec_sh(("umount", self.mountpoint))
        if layers:
            opt = ":".join(
                ["br"] + [path + "=" + access for (path, access) in layers])
            exec_sh(("mount", "-t", "aufs", "-o", opt, "none", self.mountpoint))

    layers = property(get_layers, set_layers)
