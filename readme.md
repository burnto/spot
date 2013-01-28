Spot
====

Spot sees your files and runs your commands.

Uses [fsnotify](https://github.com/howeyc/fsnotify).

Install
-------

```bash
go get github.com/burnto/spot
```

Usage
-----

```bash
spot <command>
```

Any file creations, deletions, or modifications within the working directory will trigger a restart of the process. Example:

```bash
spot go run my_serve.go
```
