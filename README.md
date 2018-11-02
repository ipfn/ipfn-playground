# IPFN – Interplanetary Functions

[![IPFN project][badge-ipfn]][org-ipfn]
[![IPFN Documentation][badge-docs]][docs]
[![Apache-2.0 License][badge-license]][LICENSE]
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
git clone https://github.com/ipfn/ipfn.git
cd ipfn
```

#### Vagrant

To start development using [Vagrant][] run following commands in source directory:

```sh
vagrant up
```

#### Docker

To start development using [Docker][] run following commands in source directory:

```sh
docker pull ipfn/env:latest
docker run -it -v `pwd`:/src ipfn/env:latest
```

### Requirements

* [Go][] >= 1.9.2
* [CMake][] >= 3.5
* [Python][] >= 3.0 (only for Python SDK)
* [node.js][] >= 8.4 (only for node.js SDK)
* [Rust][] >= 1.3 (only for Rust SDK)
* [g++][gcc] >= 5.0 (for [seastar][] server)

Lower versions could also work but are not targeted.

Installed automatically in [Vagrant](./Vagrantfile) and [Docker](./Dockerfile) by [shell scripts](./tools/devenv).

### Platform support

Interplanetary functions project development prioritizes [Linux][] platform over any other.

| Operating System | High-performance | Server    | Client   | CI                           |
|:-----------------|:-----------------|:----------|:---------|:-----------------------------|
| Fedora 29        | &#10003;         | &#10003;  | &#10003; | [![Travis CI][badge-ci]][ci] |
| Fedora 28        | &#10003;         | &#10003;  | &#10003; | [![Travis CI][badge-ci]][ci] |
| Ubuntu 18.10     | &#10003;         | &#10003;  | &#10003; | [![Travis CI][badge-ci]][ci] |
| Ubuntu 18.04     | &#10003;         | &#10003;  | &#10003; | [![Travis CI][badge-ci]][ci] |
| Ubuntu 16.04     | &#10003;         | &#10003;  | &#10003; | [![Travis CI][badge-ci]][ci] |
| Windows 10       | &#10007;         | &#10003;* | &#10003; | &#8230;                      |
| OSX              | &#10007;         | &#10003;* | &#10003; | &#8230;                      |

* \* - supported using [Vagrant][] and [VirtualBox][]

## Technologies

IPFN builds on many technologies and open source work of many different companies altogether, some of which include following:

### General

* [ipfs][] - interplanetary file system
* [linux][] - linux foundation

### Networking

* [smf][] - fastest rpc in the west
* [seastar][] - high performance framework

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
* [tensorflow][] - computation graphs playground
* [emscripten] - llvm to wasm and js compiler
* [assemblysript][] - typescript to wasm compiler

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

This software is licensed under the Apache License 2.0, see [COPYING](./COPYING) file for more details.

## Project

This source code is part of [IPFN](https://github.com/ipfn) – interplanetary functions project.

[ci]: https://travis-ci.org/ipfn/ipfn
[docs]: https://docs.ipfn.io/
[badge-ci]: https://travis-ci.org/ipfn/ipfn.svg?branch=master
[badge-license]: https://dmlc.github.io/img/apache2.svg
[badge-ipfn]: https://img.shields.io/badge/project-IPFN-blue.svg?style=flat-square
[badge-docs]: https://img.shields.io/badge/documentation-IPFN-blue.svg?style=flat-square
[org-ipfn]: https://github.com/ipfn
[COPYING]: https://github.com/ipfn/ipfn/blob/master/COPYING.txt
[LICENSE]: https://github.com/ipfn/ipfn/blob/master/LICENSE.txt
[linux]: https://www.linuxfoundation.org/
[seastar]: https://github.com/scylladb/seastar
[ipfs]: https://github.com/ipfs/go-ipfs/
[smf]: https://senior7515.github.io/smf/
[seastar]: https://github.com/scylladb/seastar/
[npm]: https://www.npmjs.com/
[wasmi]: https://github.com/paritytech/wasmi
[wagon]: https://github.com/go-interpreter/wagon
[assemblysript]: https://github.com/AssemblyScript/assemblyscript
[emscripten]: https://github.com/kripken/emscripten
[tvm]: https://github.com/dmlc/tvm/
[llvm]: https://llvm.org/
[tensorflow]: https://www.tensorflow.org/
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
