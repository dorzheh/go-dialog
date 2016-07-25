package dialog

// Copyright 2016 Valeriy Solovyov <weldpua2008@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style

type BaseDialog struct {
	environment       string
	parentId          int
	title             string
	backtitle         string
	label             string
	height            int
	width             int
	left              int
	top               int
	shadow            bool
	helpButton        bool
	helpLabel         string
	extraLabel        string
	cancelLabel       string
	okLabel           string
	beforeDtype       []string
	beforeSize        []string
	afterSize         []string
	lastCmd           []string
	exec_error        error
	exec_output       string
	catch_exitcode255 bool
}
