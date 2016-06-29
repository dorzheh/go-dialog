// Copyright 2016 Valeriy Solovyov <weldpua2008@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style

package dialog

import (
	// "bytes"
	// "fmt"
	// "os"
	// "os/exec"
	// "strconv"
	// "strings"
	"time"
)

type ProgressIface interface {
	Step(percent int, newLabel string)
	Close()
}

type DialogIface interface {
	Shadow(truefalse bool)
	SetCancelLabel(label string)
	SetSize(height int, width int)
	SetTitle(title string)
	SetBackTitle(backtitle string)
	SetLabel(label string)
	SetOkLabel(label string)
	HelpButton(truefalse bool)
	SetHelpLabel(label string)
	SetExtraLabel(label string)
	Slider(min int, max int, step int) (int, error)
	Passivepopup(text string, timeout int)
	Geticon() string
	Getcolor() string
	Combobox(item ...string) (string, error)
	// Calendar(date time.Time) (string, error)
	// Calendar(date time.Time) (string, error)
	Checklist(listHeight int, tagItemStatus ...string) ([]string, error)
	Mixedform(title string, insecure bool, tagItemStatus ...string) ([]string, error)
	Fselect(filepath string) (string, error)
	Infobox(text string)
	Inputbox(value string) (string, error)
	Inputmenu(menuHeight int, tagItem ...string) ([]string, error)
	Menu(menuHeight int, tagItem ...string) (string, error)
	Msgbox(text string)
	Passwordbox(insecure bool) (string, error)
	Pause(seconds int)
	Textbox(filepath string) error
	Timebox(date time.Time) (string, error)
	Yesno() bool
	Radiolist(listHeight int, tagItemStatus ...string) (string, error)
	Dselect(dirpath string) (string, error)
	Progressbar() ProgressIface
}
