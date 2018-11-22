# IPFN implementation in Go

[![IPFN project][badge-ipfn]][org-ipfn]
[![IPFN Documentation][badge-docs]][docs]
[![See COPYING.txt][badge-copying]][COPYING]
[![GoDoc][badge-godoc]][godoc-ipfn]
[![Travis CI][badge-ci]][ci]

Go implementation of IPFN core with command line tool and daemon.

## Packages

| Name               | Description    | Spec                         | Comment                  | Docs                 |
|:-------------------|:---------------|:-----------------------------|:-------------------------|:---------------------|
| [bccsp][bccsp]     | Crypto Service | [BCCSP][bccsp-spec]          | Source: [Fabric][fabric] | [godoc][bccsp-doc]   |
| [celldag][celldag] | Cell IPFS DAG  | [IPFS DAG][ipfs-dag]         | Prototype                | [godoc][celldag-doc] |
| [cells][cells]     | Cells & Codecs | [IPFN Cell V1][cell-spec]    |                          | [godoc][cells-doc]   |
| [codec][codec]     | IPFN Codec     |                              |                          | [godoc][codec-doc]   |
| [digest][digest]   | Multihashing   | [multihash][multihash]       |                          | [godoc][digest-doc]  |
| [sealbox][sealbox] | Seal Box       | [Web3 Secrets][web3-secrets] |                          | [godoc][sealbox-doc] |

## Common

| Name                         | Description        | Spec                               | Comment   | Docs                      |
|:-----------------------------|:-------------------|:-----------------------------------|:----------|:--------------------------|
| [base32i][base32i]           | Base32 Encoding    | [Spec][base32i-spec]               | Prototype | [godoc][base32i-doc]      |
| [codecs][codecs]             | IPFN Codec IDs     |                                    |           | [godoc][codecs-doc]       |
| [pkcid][pkcid]               | Public Key CID     | [Content ID][cid-spec]             | Prototype | [godoc][pkcid-doc]        |
| [shortaddress][shortaddress] | IPFN Short Address | [Short address][shortaddress-spec] | Prototype | [godoc][shortaddress-doc] |

## Utilities

| Name                     | Description       | Spec            | Comment                  | Docs                    |
|:-------------------------|:------------------|:----------------|:-------------------------|:------------------------|
| [byteutil][byteutil]     | Byte Utilities    |                 |                          | [godoc][byteutil-doc]   |
| [cmdutil][cmdutil]       | Console Utilities |                 |                          | [godoc][cmdutil-doc]    |
| [entropy][entropy]       | Entropy           |                 |                          | [godoc][entropy-doc]    |
| [flog][flog]             | Logging Utilities |                 | Source: [Fabric][fabric] | [godoc][flog-doc]       |
| [hexutil][hexutil]       | Hex Utilities     |                 |                          | [godoc][hexutil-doc]    |
| [hdkeychain][hdkeychain] | HD-Key Derivation | [BIP-32][bip32] |                          | [godoc][hdkeychain-doc] |

## Application

| Name                 | Description   | Spec | Comment | Docs                  |
|:---------------------|:--------------|:-----|:--------|:----------------------|
| [commands][commands] | IPFN Commands |      |         | [godoc][commands-doc] |

## Dependencies

See Go [dep](https://golang.github.io/dep/) management tool.

## License

See [COPYING][COPYING] file for licensing details.

## Project

This source code is part of [IPFN](https://github.com/ipfn) – interplanetary functions project.

[COPYING]: https://github.com/ipfn/ipfn/blob/master/COPYING.txt
[badge-ci]: https://travis-ci.org/ipfn/ipfn.svg?branch=master
[badge-copying]: https://img.shields.io/badge/license-see%20COPYING.txt-blue.svg?style=flat-square
[badge-docs]: https://img.shields.io/badge/documentation-IPFN-blue.svg?style=flat-square
[badge-godoc]: https://godoc.org/github.com/ipfn/ipfn/pkg?status.svg
[badge-ipfn]: https://img.shields.io/badge/project-IPFN-blue.svg?style=flat-square
[base32i-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/common/base32i
[base32i-spec]: https://github.com/ipfn/ipfn/blob/master/pkg/common/base32i/base32i.go#L25
[base32i]: https://github.com/ipfn/ipfn/tree/master/pkg/common/base32i
[codec]: https://github.com/ipfn/ipfn/tree/master/pkg/codec
[codec-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/codec
[codecs]: https://github.com/ipfn/ipfn/tree/master/pkg/common/codecs
[codecs-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/common/codecs
[digest]: https://github.com/ipfn/ipfn/tree/master/pkg/digest
[digest-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/digest
[bccsp-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/crypto/bccsp
[bccsp-spec]: https://jira.hyperledger.org/secure/attachment/10124/BCCSP.pdf
[bccsp]: https://github.com/ipfn/ipfn/tree/master/pkg/crypto/bccsp
[bip32]: https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki
[cell-spec]: https://github.com/ipfn/ipfn/tree/master/proto/cell.proto
[celldag-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/cells/celldag
[celldag]: https://github.com/ipfn/ipfn/tree/master/pkg/cells/celldag
[cells-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/cells
[cells]: https://github.com/ipfn/ipfn/tree/master/pkg/cells
[ci]: https://travis-ci.org/ipfn/ipfn
[cmdutil-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/utils/cmdutil
[cmdutil]: https://github.com/ipfn/ipfn/tree/master/pkg/utils/cmdutil
[commands-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/commands
[commands]: https://godoc.org/github.com/ipfn/ipfn/tree/master/pkg/commands
[docs]: https://docs.ipfn.io/
[entropy-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/crypto/entropy
[entropy]: https://github.com/ipfn/ipfn/tree/master/pkg/crypto/entropy
[fabric]: https://github.com/hyperledger/fabric
[flog-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/utils/flog
[flog]: https://github.com/ipfn/ipfn/tree/master/pkg/utils/flog
[godoc-ipfn]: https://godoc.org/github.com/ipfn/ipfn/tree/master/pkg
[hdkeychain-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/utils/hdkeychain
[hdkeychain]: https://github.com/ipfn/ipfn/tree/master/pkg/utils/hdkeychain
[hexutil-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/utils/hexutil
[hexutil]: https://github.com/ipfn/ipfn/tree/master/pkg/utils/hexutil
[byteutil-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/utils/byteutil
[byteutil]: https://github.com/ipfn/ipfn/tree/master/pkg/utils/byteutil
[ipfs-dag]: https://github.com/ipfs/specs/tree/master/merkledag
[pkcid-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/common/pkcid
[pkcid]: https://github.com/ipfn/ipfn/tree/master/pkg/common/pkcid
[org-ipfn]: https://github.com/ipfn
[sealbox-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/crypto/sealbox
[sealbox]: https://godoc.org/github.com/ipfn/ipfn/tree/master/pkg/crypto/sealbox
[shortaddress-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/common/shortaddress
[shortaddress-spec]: https://github.com/ipfn/ipfn/blob/master/pkg/common/shortaddress/address.go#L15
[shortaddress]: https://github.com/ipfn/ipfn/tree/master/pkg/common/shortaddress
[wallet-doc]: https://godoc.org/github.com/ipfn/ipfn/pkg/wallet
[web3-secrets]: https://github.com/ethereum/wiki/wiki/Web3-Secret-Storage-Definition
[cid-spec]: https://github.com/ipld/cid
[multihash]: https://multiformats.io/multihash/
