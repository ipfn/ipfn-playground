# JavaScript synaptic types definition table

[![IPFN project](https://img.shields.io/badge/project-IPFN-blue.svg?style=flat-square)](http://github.com/ipfn)
[![npm](https://img.shields.io/npm/v/synaptic-types.svg?maxAge=8640&style=flat-square)](https://www.npmjs.com/package/synaptic-types)

This package provides JavaScript definition table of IPFN **synaptic types**.

## Synaptic types

Currently this relies heavily on protocol buffers and it's clearly work-in-progress.

| Name      | Type       | Done |
|-----------|------------|------|
| bool      | boolean    | ✓    |
| number    | number     | ✓    |
| string    | string     | ✓    |
| float     | Number     | ✗    |
| double    | Number     | ✗    |
| sfixed32  | Number     | ✗    |
| fixed32   | Number     | ✗    |
| int32     | Number     | ✗    |
| int64     | Number     | ✗    |
| uint32    | Number     | ✗    |
| uint64    | Number     | ✗    |
| uint128   | BigNumber  | ✗    |
| uint256   | BigNumber  | ✗    |
| sint64    | Number     | ✗    |
| sint32    | Number     | ✗    |
| varint    | Number     | ✗    |

## Install

To include in node.js project in its directory run:

```console
$ npm install --save synaptic-types
```

## Notes

Possibilites in the future include `enum`.

## References

* [Protocol Buffers](https://developers.google.com/protocol-buffers/)

## License

                                 Apache License
                           Version 2.0, January 2004
                        http://www.apache.org/licenses/
