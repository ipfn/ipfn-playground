# IPFN – Interplanetary Functions

[![IPFN project](https://img.shields.io/badge/project-IPFN-blue.svg?style=flat-square)](//github.com/ipfn)
[![GoDoc](https://godoc.org/github.com/ipfn/ipfn?status.svg)](https://godoc.org/github.com/ipfn/ipfn)
[![IPFN Documentation](https://img.shields.io/badge/documentation-IPFN-blue.svg?style=flat-square)](//ipfn.github.io/documentation/)
[![Circle CI](https://img.shields.io/circleci/project/ipfn/ipfn.svg)](https://circleci.com/gh/ipfn/ipfn)

IPFN – Interplanetary Functions Project.

## Cells

* Cells have bodies made up of other cells.
* Cell can have a soul which fulfills a purpose.
* Cells can stimulate and produce another cells.
* Cell can contain a memory.

### Cell structure

```capnp
interface Cell {
    name?: string;
    soul?: string;
    body?: Cell[];
    memory?: any;
}
```

<!--
## Documentation

Documentation for IPFN project is on [ipfn.github.io/documentation](//ipfn.github.io/documentation/).

## Examples

Repositories containing example neurons are hosted on [ipfn-examples](//github.com/ipfn-examples) organization.
-->

## Project

This source code is part of [IPFN](//github.com/ipfn) – interplanetary functions project.
