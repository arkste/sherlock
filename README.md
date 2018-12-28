# sherlock

[![Go Report Card](https://goreportcard.com/badge/github.com/arkste/sherlock)](https://goreportcard.com/report/github.com/arkste/sherlock)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/arkste/sherlock/master/LICENSE)
[![Golang](https://img.shields.io/badge/Go-1.11-blue.svg)](https://golang.org)
![Linux](https://img.shields.io/badge/Supports-Linux-green.svg)
![macOS](https://img.shields.io/badge/Supports-macOS-green.svg)
![windows](https://img.shields.io/badge/Supports-windows-green.svg)

Find usernames across multiple social networks.

This is a multi-threaded Golang-Implementation of [sdushantha/sherlock](https://github.com/sdushantha/sherlock).

[![asciicast](https://asciinema.org/a/YOG0MX8VaaavjU4t8qSlwhVmk.svg)](https://asciinema.org/a/YOG0MX8VaaavjU4t8qSlwhVmk)

## Installation

+  Install Go (1.9+) and set your [GOPATH](https://golang.org/doc/code.html#GOPATH)
+ `go get -u github.com/arkste/sherlock`
+ cd `$GOPATH/src/github.com/arkste/sherlock`
+ `make`

or use Docker:

+ `docker pull arkste/sherlock`

or get a pre-compiled binary from [Releases](https://github.com/arkste/sherlock/releases).

## Usage

    $ sherlock

or 

    $ sherlock --username user123

## License

sherlock is released under the [MIT License](http://www.opensource.org/licenses/MIT).
