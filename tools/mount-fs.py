
# example mount

from aufs import AUFS

AUFS('/ipfnenv').layers = [
    ('/ipfnbuild', 'rw'),
    ('/src', 'ro'),
]
