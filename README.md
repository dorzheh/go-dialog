[![Build Status](https://travis-ci.org/weldpua2008/go-dialog.svg?branch=master)](https://travis-ci.org/weldpua2008/go-dialog)
[![Coverage Status](https://coveralls.io/repos/github/weldpua2008/go-dialog/badge.svg?branch=master)](https://coveralls.io/github/weldpua2008/go-dialog?branch=master)

go-dialog
=========

go-dialog is a Go wrapper for the dialog utility originally written by Savio Lam, and later rewritten by Thomas E. Dickey.

Usage
=========
```go
package main

import (
	"fmt"
	"github.com/weldpua2008/go-dialog"	
)
func main() {
   d := dialog.New(dialog.AUTO, 0)
   d.Msgbox("Hello world!")
}
```

Installation
=========
```bash
 go get "github.com/weldpua2008/go-dialog"
```

Contributors
=========
* [Valeriy Soloviov](http://github.com/weldpua2008/) weldpua2008 @gmail.com
* [Dmitry Orzhehovsky](http://github.com/dorzheh/) dorzheh @gmail.com
* Pavel Vershinin master-dev @inbox.ru
