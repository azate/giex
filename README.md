# Giex

Perhaps the fastest publicly available tool for identifying websites with publicly available ```.git``` repositories.

## Usage

Download the **giex-*** file for your platform in the latest available [release](https://github.com/azate/giex/releases)

```
$ ./giex-* -h

Usage:
  giex [flags]

Flags:
  -h, --help               help for giex
  -i, --input string       Path to the file with domains <one row one domain> (default "domains.txt")
  -t, --max-tasks uint     Maximum prepared tasks (default 200)
  -w, --max-workers uint   Maximum workers (default 100)
  -o, --output string      Path to the folder for saving the git configs (default "/tmp")
  -p, --proxy string       HTTP proxy
```
