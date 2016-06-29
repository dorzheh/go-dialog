package dialog

import (
	"fmt"
	"github.com/weldpua2008/go-dialog"
	//"strconv"
	//"time"
	"testing"
)

func NewTestDialog(environment string, parentId int) dialog.DialogIface {
	var err error
	var res = new(dialog.Dialog)
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
	return &res
}

func TestInfoBox(t *testing.T) {
	d := NewTestDialog(dialog.AUTO, 0)
	// res, err := d.Inputbox("Hello world!")
	// fmt.Println(res, err)
	res1 := d.Yesno()
	fmt.Println("%v", res1)
}
