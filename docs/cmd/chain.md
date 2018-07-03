# Chain commands

IPFN chain commands.

## `ipfn chain init` - initializes a new chain

### Usage

```sh
$ ipfn chain init -h
Initializes a new chain.

See wallet usage for more information on key derivation path.

Usage:
  ipfn chain init [config] [flags]

Examples:
  $ ipfn chain init -n mychain -k wallet:1e6:1e6 -k default/x/test:1e6:0
  $ ipfn chain init -p zFNScYMGz4wQocWbvHVqS1HcbzNzJB5JK3eAkzF9krbSLZiV8cNr:1

Flags:
  -h, --help           help for init
  -k, --key  strings   key path and power in key:power:delegated format
  -a, --addr strings   address and power in addr:power format

Global Flags:
  -c, --config string   config file (default "~/.ipfn.json")
  -v, --verbose         verbose logs output (stdout/stderr)
```

### Example

```sh
$ ipfn chain init \
    -k default/x/test:1e6:1e6 \
    -k default/x/test2:1e6:1e6 \
    -k default/x/test3:1e6:0 \
    -p zFNScYMH2PZmRrdEF3aP7HrHVM2HegvKwMa2yFMKeZ2wwQR35EXy:1e6

Wallet "default" password: ***
{
  "header": {
    "timestamp": "2018-07-03T21:00:04.7180038+02:00",
    "head_hash": "zHeadAt95i1ac3bx3yi2Pkpbseh79bDUiNiYHDRWbVdewZNDeaCH",
    "exec_hash": "zFuncm68zVfhy1ZNgLrfCfyjb3mQAaUkw4siQAEtWYEpU4yqgckS",
    "state_hash": "zEsT8raaPksxbjDKo5Nysj8cxMgE4riBzV4WdfLNA5ZAWKQAaxm4",
    "signed_hash": "zFSec2XV4fV81JPkLJDD7SubskdWuW3PBF4oMb8kBSwpDqbNy5N9"
  },
  "exec_ops": [
    "OP_GENESIS",
    "OP_ASSIGN_POWER [ OP_UINT64 1000000 OP_PUBKEY 0x020e8b587eab8b5c9a57f3b6d540f01b2ce154a1c4cddad145234e83aad81282f8 ]",
    "OP_ASSIGN_POWER [ OP_UINT64 1000000 OP_PUBKEY 0x025657544e1355ac629798c62b4a65b917e44ec7b445ee0af334dd5cb5802652ba ]",
    "OP_ASSIGN_POWER [ OP_UINT64 1000000 OP_PUBKEY 0x03ff24488ea4d80627cbf3ff2a6c391a9855a609e3aa3c5e30c501df6fe7075177 ]",
    "OP_ASSIGN_POWER [ OP_UINT64 1000000 OP_CID zFNScYMH2PZmRrdEF3aP7HrHVM2HegvKwMa2yFMKeZ2wwQR35EXy ]",
    "OP_SIGNED [ OP_ASSIGN_POWER [ OP_UINT64 1000000 ] OP_SIGNATURE 0x1b762ac624e0209e280c12e7b9c758fc411d938a381ccc3cc474aea2b14780cea5466e649dc41fdd252791b7aa39d9f1307f2bcc39e9445258d6dc68b9f3d58ae6 ]",
    "OP_SIGNED [ OP_ASSIGN_POWER [ OP_UINT64 1000000 ] OP_SIGNATURE 0x1c8bee3c13d0a1987c29e387387be56f8ff7fc2db93e0fd45fac9b75f0134d7e87684347b960fe9c5df05bd2e26386b6b1d8c65c7468c055e026aa3382d45620c1 ]"
  ],
  "signatures": [
    "HAos0y5M21Qfx8/tGpcs9jDrF2C8+sRFRKcPA+wY4u/MXpAVl1OZ8ewbU3WLOfoN02zbnAk6buFpqDHHmVby4no=",
    "HL5uUc2UXuleym23YAvdhIImkRox8wsYZTv2Ki3E4J94YYdvCvZFYPCE95fn6yBRhqEe+pmceeQGGf+d9NnbsK0="
  ]
}
```