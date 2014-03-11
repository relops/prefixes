Country Prefix Codes For Go
===========================

[![Build Status](https://travis-ci.org/relops/prefixes.png?branch=master)](https://travis-ci.org/relops/prefixes)

Installation
------------

    go get github.com/relops/prefixes


Features
--------

* Looks up country dialing code prefixes using a given number
* Can tell the difference between the US and Canada
* Handles the Vatican City

Example
-------

```go
package main

import (
	"github.com/relops/prefixes"
	"fmt"
)

func main() {
	country, err := prefixes.Lookup("46485562003")
	fmt.Printf("Country: %s\n", country.Name) // Prints Sweden
}
```
