# Quickstart

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

There is a prebuilt [ipfn/dev][vagrant-box] Vagrant box.

NOTE: This may have parts missing or be outdated.

```sh
vagrant init ipfn/dev
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

[Vagrant]: https://www.vagrantup.com/
[Docker]: https://docker.com
[Go]: https://golang.org/
[vagrant-box]: https://app.vagrantup.com/ipfn/boxes/dev
