package dialog

import (
	"fmt"
	godialog "github.com/weldpua2008/go-dialog"
	//"strconv"
	//"time"
	"testing"
)

func NewTestDialog(environment string, parentId int) godialog.DialogIface {
	// var err error
	var res = new(godialog.Dialog)
	// if environment == AUTO || environment == "" {
	// 	for _, pkg := range []string{KDE, GTK, X, CONSOLE} {
	// 		_, err = exec.LookPath(pkg)
	// 		if err == nil {
	// 			res.environment = pkg
	// 			break
	// 		}
	// 	}
	// 	if res.environment == "" {
	// 		fmt.Println("Package not found!\nPlease install " + KDE + " or " + GTK + " or " + X + " or " + CONSOLE)
	// 	}
	// } else {
	// 	_, err = exec.LookPath(environment)
	// 	if err == nil {
	// 		res.environment = environment
	// 	} else {
	// 		fmt.Println("Package not found!\nPlease install " + environment)
	// 	}
	// }

	// if res.environment == "" {
	// 	os.Exit(1)
	// }
	// res.parentId = parentId
	// res.reset()
	return res
}

func TestInfoBox(t *testing.T) {
	d := NewTestDialog(godialog.DIALOG_TEST_ENV, 0)
	// res, err := d.Inputbox("Hello world!")
	// fmt.Println(res, err)
	res1 := d.Yesno()
	x := godialog.LastCMD
	fmt.Println("%v", x)
	fmt.Println("%v", res1)
}
