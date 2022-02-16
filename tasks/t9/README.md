# goget

A very simple `wget` clone written in Go. It supports concurrent
downloads over `http`.

## System Information

Built and tested on a 64-bit Ubuntu 18.04.1 LTS system with Golang
version 1.10.1.

## Installation

The following will clone and install `goget` in `$GOPATH/bin`.
Make sure that `$GOPATH` is defined and that `$GOPATH/bin` is in
your `$PATH`.

```
$ go get github.com/chatzikalymnios/goget/cmd/goget
```

## Usage

| Option        | Description                                                                                |
|---------------|--------------------------------------------------------------------------------------------|
| `-c <int>`    | Number of concurrent downloads allowed. Default value is `GOMAXPROCS` (`nproc`, if unset). |
| `-d <string>` | Directory to save downloaded files to. Default value is the current working directory.     |
| `-h`          | Show help message.                                                                         |
| `-i <string>` | Input file to read URLs from.                                                              |

### Examples

The following will download the two specified files one by one to the
current working directory.

```
$ goget -c 1 \
    http://speedtest.xs4all.net/files/100MiB.bin \
    http://speedtest.xs4all.net/files/50MB.bin
```

The following will download the four files specified in `files.txt`
concurrently.

```
$ cat files.txt
http://speedtest.xs4all.net/files/100MiB.bin
http://speedtest.xs4all.net/files/50MB.bin
http://speedtest.xs4all.net/files/250MB.bin
http://speedtest.xs4all.net/files/100MB.bin

$ goget -c 4 -i files.txt
```

## Licence

MIT license. For more information, see the LICENSE file.
