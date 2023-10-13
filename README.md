# go-ebsnvme

[![PkgGoDev](https://pkg.go.dev/badge/github.com/mvisonneau/go-ebsnvme)](https://pkg.go.dev/mod/github.com/mvisonneau/go-ebsnvme)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvisonneau/go-ebsnvme)](https://goreportcard.com/report/github.com/mvisonneau/go-ebsnvme)
[![test](https://github.com/mvisonneau/go-ebsnvme/actions/workflows/test.yml/badge.svg)](https://github.com/mvisonneau/go-ebsnvme/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/github/mvisonneau/go-ebsnvme/badge.svg?branch=main)](https://coveralls.io/github/mvisonneau/go-ebsnvme?branch=main)
[![release](https://github.com/mvisonneau/go-ebsnvme/actions/workflows/release.yml/badge.svg)](https://github.com/mvisonneau/go-ebsnvme/actions/workflows/release.yml)

`go-ebsnvme` is a [golang](https://golang.org/) version of the [AWS ebsnvme-id python script](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nvme-ebs-volumes.html)

## TL:DR

```bash
~$ go-ebsnvme /dev/nvme0n1
sda1
/dev/sda1
vol-99cff4881d00c56a8

~$ go-ebsnvme -f volume-id /dev/nvme1n1
vol-80dfffbbee880a72c

~$ go-ebsnvme -f device-name /dev/nvme1n1
xvdf

~$ go-ebsnvme -t json -f device-path,volume-id  /dev/nvme1n1
{"device-path":"/dev/xvdf","volume-id":"vol-80dfffbbee880a72c"}
```

## Install

Have a look onto the [latest release page](https://github.com/mvisonneau/go-ebsnvme/releases/latest) and pick your flavor.

### Go

```bash
~$ go run github.com/mvisonneau/go-ebsnvme/cmd/go-ebsnvme@latest

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

```go
package main

import (
   "fmt"
   "os"

   "github.com/mvisonneau/go-ebsnvme/pkg/ebsnvme"
)

func main() {
   device, err := ebsnvme.ScanDevice("/dev/nvme0n1")
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   
   fmt.Println(device.Name)
   fmt.Println(device.Path)
   fmt.Println(device.VolumeID)
}
```

### Client

```bash
~$ go-ebsnvme -h
NAME:
   go-ebsnvme - Find details about currently attached AWS EBS NVMe volumes

USAGE:
   go-ebsnvme [opts] <block_device>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --output-type value, -t value                                      print results in whether "text" or "json" (default: "text")
   --output-field value, -f value [ --output-field value, -f value ]  filter out printed fields (default: "device-name", "device-path", "volume-id")
   --help, -h                                                         show help
```

## Contribute

Contributions are more than welcome! Feel free to submit a [PR](https://github.com/mvisonneau/go-ebsnvme/pulls).
