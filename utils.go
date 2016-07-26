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

const (
	DIALOG_PACKAGE_AUTO_NOT_FOUND = "Package not found!\nPlease install " + KDE + " or " + GTK + " or " + X + " or " + CONSOLE
)

// helper function that checks path and return error, nil
func getPathOeRaiseError(environment string) error {
	var err error
	switch environment {
	case CONSOLE, KDE, GTK, X:
		_, execution_error := exec.LookPath(environment)
		if execution_error != nil {
			err = fmt.Errorf("Package not found!\nPlease install " + environment)
		}
	case AUTO:
		env := ""
		for _, pkg := range []string{KDE, GTK, X, CONSOLE} {
			_, execution_error := exec.LookPath(pkg)
			if execution_error == nil {
				env = pkg
				break
			}
		}
		if env == "" {
			err = fmt.Errorf(DIALOG_PACKAGE_AUTO_NOT_FOUND)
		}
	case DIALOG_TEST_ENV:
		break
	default:
		err = fmt.Errorf("Unknown package " + environment)
	}
	return err

}

func DialogFindPathOrExit(environment string) {
	err := getPathOeRaiseError(environment)
	if err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}

}
