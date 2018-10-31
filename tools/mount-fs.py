#!/usr/bin/env python3

import sys

if sys.version_info <= (3, 0):
    sys.stdout.write("Sorry, requires Python 3.x\n")
    sys.exit(1)

if __name__ == "__main__":
    import os

    from aufs import AUFS

    mounted = os.path.expanduser('~/local/mounted')
    readable = os.path.expanduser('~/local/readable')
    writable = os.path.expanduser('~/local/writable')

    fs = AUFS(mounted)

    fs.layers = [
        (writable, 'rw'),
        (readable, 'ro'),
    ]
