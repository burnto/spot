Spot
====

"See spot run"

Spot is a dead simple, feature-free command for Go that watches your filesystem and (re)runs your command when a change is observed. It uses [fsnotify](https://github.com/howeyc/fsnotify).

If you need something beyond this dead simple, single-use reloading offered here, check out other monitor/exec tools that are far more featureful, such as [watchdog](https://github.com/gorakhargosh/watchdog) (Python) or [watchr](https://github.com/bevry/watchr/) (node.js). 

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

License
-------

The MIT License (MIT)
Copyright © 2013 <copyright holders>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
