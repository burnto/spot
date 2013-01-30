Spot
====

Spot is a stupidly simple tool for watching your filesystem and (re)running your command when a change is observed. If the first argument is a .go file, it will build before executing.

Check out other monitor/exec tools that are far more featureful:

* [watchdog](https://github.com/gorakhargosh/watchdog) (Python)
* [watchr](https://github.com/bevry/watchr/) (node.js)
* [devsrvr](https://bitbucket.org/liamstask/devsrvr) (Go)

Spot uses [fsnotify](https://github.com/howeyc/fsnotify).

Install
-------

```bash
go get github.com/burnto/spot
```

Usage
-----

You can supply either a go file or an executable, followed by any number of arguments.

```bash
spot <program.go> [args]...
spot <executable> [args]...
```

If the file is a go file, it will be built and then run.

Any file creations, deletions, or modifications within the working directory will trigger a restart of the process. Example:

```bash
spot my_serve.go
```

License
-------

The MIT License (MIT)
Copyright © 2013 <copyright holders>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
