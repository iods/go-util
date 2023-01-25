Package `boolutil` v0.1.0
=======================

Package `boolutil` adds and extends default tool functions available for booleans in Go.

> **WARNING:** This is a private package and is not currently maintained for public use.

tl;dr
-----

 * does some things.
 * does some more things.
 * does a couple more things.


Getting Started
---------------

### Installation

```shell
go get github.com/iods/go-util/type/boolutil
```

### Built With
* [fatih/color](https://github.com/fatih/color)

### Rationale


Function API
------------

### Flip()

In any action you can flip the result or value of a boolean.

```
func Flip(b bool) bool {}
```

### ToInt()

In any action you can return a true/false argument as an int (1 or 0).

```
func ToInt(b bool) int {}
```


Acknowledgements
----------------

This package and others in this repository, are heavily influenced by the following:

 * https://github.com/gookit/
 * https://github.com/Darevski/