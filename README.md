[![Build Status](https://travis-ci.org/alexanderGugel/resolve.svg)](https://travis-ci.org/alexanderGugel/resolve)

resolve
=======

`require(X)`` from module at path `Y`.

Implements Node's [`require.resolve` algorithm](http://nodejs.org/docs/v0.4.8/api/all.html#all_Together...) in Go.

Allows you to resolve `require` statements just like [Node](https://nodejs.org/) or [browserify](https://github.com/substack/node-browserify).

Usage
-----

```go
helloDependency, _ := Resolve("./hello.js", pwd)
log.Println(helloDependency.Pathname)
```

Works with `node_modules`, directories and files. Correctly resolves main files as specified in `package.json`, e.g. `hello` could have also been resolved to `....../my_modules/node_modules/hello/some-file-as-specified-in-package-json.js`

Credits
-------

* [substack/node-resolve](https://github.com/substack/node-resolve)
