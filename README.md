# IPFN – Interplanetary Functions

[![IPFN project][badge-ipfn]][org-ipfn]
[![IPFN Documentation][badge-docs]][docs]
[![Apache-2.0 License][badge-copying]][COPYING]
[![GoDoc][badge-godoc]][godoc-ipfn]
[![Travis CI][badge-ci]][ci]

IPFN – Interplanetary Functions Project Development Environment

NOTE: at this point it's empty development environment.

## Quickstart

Installation of IPFN client requires only [Go][] compiler installed.

```sh
go get -u github.com/ipfn/ipfn
```

In the future relases will be available to download for all major platforms.

### Development

If you want to start IPFN core development you can give it a quick try with [Vagrant][] or [Docker][]:

#### Clone

```sh
git clone --recursive https://github.com/ipfn/ipfn.git && cd ipfn
```

#### Vagrant

To start development using [Vagrant][] run following commands in source directory:

```sh
vagrant up
```

#### Docker

To start development using [Docker][] you can download latest environment image:

```sh
docker pull ipfn/env:latest
```

Start in IPFN source code directory or replace the `$(pwd)` with it's location:

```sh
docker run -it -v $(pwd):/src ipfn/env:latest
```

This is only required to work on IPFN and not to use it.

### Requirements

* [Go][] >= 1.9.2
* [CMake][] >= 3.5
* [g++][gcc] >= 5.0 (for [seastar][] server)
* [Python][] >= 3.0 (only for Python SDK)
* [node.js][] >= 8.4 (only for node.js SDK)
* [Rust][] nightly (only for Rust SDK)

Lower versions could also work but are not targeted.

Installed automatically in [Vagrant](./Vagrantfile) and [Docker](./Dockerfile) by [shell scripts](./tools/devenv).

### Platform support

Interplanetary functions project development prioritizes [Linux][] platform over any other.

| Operating System | High-performance | Server    | Client   | CI                           |
|:-----------------|:-----------------|:----------|:---------|:-----------------------------|
| Debian 6+        | &#10003;         | &#10003;  | &#10003; | [![Travis CI][badge-ci]][ci] |
| Fedora 21+       | &#10003;         | &#10003;  | &#10003; | [![Travis CI][badge-ci]][ci] |
| Ubuntu 14.04+    | &#10003;         | &#10003;  | &#10003; | [![Travis CI][badge-ci]][ci] |
| OSX              | &#10007;         | &#10003;* | &#10003; | &#8230;                      |
| Windows          | &#10007;         | &#10003;* | &#10003; | &#8230;                      |

* \* - supported using [Vagrant][] and [VirtualBox][]

It should be possible to build on any other host using [Docker][].

## Technologies

IPFN builds on many technologies and open source work of many different companies altogether, some of which include following:

### General

* [ipfs][] - interplanetary file system
* [linux][] - linux foundation

### Networking

* [smf][] - fastest rpc in the west
* [seastar][] - high performance framework

### Storage

* [scylla][] - high perfomance data store

### Sandboxing

* [runc][] - linux containers
* [runv][] - virtual machine containers
* [kata-runtime][] - virtual machine containers
* [docker][] - container build system
* [containerd][] - container management system
* [qemu][] - full system emulator

### Compilation

* [tvm][] - deep learning compiler stack
* [llvm][] - compiler infrastructure
* [halide][] - language for portable computation
* [onnx][] - interchangeable graph format
* [tensorflow][] - computation graphs playground
* [emscripten] - llvm to wasm and js compiler
* [assemblysript][] - typescript to wasm compiler

### Protocols

* [protobuf][] - serialization format
* [flatbuffers][] - non-serialization format

### Web Runtime

* [npm][] - javascript package manager
* [wasmi][] - wasm interpreter in rust
* [wagon][] - wasm interpreter in go

### Blockchain

* [bitcoin][] - cryptography toolchains
* [ethereum][] - key management infrastructure
* [fabric][] - key management infrastructure
* [parity][] - secure blockchain infrastructure

## License

Contributors licensed under the Apache License 2.0.
See [COPYING](./COPYING) file for licensing details.

## Project

This source code is part of [IPFN](https://github.com/ipfn) – interplanetary functions project.

[ci]: https://travis-ci.org/ipfn/ipfn
[docs]: https://docs.ipfn.io/
[badge-ci]: https://travis-ci.org/ipfn/ipfn.svg?branch=master
[badge-copying]: https://img.shields.io/badge/license-see%20COPYING.txt-blue.svg?style=flat-square
[badge-ipfn]: https://img.shields.io/badge/project-IPFN-blue.svg?style=flat-square
[badge-docs]: https://img.shields.io/badge/documentation-IPFN-blue.svg?style=flat-square
[badge-godoc]: https://godoc.org/github.com/ipfn/ipfn/go?status.svg
[godoc-ipfn]: https://godoc.org/github.com/ipfn/ipfn/go
[org-ipfn]: https://github.com/ipfn
[COPYING]: https://github.com/ipfn/ipfn/blob/master/COPYING.txt
[linux]: https://www.linuxfoundation.org/
[seastar]: https://github.com/scylladb/seastar
[ipfs]: https://github.com/ipfs/go-ipfs/
[smf]: https://github.com/smfrpc/smf
[seastar]: https://github.com/scylladb/seastar
[scylla]: https://github.com/scylladb/scylla
[npm]: https://www.npmjs.com/
[wasmi]: https://github.com/paritytech/wasmi
[wagon]: https://github.com/go-interpreter/wagon
[assemblysript]: https://github.com/AssemblyScript/assemblyscript
[emscripten]: https://github.com/kripken/emscripten
[tvm]: https://github.com/dmlc/tvm/
[llvm]: https://llvm.org/
[protobuf]: https://github.com/protocolbuffers/protobuf
[flatbuffers]: https://github.com/google/flatbuffers
[halide]: https://github.com/halide/Halide
[tensorflow]: https://www.tensorflow.org/
[onnx]: https://onnx.ai
[runc]: https://github.com/opencontainers/runc
[runv]: https://github.com/hyperhq/runv
[kata-runtime]: https://github.com/kata-containers/runtime
[docker]: https://github.com/docker/docker-ce
[containerd]: https://github.com/containerd/containerd
[qemu]: https://www.qemu.org/
[bitcoin]: https://github.com/btcsuite
[ethereum]: https://github.com/ethereum
[fabric]: https://github.com/hyperledger/fabric
[parity]: https://github.com/paritytech
[Vagrant]: https://www.vagrantup.com/
[VirtualBox]: https://www.virtualbox.org/
[Go]: https://golang.org/
[node.js]: https://nodejs.org
[Rust]: https://www.rust-lang.org/en-US/
[Python]: https://www.python.org/
[CMake]: https://cmake.org/
[gcc]: https://www.gnu.org/software/gcc/
[Fedora]: https://getfedora.org/
[Ubuntu]: https://www.ubuntu.com/
