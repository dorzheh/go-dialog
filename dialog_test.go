package dialog

import (
	"fmt"
	// "github.com/weldpua2008/go-dialog"
	//"strconv"
	//"time"
	// "strings"
	// "reflect"
	"testing"
)

func NewTestDialog(environment string, parentId int) DialogIface {
	var res = new(Dialog)
	LastCMD = []string{}
	return res
}
func NewTestDialogRAW(environment string, parentId int) Dialog {
	var res = new(Dialog)
	LastCMD = []string{}
	return *res
}

func tearDown() {
	LastCMD = []string{}
}

func TestYesNo(t *testing.T) {
	d := NewTestDialog(DIALOG_TEST_ENV, 0)
	d.Yesno()
	x := LastCMD
	expected_str := "[ --no-shadow --yesno  0 0 --attach 0]"

	if fmt.Sprintf("%v", x) != expected_str {
		t.Errorf("Expected %v, actual %v ", expected_str, x)
	}

}

// tests for structure changes
func TestHelpButtonTrue(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)
	expected_val := true
	d.HelpButton(expected_val)

	if d.helpButton != expected_val {
		t.Errorf("Expected %v, actual %v ", expected_val, d.helpButton)
	}
	// x := LastCMD
	// fmt.Sprintf("%v \n", LastCMD)
	// expected_str := "[]"
	// if fmt.Sprintf("%v", x) != expected_str {
	// 	t.Errorf("Expected %v, actual %v ", expected_str, x)
	// }
}

func TestHelpButtonFalse(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)
	expected_val := false
	d.HelpButton(expected_val)

	if d.helpButton != expected_val {
		t.Errorf("Expected %v, actual %v ", expected_val, d.helpButton)
	}
}

func TestLabel(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)
	expected_val := "label"
	d.SetHelpLabel(expected_val)

	if d.helpLabel != expected_val {
		t.Errorf("Expected %v, actual %v ", expected_val, d.helpButton)
	}
}
