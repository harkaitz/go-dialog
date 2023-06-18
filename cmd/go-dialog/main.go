package main
import (
	"fmt"
	"github.com/harkaitz/go-dialog"
	"time"
	"os"
)
func main() {
	var ok  bool
	var err error
	var t   time.Time
	var s   string
	var l   []string
	
	dialog.AddG("--backtitle", "DIALOG TEST PROGRAM")
	
	switch os.Args[1] {
	case "Calendar":
		t, ok, err = dialog.Calendar("Calendar widget", time.Now())
		fmt.Printf("%s %s %s\n", t, ok, err)
	case "CheckList":
		l, ok, err = dialog.CheckList("CheckList widget", []dialog.TagItemStatus {
			{"tag1", "item1", "status1"},
			{"tag2", "item2", "status2"},
		})
		fmt.Printf("%d %s %s %s\n", len(l), l, ok, err)
	case "DSelect":
		s, ok, err = dialog.DSelect("/etc")
		fmt.Printf("%s %s %s\n", s, ok, err)
	case "EditBox":
		s, ok, err = dialog.EditBox("/etc/hosts")
		fmt.Printf("%s %s %s\n", s, ok, err)
	case "Form":
		l, ok, err = dialog.Form("Form widget", 20, 10, []dialog.FormField {
			{"Name"    , ""  , 1, 1},
			{"Telefono", ""  , 2, 1},
			{"Edad"    , "20", 3, 1}, 
		})
		fmt.Printf("%d %s %s %s\n", len(l), l, ok, err)
	case "FSelect":
		s, ok, err = dialog.FSelect("/etc/hosts")
		fmt.Printf("%s %s %s\n", s, ok, err)
	case "InfoBox":
		_, ok, err = dialog.InfoBox("InfoBox widget")
		fmt.Printf("%s %s\n", ok, err)
	case "InputBox":
		s, ok, err = dialog.InputBox("InputBox", "default")
		fmt.Printf("<%s> %s %s\n", s, ok, err)
	case "Menu":
		s, ok, err = dialog.Menu("Menu", []dialog.MenuItem {
			{"1", "Opción 1"},
			{"2", "Opción 2"},
		})
		fmt.Printf("<%s> %s %s\n", s, ok, err)
	case "MsgBox":
		ok, err = dialog.MsgBox("MsgBox widget")
		fmt.Printf("<%s> %s %s\n", s, ok, err)
	case "PasswordBox":
		s, ok, err = dialog.PasswordBox("PasswordBox widget", "default")
		fmt.Printf("<%s> %s %s\n", s, ok, err)
	case "Pause":
		ok, err = dialog.Pause("Pause widget", 5)
		fmt.Printf("%s %s\n", ok, err)
	case "Radiolist":
		s, ok, err = dialog.RadioList("Radiolist widget", []dialog.TagItemStatus {
			{"tag1", "item1", "status1"},
			{"tag2", "item2", "status2"},
		})
		fmt.Printf("<%s> %s %s\n", s, ok, err)
	case "TextBox":
		ok, err = dialog.TextBox("/etc/hosts")
		fmt.Printf("%s %s\n", s, ok, err)
	case "YesNo":
		ok, err = dialog.YesNo("Yes/No")
		fmt.Printf("%s %s\n", ok, err)
	}
	
	
}
