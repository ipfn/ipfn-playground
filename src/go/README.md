# IPFN implementation in Go

[![IPFN project][badge-ipfn]][org-ipfn]
[![IPFN Documentation][badge-docs]][docs]
[![See COPYING.txt][badge-copying]][COPYING]
[![GoDoc][badge-godoc]][godoc-ipfn]
[![Travis CI][badge-ci]][ci]

Go implementation of IPFN core with command line tool and daemon.

## Packages

| Name                         | Description        | Spec                               | Comment                  | Docs                      |
|:-----------------------------|:-------------------|:-----------------------------------|:-------------------------|:--------------------------|
| [base32i][base32i]           | Base32 Encoding    | [Spec][base32i-spec]               |                          | [godoc][base32i-doc]      |
| [bccsp][bccsp]               | Crypto Service     | [BCCSP][bccsp-spec]                | Source: [Fabric][fabric] | [godoc][bccsp-doc]        |
| [celldag][celldag]           | Cell IPFS DAG      | [IPFS DAG][ipfs-dag]               |                          | [godoc][celldag-doc]      |
| [cells][cells]               | Cells & Codecs     | [IPFN Cells][cell-spec]            |                          | [godoc][cells-doc]        |
| [cmdutil][cmdutil]           | Console Utilities  |                                    |                          | [godoc][cmdutil-doc]      |
| [commands][commands]         | IPFN Commands      |                                    |                          | [godoc][commands-doc]     |
| [cryptoutil][cryptoutil]     | Crypto Utilities   |                                    |                          | [godoc][cryptoutil-doc]   |
| [entropy][entropy]           | Entropy            |                                    |                          | [godoc][entropy-doc]      |
| [flog][flog]                 | Logging utilities  |                                    | Source: [Fabric][fabric] | [godoc][flog-doc]         |
| [hashutil][hashutil]         | Hashing Utilities  |                                    |                          | [godoc][hashutil-doc]     |
| [hdkeychain][hdkeychain]     | Hex Utilities      | [BIP-32][bip32]                    |                          | [godoc][hdkeychain-doc]   |
| [hexutil][hexutil]           | Hex Utilities      |                                    |                          | [godoc][hexutil-doc]      |
| [keypair][keypair]           | Key Pair Utilities |                                    |                          | [godoc][keypair-doc]      |
| [sealbox][sealbox]           | Seal Box           | [Web3 Secrets][web3-secrets]       |                          | [godoc][sealbox-doc]      |
| [shortaddress][shortaddress] | Short address      | [Short address][shortaddress-spec] | Prototype                | [godoc][shortaddress-doc] |

## Dependencies

See Go [dep](https://golang.github.io/dep/) management tool.

## License

                                 Apache License
                           Version 2.0, January 2004
                        http://www.apache.org/licenses/

## Project

This source code is part of [IPFN](https://github.com/ipfn) – interplanetary functions project.

[COPYING]: https://github.com/ipfn/ipfn/blob/master/COPYING.txt
[badge-ci]: https://travis-ci.org/ipfn/ipfn.svg?branch=master
[badge-copying]: https://img.shields.io/badge/license-see%20COPYING.txt-blue.svg?style=flat-square
[badge-docs]: https://img.shields.io/badge/documentation-IPFN-blue.svg?style=flat-square
[badge-godoc]: https://godoc.org/github.com/ipfn/ipfn/src/go?status.svg
[badge-ipfn]: https://img.shields.io/badge/project-IPFN-blue.svg?style=flat-square
[base32i-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/utils/base32i
[base32i-spec]: https://github.com/ipfn/ipfn/blob/master/src/go/utils/base32i/base32i.go#L25
[base32i]: https://github.com/ipfn/ipfn/tree/master/src/go/utils/base32i
[bccsp-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/crypto/bccsp
[bccsp-spec]: https://jira.hyperledger.org/secure/attachment/10124/BCCSP.pdf
[bccsp]: https://godoc.org/github.com/ipfn/ipfn/src/go/crypto/bccsp
[bip32]: https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki
[cell-spec]: https://github.com/ipfn/ipfn/tree/master/src/proto/cell.proto
[celldag-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/cells/celldag
[celldag]: https://godoc.org/github.com/ipfn/ipfn/src/go/cells/celldag
[cells-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/cells
[cells]: https://godoc.org/github.com/ipfn/ipfn/src/go/cells
[ci]: https://travis-ci.org/ipfn/ipfn
[cmdutil-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/utils/cmdutil
[cmdutil]: https://godoc.org/github.com/ipfn/ipfn/src/go/utils/cmdutil
[commands-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/commands
[commands]: https://godoc.org/github.com/ipfn/ipfn/src/go/commands
[cryptoutil-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/utils/cryptoutil
[cryptoutil]: https://godoc.org/github.com/ipfn/ipfn/src/go/utils/cryptoutil
[docs]: https://docs.ipfn.io/
[entropy-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/utils/entropy
[entropy]: https://github.com/ipfn/ipfn/tree/master/src/go/utils/entropy
[fabric]: https://github.com/hyperledger/fabric
[flog-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/utils/flog
[flog]: https://godoc.org/github.com/ipfn/ipfn/src/go/utils/flog
[godoc-ipfn]: https://godoc.org/github.com/ipfn/ipfn/src/go
[hdkeychain-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/utils/hdkeychain
[hdkeychain]: https://github.com/ipfn/ipfn/tree/master/src/go/utils/hdkeychain
[hashutil-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/utils/hashutil
[hashutil]: https://github.com/ipfn/ipfn/tree/master/src/go/utils/hashutil
[hexutil-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/utils/hexutil
[hexutil]: https://github.com/ipfn/ipfn/tree/master/src/go/utils/hexutil
[ipfs-dag]: https://github.com/ipfs/specs/tree/master/merkledag
[keypair-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/keypair
[keypair]: https://godoc.org/github.com/ipfn/ipfn/src/go/keypair
[org-ipfn]: https://github.com/ipfn
[sealbox-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/crypto/sealbox
[sealbox]: https://godoc.org/github.com/ipfn/ipfn/src/go/crypto/sealbox
[shortaddress-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/utils/shortaddress
[shortaddress-spec]: https://github.com/ipfn/ipfn/blob/master/src/go/utils/shortaddress/address.go#L15
[shortaddress]: https://github.com/ipfn/ipfn/tree/master/src/go/utils/shortaddress
[wallet-doc]: https://godoc.org/github.com/ipfn/ipfn/src/go/wallet
[web3-secrets]: https://github.com/ethereum/wiki/wiki/Web3-Secret-Storage-Definition
