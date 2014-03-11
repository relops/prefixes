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

Broken Examples
---------------

We're keen to hear of any examples that don't work - in this case, please raise an issue.

Credits
-------

For providing the input data, kudos goes to:

* https://github.com/mledoze/countries
* http://en.wikipedia.org/wiki/List_of_North_American_Numbering_Plan_area_codes

Please check out the terms under which you can use their data.

License
-------

The MIT License (MIT)

Copyright (c) [2014] [RelOps Ltd]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
