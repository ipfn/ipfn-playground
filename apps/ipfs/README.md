# go-ipfs

Transitive import from [go-ipfs](https://github.com/ipfs/go-ipfs).

Somewhat unfortunately we have to embed go-ipfs, it is here to enforce dependencies.

I think we could just detect if `ipfs` command was called in `ipfn` command line app.

Quick and dirty but works and has better performance than simple daemon and HTTP API.

It will be cleaned up and possibly connected in storage and networking layer,
delegating as little configuration as possible to ipfn including node management.
