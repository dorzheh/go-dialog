package dialog

// Copyright 2016 Valeriy Solovyov <weldpua2008@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style

import (
	// "bytes"
	"fmt"
	"os"
	"os/exec"
	// "strconv"
	// "strings"
	// "log"
	// "time"
)

func DialogFindPathOrExit(environment string) {
	var err error
	// fmt.Println(environment)
	switch environment {
	case CONSOLE, KDE, GTK, X:
		_, err = exec.LookPath(environment)
		if err != nil {
			fmt.Println("Package not found!\nPlease install " + environment)
			os.Exit(1)
		}
	case AUTO:
		env := ""
		for _, pkg := range []string{KDE, GTK, X, CONSOLE} {
			_, err = exec.LookPath(pkg)
			if err == nil {
				env = pkg
				break
			}
		}
		if env == "" {
			fmt.Println("Package not found!\nPlease install " + KDE + " or " + GTK + " or " + X + " or " + CONSOLE)
			os.Exit(1)
		}
	case DIALOG_TEST_ENV:
		break
	default:
		fmt.Println("Unknown package " + environment)
		os.Exit(126)
	}
}
