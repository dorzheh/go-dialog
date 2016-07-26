package dialog

// Copyright 2016 Valeriy Solovyov <weldpua2008@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style

import (
	// 	// "bytes"
	// "fmt"
	"os"
	"os/exec"
	"testing"
)

// func RunDialogFindPathOrExit(t *testing.T) {
// 	// if  == "1" {
// 	// DialogFindPathOrExit(os.Getenv("BE_CRASHER"))
// 	// return
// 	// }
// 	if os.Getenv("BE_CRASHER") != "" {
// 		DialogFindPathOrExit(os.Getenv("BE_CRASHER"))
// 		return
// 	}

// }

// func TestDialogFindPathOrExit(t *testing.T) {
// 	fixtures := []struct {
// 		Env string
// 	}{
// 		{Env: CONSOLE},
// 	}

// 	if len(fixtures) < 1 {
// 		t.Fatalf("Failed because you should have test cases")
// 	}

// 	for _, tt := range fixtures {

// 		os.Setenv("PATH", "")
// 		os.Setenv("BE_CRASHER", "")
// 		cmd := exec.Command(os.Args[0], "-test.run=RunDialogFindPathOrExit")
// 		// t.Fatalf(cmd.Path)
// 		// cmd.Path = ""
// 		env_str := "BE_CRASHER=" + tt.Env
// 		cmd.Env = append(os.Environ(), env_str)
// 		err := cmd.Run()
// 		// bs, err := cmd.CombinedOutput()
// 		// t.Log(bs)
// 		// t.Log(err)
// 		e, ok := err.(*exec.ExitError)

// 		if ok && !e.Success() {
// 			t.Log(e.String())
// 			return
// 		}

// 		t.Fatalf("process ran with err %v , want exit status 1", err)
// 	}
// }

func TestDialogFindPathOrExitConsole(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "TestDialogFindPathOrExitConsole" {
		DialogFindPathOrExit(CONSOLE)
		return
	}
	os.Setenv("PATH", "")
	cmd := exec.Command(os.Args[0], "-test.run=TestDialogFindPathOrExitConsole")
	// t.Fatalf(cmd.Path)
	// cmd.Path = ""
	cmd.Env = append(os.Environ(), "BE_CRASHER=TestDialogFindPathOrExitConsole")
	err := cmd.Run()
	// bs, err := cmd.CombinedOutput()
	// t.Log(bs)
	// t.Log(err)
	e, ok := err.(*exec.ExitError)

	if ok && !e.Success() {
		if e.String() != "exit status 1" {
			t.Fatalf(" %v , want exit status 1", e.String())
		}
		return
	}

	t.Fatalf("process ran with err %v , want exit status 1", err)

}

func TestDialogFindPathOrExitKDE(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "TestDialogFindPathOrExitKDE" {
		DialogFindPathOrExit(KDE)
		return
	}
	os.Setenv("PATH", "")
	cmd := exec.Command(os.Args[0], "-test.run=TestDialogFindPathOrExitKDE")
	// t.Fatalf(cmd.Path)
	// cmd.Path = ""
	cmd.Env = append(os.Environ(), "BE_CRASHER=TestDialogFindPathOrExitKDE")
	err := cmd.Run()
	// bs, err := cmd.CombinedOutput()
	// t.Log(bs)
	// t.Log(err)
	e, ok := err.(*exec.ExitError)

	if ok && !e.Success() {
		if e.String() != "exit status 1" {
			t.Fatalf(" %v , want exit status 1", e.String())
		}
		return
	}

	t.Fatalf("process ran with err %v , want exit status 1", err)

}

func TestDialogFindPathOrExitGTK(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "TestDialogFindPathOrExitGTK" {
		DialogFindPathOrExit(GTK)
		return
	}
	os.Setenv("PATH", "")
	cmd := exec.Command(os.Args[0], "-test.run=TestDialogFindPathOrExitGTK")
	// t.Fatalf(cmd.Path)
	// cmd.Path = ""
	cmd.Env = append(os.Environ(), "BE_CRASHER=TestDialogFindPathOrExitGTK")
	err := cmd.Run()
	// bs, err := cmd.CombinedOutput()
	// t.Log(bs)
	// t.Log(err)
	e, ok := err.(*exec.ExitError)

	if ok && !e.Success() {
		if e.String() != "exit status 1" {
			t.Fatalf(" %v , want exit status 1", e.String())
		}
		return
	}

	t.Fatalf("process ran with err %v , want exit status 1", err)

}

func TestDialogFindPathOrExitX(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "TestDialogFindPathOrExitX" {
		DialogFindPathOrExit(X)
		return
	}
	os.Setenv("PATH", "")
	cmd := exec.Command(os.Args[0], "-test.run=TestDialogFindPathOrExitX")
	// t.Fatalf(cmd.Path)
	// cmd.Path = ""
	cmd.Env = append(os.Environ(), "BE_CRASHER=TestDialogFindPathOrExitX")
	err := cmd.Run()
	// bs, err := cmd.CombinedOutput()
	// t.Log(bs)
	// t.Log(err)
	e, ok := err.(*exec.ExitError)

	if ok && !e.Success() {
		if e.String() != "exit status 1" {
			t.Fatalf(" %v , want exit status 1", e.String())
		}
		return
	}
	t.Fatalf("process ran with err %v , want exit status 1", err)

}

func TestDialogFindPathOrExitAUTO(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "TestDialogFindPathOrExitAUTO" {
		DialogFindPathOrExit(AUTO)
		return
	}
	os.Setenv("PATH", "")
	cmd := exec.Command(os.Args[0], "-test.run=TestDialogFindPathOrExitAUTO")
	// t.Fatalf(cmd.Path)
	// cmd.Path = ""
	cmd.Env = append(os.Environ(), "BE_CRASHER=TestDialogFindPathOrExitAUTO")
	err := cmd.Run()
	// bs, err := cmd.CombinedOutput()
	// t.Log(bs)
	// t.Log(err)
	e, ok := err.(*exec.ExitError)

	if ok && !e.Success() {
		if e.String() != "exit status 1" {
			t.Fatalf(" %v , want exit status 1", e.String())
		}
		return
	}
	t.Fatalf("process ran with err %v , want exit status 1", err)
}

func TestDialogFindPathOrExitDIALOG_TEST_ENV(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "TestDialogFindPathOrExitDIALOG_TEST_ENV" {
		DialogFindPathOrExit(DIALOG_TEST_ENV)
		return
	}
	os.Setenv("PATH", "")
	cmd := exec.Command(os.Args[0], "-test.run=TestDialogFindPathOrExitDIALOG_TEST_ENV")
	// t.Fatalf(cmd.Path)
	// cmd.Path = ""
	cmd.Env = append(os.Environ(), "BE_CRASHER=TestDialogFindPathOrExitDIALOG_TEST_ENV")
	err := cmd.Run()
	// bs, err := cmd.CombinedOutput()
	// t.Log(bs)
	// t.Log(err)
	// e, ok := err.(*exec.ExitError)

	// if !ok && e.Success() {
	// 	if e.String() != "exit status 1" {
	// 		t.Fatalf(" %v , want exit status 1", e.String())
	// 	}
	// 	return
	// }
	if err != nil {
		t.Fatalf("process ran with err %v , want exit status 0", err)
	}

}
