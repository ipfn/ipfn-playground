# IPFN – Interplanetary Functions

[![IPFN project](https://img.shields.io/badge/project-IPFN-blue.svg?style=flat-square)](//github.com/ipfn)
[![IPFN Documentation](https://img.shields.io/badge/documentation-IPFN-blue.svg?style=flat-square)](//ipfn.github.io/documentation/)

IPFN – Interplanetary Functions Project.

## Cells

Cells have bodies made up of other cells.
Cell can have a soul which fulfills a purpose.
Cells can stimulate and produce another cells.
Cell can contain a memory.

```capnp
struct Cell {
  name   @0 :Text;
  soul   @1 :Text;
  body   @2 :List(Cell);
  feed   @4 :List(Text);
  memory @3 :Tensor;
}
```

<!--
## Documentation

Documentation for IPFN project is on [ipfn.github.io/documentation](//ipfn.github.io/documentation/).

## Examples

Repositories containing example neurons are hosted on [ipfn-examples](//github.com/ipfn-examples) organization.
-->

## Project

This repository is part of [IPFN](//github.com/ipfn) – interplanetary functions project.
