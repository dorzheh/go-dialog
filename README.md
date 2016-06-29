go-dialog
=========

go-dialog is a Go wrapper for the dialog utility originally written by Savio Lam, and later rewritten by Thomas E. Dickey.

Usage
=========
package main

import (
	"fmt"
	"github.com/weldpua2008/go-dialog"	
)
func main() {
   d := dialog.New(dialog.AUTO, 0)
   d.Msgbox("Hello world!")
}


Installation
=========
```bash
 go get "github.com/weldpua2008/go-dialog"
```

Contributors
=========
[Valeriy Soloviov](github.com/weldpua2008/) <weldpua2008 @gmail.com>
[Dmitry Orzhehovsky](github.com/dorzheh/) <dorzheh @gmail.com> 
Pavel Vershinin <master-dev @inbox.ru>
