# mvisonneau/go-ebsnvme

[![GoDoc](https://godoc.org/github.com/mvisonneau/go-ebsnvme?status.svg)](https://godoc.org/github.com/mvisonneau/go-ebsnvme/app)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvisonneau/go-ebsnvme)](https://goreportcard.com/report/github.com/mvisonneau/go-ebsnvme)
[![Build Status](https://cloud.drone.io/api/badges/mvisonneau/go-ebsnvme/status.svg)](https://cloud.drone.io/mvisonneau/go-ebsnvme)
[![Coverage Status](https://coveralls.io/repos/github/mvisonneau/go-ebsnvme/badge.svg?branch=master)](https://coveralls.io/github/mvisonneau/go-ebsnvme?branch=master)

`go-ebsnvme` is a [golang](https://golang.org/) version of the [AWS ebsnvme-id python script](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nvme-ebs-volumes.html)

## TL:DR

```bash
~$ go-ebsvnme /dev/nvme0n1
vol-99cff4881d00c56a8
/dev/sda1

~$ go-ebsvnme --volume-id /dev/nvme1n1
vol-80dfffbbee880a72c

~$ go-ebsvnme --device-name /dev/nvme1n1
/dev/xvdf
```

## Install

Have a look onto the [latest release page](https://github.com/mvisonneau/go-ebsnvme/releases/latest) and pick your flavor.

### Go

```bash
~$ go get -u github.com/mvisonneau/go-ebsnvme
```

### Docker

```bash
~$ docker run -it --rm mvisonneau/go-ebsnvme
```

### Binaries, DEB and RPM packages

For the following ones, you need to know which version you want to install, to fetch the latest available :

```bash
~$ export VERSION=$(curl -s "https://api.github.com/repos/mvisonneau/go-ebsnvme/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
```

```bash
# Binary (eg: linux/amd64)
~$ wget https://github.com/mvisonneau/go-ebsnvme/releases/download/${VERSION}/go-ebsnvme_${VERSION}_linux_amd64.tar.gz
~$ tar zxvf go-ebsnvme_${VERSION}_linux_amd64.tar.gz -C /usr/local/bin

# DEB package (eg: linux/386)
~$ wget https://github.com/mvisonneau/go-ebsnvme/releases/download/${VERSION}/go-ebsnvme_${VERSION}_linux_386.deb
~$ dpkg -i go-ebsnvme_${VERSION}_linux_386.deb

# RPM package (eg: linux/arm64)
~$ wget https://github.com/mvisonneau/go-ebsnvme/releases/download/${VERSION}/go-ebsnvme_${VERSION}_linux_arm64.rpm
~$ rpm -ivh go-ebsnvme_${VERSION}_linux_arm64.rpm
```

## Usage

### Library

```bash
go import github.com/mvisonneau/go-ebsnvme/pkg
```

### Client

```bash
~$ go-ebsnvme -h
NAME:
   go-ebsnvme - Fetch information about AWS EBS NVMe volumes

USAGE:
   go-ebsnvme <block_device> [--volume-id|--device-name]

VERSION:
   <devel>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --volume-id, -i    only print the EBS volume-id
   --device-name, -n  only print the name of the block device
   --help, -h         show help
   --version, -v      print the version
```

## Develop / Test

If you use docker, you can easily get started using :

```bash
~$ make dev-env
# You should then be able to use go commands to work onto the project, eg:
~docker$ make install
~docker$ go-ebsnvme
```

## Contribute

Contributions are more than welcome! Feel free to submit a [PR](https://github.com/mvisonneau/go-ebsnvme/pulls).
